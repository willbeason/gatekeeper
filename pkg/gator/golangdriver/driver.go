package golangdriver

import (
	"context"

	"github.com/open-policy-agent/frameworks/constraint/pkg/client/drivers"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"github.com/open-policy-agent/opa/storage"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GVKNN struct {
	schema.GroupVersionKind
	Namespace string
	Name string
}

type Driver struct{
	templates map[string]Template
	constraints map[string]map[string]Constraint

	storage map[GVKNN]*unstructured.Unstructured
}

func (d Driver) AddTemplate(ctx context.Context, ct *templates.ConstraintTemplate) error {
	kind := ct.Spec.CRD.Spec.Names.Kind
	entry := ct.Spec.Targets[0].Rego

	d.templates[kind] = library[entry]
	if _, found := d.constraints[kind]; !found {
		d.constraints[kind] = make(map[string]Constraint)
	}

	return nil
}

func (d Driver) RemoveTemplate(ctx context.Context, ct *templates.ConstraintTemplate) error {
	kind := ct.Spec.CRD.Spec.Names.Kind

	delete(d.templates, kind)
	delete(d.constraints, kind)

	return nil
}

func (d Driver) AddConstraint(ctx context.Context, constraint *unstructured.Unstructured) error {
	kind := constraint.GetKind()
	c := d.templates[kind](constraint)

	d.constraints[kind][constraint.GetName()] = c

	return nil
}

func (d Driver) RemoveConstraint(ctx context.Context, constraint *unstructured.Unstructured) error {
	kind := constraint.GetKind()
	kindConstraints := d.constraints[kind]

	delete(kindConstraints, constraint.GetName())

	return nil
}

func (d Driver) AddData(ctx context.Context, target string, path storage.Path, data interface{}) error {
	obj := &unstructured.Unstructured{Object: data.(map[string]interface{})}
	gvk := obj.GroupVersionKind()

	gvknn := GVKNN{
		GroupVersionKind: gvk,
		Namespace: obj.GetNamespace(),
		Name: obj.GetName(),
	}

	d.storage[gvknn] = obj

	return nil
}

func (d Driver) RemoveData(ctx context.Context, target string, path storage.Path) error {
	obj := &unstructured.Unstructured{Object: data.(map[string]interface{})}
	gvk := obj.GroupVersionKind()

	gvknn := GVKNN{
		GroupVersionKind: gvk,
		Namespace: obj.GetNamespace(),
		Name: obj.GetName(),
	}

	delete(d.storage, gvknn)

	return nil
}

func (d Driver) Query(ctx context.Context, target string, constraints []*unstructured.Unstructured, review interface{}, opts ...drivers.QueryOpt) ([]*types.Result, *string, error) {
	obj := &unstructured.Unstructured{
		Object: review.(map[string]interface{}),
	}

	var results []*types.Result
	for _, constraints := range d.constraints {
		for _, constraint := range constraints {
			result := constraint(d.storage, obj)
			if result != nil {
				results = append(results, result)
			}
		}
	}

	return results, nil, nil
}

func (d Driver) Dump(ctx context.Context) (string, error) {
	// TODO implement me
	panic("implement me")
}
var _ drivers.Driver = Driver{}
