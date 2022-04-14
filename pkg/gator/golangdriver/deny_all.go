package golangdriver

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func DenyAll(_ interface{}) Constraint {
	return func(_ map[string]*unstructured.Unstructured, _ *unstructured.Unstructured) *types.Result {
		return &types.Result{
			Msg: "deny all",
		}
	}
}
