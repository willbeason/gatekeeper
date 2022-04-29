package golangdriver

import (
	"fmt"

	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const UniqueIngressHostKey = "unique-ingress-host"

func UniqueIngressHost(_ interface{}) ConstraintFn {
	return func(storage map[string]*unstructured.Unstructured, review *unstructured.Unstructured) *types.Result {
		hosts := getHosts(review.Object)
		for host := range hosts {
			fmt.Println("New host", host)
		}

		for _, stored := range storage {
			gvk := stored.GroupVersionKind()
			if gvk.Kind != "Ingress" {
				continue
			}
			if gvk.Group != "extensions" && gvk.Group != "networking.k8s.io" {
				continue
			}
			if review.GetName() == stored.GetName() && review.GetNamespace() == stored.GetNamespace() {
				fmt.Println("Identical ingress")
				continue
			}

			fmt.Println("Existing ingress", stored.GetNamespace(), stored.GetName())

			storedHosts := getHosts(stored.Object)

			for h := range storedHosts {
				fmt.Println("Existing host", h)

				if hosts[h] {
					conflict := h
					return &types.Result{
						Msg: fmt.Sprintf("ingress host conflicts with an existing ingress (go): %q", conflict),
					}
				}
			}
		}

		return nil
	}
}

func getHosts(obj map[string]interface{}) map[string]bool {
	hosts := make(map[string]bool)

	hostsSlice, _, _ := unstructured.NestedSlice(obj, "spec", "rules")
	for _, hostI := range hostsSlice {
		hostM := hostI.(map[string]interface{})
		host := hostM["host"].(string)
		hosts[host] = true
	}

	return hosts
}
