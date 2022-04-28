package golangdriver

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/constraints"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	target2 "github.com/open-policy-agent/gatekeeper/pkg/target"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func FakeConstraintTemplate(kind, entry string) *templates.ConstraintTemplate {
	return &templates.ConstraintTemplate{
		Spec: templates.ConstraintTemplateSpec{
			CRD: templates.CRD{
				Spec: templates.CRDSpec{
					Names: templates.Names{
						Kind: kind,
					},
				},
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				Annotation: entry,
			},
		},
	}
}

func FakeConstraint(kind, name string) *unstructured.Unstructured {
	obj := &unstructured.Unstructured{Object: make(map[string]interface{})}

	obj.SetKind(kind)
	obj.SetName(name)

	return obj
}

func TestDriver_Query(t *testing.T) {
	tests := []struct {
		name        string
		templates   []*templates.ConstraintTemplate
		constraints []*unstructured.Unstructured
		storage     []*unstructured.Unstructured
		review      map[string]interface{}
		want        []*types.Result
	}{
		{
			name: "unique ingress host success",
			templates: []*templates.ConstraintTemplate{
				FakeConstraintTemplate("bar", UniqueIngressHostKey),
			},
			constraints: []*unstructured.Unstructured{
				FakeConstraint("bar", "bar-1"),
			},
			storage: []*unstructured.Unstructured{
				FakeIngress("foo-1", "foo"),
			},
			review: FakeIngress("qux-1", "qux").Object,
			want:   nil,
		},
		{
			name: "unique ingress host failure",
			templates: []*templates.ConstraintTemplate{
				FakeConstraintTemplate("bar", UniqueIngressHostKey),
			},
			constraints: []*unstructured.Unstructured{
				FakeConstraint("bar", "bar-1"),
			},
			storage: []*unstructured.Unstructured{
				FakeIngress("foo-1", "foo"),
			},
			review: FakeIngress("foo-2", "foo").Object,
			want: []*types.Result{{
				Msg:               `ingress host conflicts with an existing ingress (go): "foo"`,
				Constraint:        FakeConstraint("bar", "bar-1"),
				EnforcementAction: constraints.EnforcementActionDeny,
			}},
		},
		{
			name: "replace Template",
			templates: []*templates.ConstraintTemplate{
				FakeConstraintTemplate("bar", UniqueIngressHostKey),
				FakeConstraintTemplate("bar", "x-"+UniqueIngressHostKey),
			},
			constraints: []*unstructured.Unstructured{
				FakeConstraint("bar", "bar-1"),
			},
			storage: []*unstructured.Unstructured{
				FakeIngress("foo-1", "foo"),
			},
			review: FakeIngress("foo-2", "foo").Object,
			want:   nil,
		},
	}

	target := &target2.K8sValidationTarget{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := NewDriver()
			ctx := context.Background()

			for _, template := range tt.templates {
				err := driver.AddTemplate(ctx, template)
				if err != nil {
					t.Fatal(err)
				}
			}

			for _, constraint := range tt.constraints {
				err := driver.AddConstraint(ctx, constraint)
				if err != nil {
					t.Fatal(err)
				}
			}

			for _, stored := range tt.storage {
				_, path, _, _ := target.ProcessData(stored)
				err := driver.AddData(ctx, "", path, stored.Object)
				if err != nil {
					t.Fatal(err)
				}
			}

			u := &unstructured.Unstructured{Object: tt.review}
			jsn, err := u.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}

			obj := &target2.GkReview{AdmissionRequest: v1.AdmissionRequest{Object: runtime.RawExtension{Raw: jsn}}}

			got, _, err := driver.Query(ctx, "", tt.constraints, obj)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
