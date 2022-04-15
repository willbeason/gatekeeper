package golangdriver

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func FakeIngress(name string, hosts ...string) *unstructured.Unstructured {
	result := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}

	result.SetName(name)

	result.SetGroupVersionKind(schema.GroupVersionKind{
		Kind:  "Ingress",
		Group: "networking.k8s.io",
	})

	hostList := make([]interface{}, 0, len(hosts))
	for _, host := range hosts {
		hostMap := map[string]interface{}{
			"host": host,
		}
		hostList = append(hostList, hostMap)
	}

	_ = unstructured.SetNestedSlice(result.Object, hostList, "spec", "rules")

	return result
}

func TestUniqueIngressHost(t *testing.T) {
	constraint := UniqueIngressHost(nil)

	tests := []struct {
		name    string
		storage map[string]*unstructured.Unstructured
		review  *unstructured.Unstructured
		want    *types.Result
	}{
		{
			name:    "no storage",
			storage: nil,
			review:  FakeIngress("foo-1"),
			want:    nil,
		},
		{
			name: "conflict",
			storage: map[string]*unstructured.Unstructured{
				"foo": FakeIngress("foo-1", "foo"),
			},
			review: FakeIngress("foo-2", "foo"),
			want: &types.Result{
				Msg: `ingress host conflicts with an existing ingress (go): "foo"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := constraint(tt.storage, tt.review)

			if diff := cmp.Diff(tt.want, result); diff != "" {
				t.Error(diff)
			}
		})
	}
}
