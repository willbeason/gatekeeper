package webhook

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"

	opa "github.com/open-policy-agent/frameworks/constraint/pkg/client"
	"github.com/open-policy-agent/gatekeeper/pkg/controller/config/process"
	"github.com/open-policy-agent/gatekeeper/pkg/mutation"
	"github.com/pkg/errors"
	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	exemptNamespace       = newNSSet()
	exemptNamespacePrefix = newNSSet()
)

func init() {
	AddToManagerFuncs = append(AddToManagerFuncs, AddLabelWebhook)
	flag.Var(exemptNamespace, "exempt-namespace", "The specified namespace is allowed to set the admission.gatekeeper.sh/ignore label. To exempt multiple namespaces, this flag can be declared more than once.")
	flag.Var(exemptNamespacePrefix, "exempt-namespace-prefix", "A namespace with the specified prefix is allowed to set the admission.gatekeeper.sh/ignore label. To exempt multiple prefixes, this flag can be declared more than once.")
}

const ignoreLabel = "admission.gatekeeper.sh/ignore"

type nsSet map[string]bool

var _ flag.Value = nsSet{}

func newNSSet() nsSet {
	return make(map[string]bool)
}

func (l nsSet) String() string {
	contents := make([]string, 0)
	for k := range l {
		contents = append(contents, k)
	}
	return fmt.Sprintf("%s", contents)
}

func (l nsSet) Set(s string) error {
	l[s] = true
	return nil
}

// +kubebuilder:webhook:verbs=CREATE;UPDATE,path=/v1/admitlabel,mutating=false,failurePolicy=fail,groups="",resources=namespaces,versions=*,name=check-ignore-label.gatekeeper.sh

// AddLabelWebhook registers the label webhook server with the manager
func AddLabelWebhook(mgr manager.Manager, _ *opa.Client, _ *process.Excluder, mutationCache *mutation.System) error {
	wh := &admission.Webhook{Handler: &namespaceLabelHandler{}}
	// TODO(https://github.com/open-policy-agent/gatekeeper/issues/661): remove log injection if the race condition in the cited bug is eliminated.
	// Otherwise we risk having unstable logger names for the webhook.
	if err := wh.InjectLogger(log); err != nil {
		return err
	}
	mgr.GetWebhookServer().Register("/v1/admitlabel", wh)
	return nil
}

var _ admission.Handler = &namespaceLabelHandler{}

type namespaceLabelHandler struct{}

func (h *namespaceLabelHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	if req.Operation == admissionv1.Delete {
		return admission.Allowed("Delete is always allowed")
	}
	if req.AdmissionRequest.Kind.Group != "" || req.AdmissionRequest.Kind.Kind != namespaceKind {
		return admission.Allowed("Not a namespace")
	}
	obj := &unstructured.Unstructured{}
	if err := json.Unmarshal(req.Object.Raw, obj); err != nil {
		r := admission.Denied(errors.Wrap(err, "while deserializing resource").Error())
		r.Result.Code = http.StatusInternalServerError
		return r
	}
	if exemptNamespace[obj.GetName()] || matchesPrefix(obj.GetName()) {
		return admission.Allowed(fmt.Sprintf("Namespace %s is allowed to set %s", obj.GetName(), ignoreLabel))
	}
	for label := range obj.GetLabels() {
		if label == ignoreLabel {
			return admission.Denied(fmt.Sprintf("Only exempt namespace can have the %s label", ignoreLabel))
		}
	}
	return admission.Allowed(fmt.Sprintf("Namespace is not setting the %s label", ignoreLabel))
}

func matchesPrefix(s string) bool {
	for p := range exemptNamespacePrefix {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
