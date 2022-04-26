package cel

import (
	"context"

	"github.com/google/cel-policy-templates-go/policy"
	"github.com/open-policy-agent/frameworks/constraint/pkg/client/drivers"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"github.com/open-policy-agent/opa/storage"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Driver struct {
	engines map[string]policy.Engine
}

var _ drivers.Driver = &Driver{}

func (d *Driver) AddTemplate(ctx context.Context, ct *templates.ConstraintTemplate) error {
	_, err := policy.NewEngine()
	if err != nil {
		return err
	}

	return nil
}

func (d *Driver) RemoveTemplate(ctx context.Context, ct *templates.ConstraintTemplate) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) AddConstraint(ctx context.Context, constraint *unstructured.Unstructured) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) RemoveConstraint(ctx context.Context, constraint *unstructured.Unstructured) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) AddData(ctx context.Context, target string, path storage.Path, data interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) RemoveData(ctx context.Context, target string, path storage.Path) error {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) Query(ctx context.Context, target string, constraints []*unstructured.Unstructured, review interface{}, opts ...drivers.QueryOpt) ([]*types.Result, *string, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Driver) Dump(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}
