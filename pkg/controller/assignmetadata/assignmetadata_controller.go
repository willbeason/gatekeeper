/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package assignmetadata

import (
	"context"
	"fmt"

	opa "github.com/open-policy-agent/frameworks/constraint/pkg/client"
	mutationsv1alpha1 "github.com/open-policy-agent/gatekeeper/apis/mutations/v1alpha1"
	"github.com/open-policy-agent/gatekeeper/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/pkg/mutation"
	"github.com/open-policy-agent/gatekeeper/pkg/mutation/types"
	"github.com/open-policy-agent/gatekeeper/pkg/readiness"
	"github.com/open-policy-agent/gatekeeper/pkg/watch"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	log = logf.Log.WithName("controller").WithValues(logging.Process, "assignmetadata_controller")
)

var gvkAssignMetadata = schema.GroupVersionKind{
	Group:   mutationsv1alpha1.GroupVersion.Group,
	Version: mutationsv1alpha1.GroupVersion.Version,
	Kind:    "AssignMetadata",
}

type Adder struct {
	MutationCache *mutation.System
	Tracker       *readiness.Tracker
}

// Add creates a new AssignMetadata Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func (a *Adder) Add(mgr manager.Manager) error {
	r := newReconciler(mgr, a.MutationCache, a.Tracker)
	return add(mgr, r)
}

func (a *Adder) InjectOpa(o *opa.Client) {}

func (a *Adder) InjectWatchManager(w *watch.Manager) {}

func (a *Adder) InjectControllerSwitch(cs *watch.ControllerSwitch) {}

func (a *Adder) InjectTracker(t *readiness.Tracker) {
	a.Tracker = t
}

func (a *Adder) InjectMutationCache(mutationCache *mutation.System) {
	a.MutationCache = mutationCache
}

// newReconciler returns a new reconcile.Reconciler

func newReconciler(mgr manager.Manager, mutationCache *mutation.System, tracker *readiness.Tracker) *Reconciler {
	r := &Reconciler{system: mutationCache, Client: mgr.GetClient(), tracker: tracker}
	return r
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	if !*mutation.MutationEnabled {
		return nil
	}

	// Create a new controller
	c, err := controller.New("assignmetadata-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to AssignMetadata
	if err = c.Watch(
		&source.Kind{Type: &mutationsv1alpha1.AssignMetadata{}},
		&handler.EnqueueRequestForObject{}); err != nil {
		return err
	}
	return nil
}

// Reconciler reconciles a AssignMetadata object
type Reconciler struct {
	client.Client
	system  *mutation.System
	tracker *readiness.Tracker
}

// +kubebuilder:rbac:groups=mutations.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete

// Reconcile reads that state of the cluster for a AssignMetadata object and makes changes based on the state read
// and what is in the AssignMetadata.Spec
func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log.Info("Reconcile", "request", request)
	deleted := false
	assignMetadata := &mutationsv1alpha1.AssignMetadata{}
	err := r.Get(ctx, request.NamespacedName, assignMetadata)
	if err != nil {
		if !errors.IsNotFound(err) {
			return reconcile.Result{}, err
		}
		deleted = true
		assignMetadata = &mutationsv1alpha1.AssignMetadata{
			ObjectMeta: metav1.ObjectMeta{
				Name:      request.NamespacedName.Name,
				Namespace: request.NamespacedName.Namespace,
			},
			TypeMeta: metav1.TypeMeta{
				Kind:       "AssignMetadata",
				APIVersion: fmt.Sprintf("%s/%s", mutationsv1alpha1.GroupVersion.Group, mutationsv1alpha1.GroupVersion.Version),
			},
		}
	}
	deleted = deleted || !assignMetadata.GetDeletionTimestamp().IsZero()
	tracker := r.tracker.For(gvkAssignMetadata)

	if deleted {
		id, err := types.MakeID(assignMetadata)
		if err != nil {
			log.Error(err, "Failed to get id out of metadata")
			return ctrl.Result{}, nil
		}

		if err := r.system.Remove(id); err != nil {
			log.Error(err, "Remove failed", "resource", request.NamespacedName)
		}
		return ctrl.Result{}, nil
	}

	mutator, err := mutation.MutatorForAssignMetadata(assignMetadata)
	if err != nil {
		log.Error(err, "Creating mutator for resource failed", "resource", request.NamespacedName)
		tracker.CancelExpect(assignMetadata)
		return ctrl.Result{}, err
	}
	if !deleted {
		if err := r.system.Upsert(mutator); err != nil {
			log.Error(err, "Insert failed", "resource", request.NamespacedName)
			tracker.TryCancelExpect(assignMetadata)
		} else {
			tracker.Observe(assignMetadata)
		}
	} else {
		if err := r.system.Remove(mutator.ID()); err != nil {
			log.Error(err, "Remove failed", "resource", request.NamespacedName)
		}
		tracker.CancelExpect(assignMetadata)
	}

	return ctrl.Result{}, nil
}
