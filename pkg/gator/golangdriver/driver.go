package golangdriver

import (
	"context"
	"strings"

	constraints2 "github.com/open-policy-agent/frameworks/constraint/pkg/apis/constraints"
	"github.com/open-policy-agent/frameworks/constraint/pkg/client/drivers"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	target2 "github.com/open-policy-agent/gatekeeper/pkg/target"
	"github.com/open-policy-agent/opa/storage"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const Annotation = "golib"

type Driver struct {
	templates   map[string]Template
	constraints map[string]map[string]Constraint

	storage map[string]*unstructured.Unstructured
}

func NewDriver() *Driver {
	return &Driver{
		templates:   make(map[string]Template),
		constraints: make(map[string]map[string]Constraint),
		storage:     make(map[string]*unstructured.Unstructured),
	}
}

func (d *Driver) AddTemplate(ctx context.Context, ct *templates.ConstraintTemplate) error {
	kind := ct.Spec.CRD.Spec.Names.Kind
	entry := ct.Annotations[Annotation]

	templateFn, found := library[entry]
	if !found {
		templateFn = func(params interface{}) ConstraintFn {
			return func(storage map[string]*unstructured.Unstructured, review *unstructured.Unstructured) *types.Result {
				return nil
			}
		}
	}
	d.templates[kind] = templateFn

	if constraints, found := d.constraints[kind]; found {
		for name, constraint := range constraints {
			constraint.fn = templateFn(constraint)
			constraints[name] = constraint
		}
		d.constraints[kind] = constraints
	} else {
		d.constraints[kind] = make(map[string]Constraint)
	}

	return nil
}

func (d *Driver) RemoveTemplate(ctx context.Context, ct *templates.ConstraintTemplate) error {
	kind := ct.Spec.CRD.Spec.Names.Kind

	delete(d.templates, kind)
	delete(d.constraints, kind)

	return nil
}

func (d *Driver) AddConstraint(ctx context.Context, constraint *unstructured.Unstructured) error {
	kind := constraint.GetKind()

	template, found := d.templates[kind]
	if !found {
		return nil
	}

	c := template(constraint)

	kindConstraints, found := d.constraints[kind]
	if !found {
		return nil
	}

	kindConstraints[constraint.GetName()] = Constraint{
		constraint: constraint,
		fn:         c,
	}

	return nil
}

func (d *Driver) RemoveConstraint(ctx context.Context, constraint *unstructured.Unstructured) error {
	kind := constraint.GetKind()
	kindConstraints := d.constraints[kind]

	delete(kindConstraints, constraint.GetName())

	return nil
}

func (d *Driver) AddData(ctx context.Context, target string, path storage.Path, data interface{}) error {
	obj := &unstructured.Unstructured{Object: data.(map[string]interface{})}
	key := ToKey(path)

	d.storage[key] = obj

	return nil
}

func (d *Driver) RemoveData(ctx context.Context, target string, path storage.Path) error {
	key := ToKey(path)

	delete(d.storage, key)

	return nil
}

func (d *Driver) Query(ctx context.Context, target string, constraints []*unstructured.Unstructured, review interface{}, opts ...drivers.QueryOpt) ([]*types.Result, *string, error) {
	gkr := review.(*target2.GkReview)

	obj := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}

	err := obj.UnmarshalJSON(gkr.Object.Raw)
	if err != nil {
		return nil, nil, err
	}

	var results []*types.Result
	for _, constraint := range constraints {
		kind := constraint.GetKind()
		kindConstraints, found := d.constraints[kind]
		if !found {
			continue
		}

		name := constraint.GetName()
		toRun, found := kindConstraints[name]
		if !found {
			continue
		}

		result := toRun.fn(d.storage, obj)
		if result != nil {
			result.Constraint = constraint.DeepCopy()
			result.EnforcementAction = constraints2.EnforcementActionDeny
			result.Target = target
			results = append(results, result)
		}
	}

	return results, nil, nil
}

func (d *Driver) Dump(ctx context.Context) (string, error) {
	// TODO implement me
	panic("implement me")
}

var _ drivers.Driver = &Driver{}

func ToKey(path []string) string {
	return strings.Join(path, "/")
}
