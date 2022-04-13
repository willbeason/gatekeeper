package golangdriver

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Template func(params interface{}) Constraint

type Constraint func(storage map[string]*unstructured.Unstructured, review *unstructured.Unstructured) *types.Result

type ConstraintKey struct {
	Kind string
	Name string
}

var library = map[string]Template{
	UniqueIngressHostKey: UniqueIngressHost,
}
