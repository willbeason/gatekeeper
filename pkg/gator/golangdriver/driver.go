package golangdriver

import (
	"context"
	"strings"

	constraints2 "github.com/open-policy-agent/frameworks/constraint/pkg/apis/constraints"
	"github.com/open-policy-agent/frameworks/constraint/pkg/client/drivers"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
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

	d.templates[kind] = library[entry]
	if _, found := d.constraints[kind]; !found {
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
	c := d.templates[kind](constraint)

	d.constraints[kind][constraint.GetName()] = c

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
	obj := &unstructured.Unstructured{
		Object: review.(map[string]interface{}),
	}

	var results []*types.Result
	for _, constraint := range constraints {
		kind := constraint.GetKind()
		kindConstraints := d.constraints[kind]
		name := constraint.GetName()
		toRun := kindConstraints[name]

		result := toRun(d.storage, obj)
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
