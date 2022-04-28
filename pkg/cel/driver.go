package cel

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-policy-templates-go/policy"
	"github.com/google/cel-policy-templates-go/policy/model"
	constraints2 "github.com/open-policy-agent/frameworks/constraint/pkg/apis/constraints"
	"github.com/open-policy-agent/frameworks/constraint/pkg/client/drivers"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	target2 "github.com/open-policy-agent/gatekeeper/pkg/target"
	"github.com/open-policy-agent/opa/storage"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func NewDriver() *Driver {
	return &Driver{
		engines: make(map[string]*policy.Engine),
	}
}

type Driver struct {
	engines map[string]*policy.Engine
}

var _ drivers.Driver = &Driver{}

func (d *Driver) AddTemplate(ctx context.Context, ct *templates.ConstraintTemplate) error {
	env, _ := cel.NewEnv(cel.Declarations(
		decls.NewVar("resource", decls.Any),
	))

	if len(ct.Spec.Targets) == 0 {
		return nil
	}

	celCode := ct.Spec.Targets[0].CELX
	if celCode == "" {
		return nil
	}

	engine, err := policy.NewEngine(policy.StandardExprEnv(env))
	if err != nil {
		return err
	}

	templ, iss := engine.CompileTemplate(model.StringSource(celCode, ""))
	if iss.Err() != nil {
		return iss.Err()
	}

	err = engine.SetTemplate(templ.Metadata.Name, templ)
	if err != nil {
		return err
	}

	d.engines[ct.Name] = engine
	return nil
}

func (d *Driver) RemoveTemplate(ctx context.Context, ct *templates.ConstraintTemplate) error {
	delete(d.engines, ct.Name)

	return nil
}

func (d *Driver) AddConstraint(ctx context.Context, constraint *unstructured.Unstructured) error {
	celCode, found, err := unstructured.NestedString(constraint.Object, "spec", "cel")
	if err != nil {
		return err
	}
	if !found {
		return nil
	}

	engineName := strings.ToLower(constraint.GetKind())

	engine, found := d.engines[engineName]
	if !found {
		return fmt.Errorf("no engine %q", engineName)
	}

	inst, iss := engine.CompileInstance(model.StringSource(celCode, ""))
	if iss.Err() != nil {
		return iss.Err()
	}

	return engine.AddInstance(inst)
}

func (d *Driver) RemoveConstraint(ctx context.Context, constraint *unstructured.Unstructured) error {
	// CEL doesn't support removing instances from an engine.
	return nil
}

func (d *Driver) AddData(ctx context.Context, target string, path storage.Path, data interface{}) error {
	// CEL doesn't support referential Constraints.
	return nil
}

func (d *Driver) RemoveData(ctx context.Context, target string, path storage.Path) error {
	return nil
}

func (d *Driver) Query(ctx context.Context, target string, constraints []*unstructured.Unstructured, review interface{}, opts ...drivers.QueryOpt) ([]*types.Result, *string, error) {
	// CEL doesn't support running a subset of added Constraints.
	// At least, not by name. You could jury-rig something together with Selectors which would do this, but it'd be messy.
	engines := make(map[string]bool)
	for _, constraint := range constraints {
		engines[strings.ToLower(constraint.GetKind())] = true
	}

	gkr := review.(*target2.GkReview)

	obj := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}

	err := obj.UnmarshalJSON(gkr.Object.Raw)
	if err != nil {
		return nil, nil, err
	}

	resource := map[string]interface{}{
		"resource": obj.Object,
	}

	var allDecisions []model.DecisionValue
	for name := range engines {
		engine, found := d.engines[name]
		if !found {
			continue
		}
		decisions, err := engine.EvalAll(resource)
		if err != nil {
			return nil, nil, err
		}
		allDecisions = append(allDecisions, decisions...)
	}
	if len(allDecisions) == 0 {
		return nil, nil, nil
	}

	results := make([]*types.Result, len(allDecisions))
	for i, decision := range allDecisions {
		results[i] = &types.Result{
			Metadata: map[string]interface{}{
				"name": decision.Name(),
			},
			Msg:               decision.String(),
			EnforcementAction: constraints2.EnforcementActionDeny,
		}
	}

	return results, nil, nil
}

func (d *Driver) Dump(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}
