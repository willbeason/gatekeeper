package cel

import (
	"context"
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-policy-templates-go/policy"
	"github.com/google/cel-policy-templates-go/policy/model"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/constraints"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"github.com/open-policy-agent/gatekeeper/pkg/target"
	v1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

const requireLabelsTemplate = `
apiVersion: policy.acme.co/v1
kind: PolicyTemplate
metadata:
  name: required_labels
  namespace: acme
schema:
  type: object
  required:
    - labels
  properties:
    labels:
      type: object
      additionalProperties: {}
evaluator:
  terms:
    want: rule.labels
    missing: want.filter(l, !(l in resource.metadata.labels))
    invalid: >
      resource.metadata.labels.filter(l,
        l in want && want[l] != resource.metadata.labels[l])
  productions:
    - match: missing.size() > 0
      decision: policy.violation1
      output:
        message: missing one or more required labels
        details:
          data: missing
    - match: invalid.size() > 0
      decision: policy.violation2
      output:
        message: invalid values provided on one or more labels
        details:
          data: invalid
`

const requireLabelsInstance = `
apiVersion: policy.acme.co/v1
kind: required_labels
metadata:
  name: prod_labels
  namespace: acme
rules:
  - labels:
      env: prod
      ssh: enabled
      verified: true
`

func newTemplate(name string, cel string) *templates.ConstraintTemplate {
	result := &templates.ConstraintTemplate{}
	result.Name = name

	result.Spec.Targets = []templates.Target{{
		CELX: cel,
	}}

	return result
}

func newConstraint(kind string, cel string) *unstructured.Unstructured {
	result := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}

	result.SetKind(kind)

	err := unstructured.SetNestedField(result.Object, cel, "spec", "cel")
	if err != nil {
		panic(err)
	}

	return result
}

func TestDriver_Query(t *testing.T) {
	env, _ := cel.NewEnv(cel.Declarations(
		decls.NewVar("resource", decls.Any),
	))

	tests := []struct {
		name       string
		ct         *templates.ConstraintTemplate
		constraint *unstructured.Unstructured
		obj        *unstructured.Unstructured
		want       []*types.Result
	}{{
		name:       "require labels satisfied",
		ct:         newTemplate("requirelabels", requireLabelsTemplate),
		constraint: newConstraint("RequireLabels", requireLabelsInstance),
		obj: &unstructured.Unstructured{Object: map[string]interface{}{
			"kind": "Foo",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"env":      "prod",
					"ssh":      "enabled",
					"verified": true,
				},
			},
		}},
		want: nil,
	}, {
		name:       "require labels violated",
		ct:         newTemplate("requirelabels", requireLabelsTemplate),
		constraint: newConstraint("RequireLabels", requireLabelsInstance),
		obj: &unstructured.Unstructured{Object: map[string]interface{}{
			"kind": "Foo",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"env":       "prod",
					"ssh":       "enabled",
					"verifiedd": true,
				},
			},
		}},
		want: []*types.Result{{
			EnforcementAction: constraints.EnforcementActionDeny,
			Metadata: map[string]interface{}{
				"name": "policy.violation1",
			},
		}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := NewDriver()

			ctx := context.Background()
			err := driver.AddTemplate(ctx, tt.ct)
			if err != nil {
				t.Fatal(err)
			}

			err = driver.AddConstraint(ctx, tt.constraint)
			if err != nil {
				t.Fatal(err)
			}

			engine, err := policy.NewEngine(policy.StandardExprEnv(env))
			if err != nil {
				t.Fatal(err)
			}

			templ, iss := engine.CompileTemplate(model.StringSource(tt.ct.Spec.Targets[0].CELX, ""))
			if iss.Err() != nil {
				t.Fatal(iss.Err())
			}

			err = engine.SetTemplate("required_labels", templ)
			if err != nil {
				t.Fatal(err)
			}

			inst, iss := engine.CompileInstance(model.StringSource(tt.constraint.Object["spec"].(map[string]interface{})["cel"].(string), ""))
			if iss.Err() != nil {
				t.Fatal(iss.Err())
			}

			err = engine.AddInstance(inst)
			if err != nil {
				t.Fatal(err)
			}

			raw, err := tt.obj.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}

			decisions, _, err := driver.Query(ctx, "", []*unstructured.Unstructured{tt.constraint}, &target.GkReview{AdmissionRequest: v1.AdmissionRequest{Object: runtime.RawExtension{Raw: raw}}})
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.want, decisions, cmpopts.IgnoreFields(types.Result{}, "Msg")); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
