/*
Copyright 2023 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hooks

import (
	"sort"

	"github.com/flant/addon-operator/pkg/module_manager/go_hook"
	"github.com/flant/addon-operator/sdk"
	"github.com/flant/shell-operator/pkg/kube_events_manager/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/deckhouse/deckhouse/modules/110-istio/hooks/lib"
)

type IstioNamespaceFilterResult struct {
	Name                    string
	DeletionTimestampExists bool
	Revision                string
}

func applyNamespaceFilter(obj *unstructured.Unstructured) (go_hook.FilterResult, error) {
	_, deletionTimestampExists := obj.GetAnnotations()["deletionTimestamp"]

	var namespaceInfo = IstioNamespaceFilterResult{
		Name:                    obj.GetName(),
		DeletionTimestampExists: deletionTimestampExists,
	}

	return namespaceInfo, nil
}

func applyDiscoveryAppIstioPodFilter(obj *unstructured.Unstructured) (go_hook.FilterResult, error) {
	var namespaceInfo = IstioNamespaceFilterResult{
		Name: obj.GetNamespace(),
	}
	return namespaceInfo, nil
}

var _ = sdk.RegisterFunc(&go_hook.HookConfig{
	Queue: lib.Queue("discovery"),
	Kubernetes: []go_hook.KubernetesConfig{
		{
			Name:          "namespaces_global_revision",
			ApiVersion:    "v1",
			Kind:          "Namespace",
			FilterFunc:    applyNamespaceFilter,
			LabelSelector: &metav1.LabelSelector{MatchLabels: map[string]string{"istio-injection": "enabled"}},
		},
		{
			Name:       "namespaces_definite_revision",
			ApiVersion: "v1",
			Kind:       "Namespace",
			FilterFunc: applyNamespaceFilter,
			LabelSelector: &metav1.LabelSelector{
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{
						Key:      "istio.io/rev",
						Operator: metav1.LabelSelectorOpExists,
					},
				},
			},
		},
		{
			Name:       "istio_pod_global_rev",
			ApiVersion: "v1",
			Kind:       "Pod",
			FilterFunc: applyDiscoveryAppIstioPodFilter,
			LabelSelector: &metav1.LabelSelector{
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{
						Key:      "sidecar.istio.io/inject",
						Operator: metav1.LabelSelectorOpIn,
						Values:   []string{"true"},
					},
					{
						Key:      "istio.io/rev",
						Operator: metav1.LabelSelectorOpDoesNotExist,
					},
				},
			},
		},
		{
			Name:       "istio_pod_definite_rev",
			ApiVersion: "v1",
			Kind:       "Pod",
			FilterFunc: applyDiscoveryAppIstioPodFilter,
			NamespaceSelector: &types.NamespaceSelector{
				LabelSelector: &metav1.LabelSelector{
					MatchExpressions: []metav1.LabelSelectorRequirement{
						{
							Key:      "istio.io/rev",
							Operator: metav1.LabelSelectorOpDoesNotExist,
						},
						{
							Key:      "istio-injection",
							Operator: metav1.LabelSelectorOpNotIn,
							Values:   []string{"enabled"},
						},
					},
				},
			},
			LabelSelector: &metav1.LabelSelector{
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{
						Key:      "istio.io/rev",
						Operator: metav1.LabelSelectorOpExists,
					},
				},
			},
		},
	},
}, applicationNamespacesDiscovery)

func applicationNamespacesDiscovery(input *go_hook.HookInput) error {
	var applicationNamespaces = make([]string, 0)
	var namespaces = make([]go_hook.FilterResult, 0)
	namespaces = append(namespaces, input.Snapshots["namespaces_definite_revision"]...)
	namespaces = append(namespaces, input.Snapshots["namespaces_global_revision"]...)
	namespaces = append(namespaces, input.Snapshots["istio_pod_global_rev"]...)
	namespaces = append(namespaces, input.Snapshots["istio_pod_definite_rev"]...)
	for _, ns := range namespaces {
		nsInfo := ns.(IstioNamespaceFilterResult)
		if nsInfo.DeletionTimestampExists {
			continue
		}
		if !lib.Contains(applicationNamespaces, nsInfo.Name) {
			applicationNamespaces = append(applicationNamespaces, nsInfo.Name)
		}
	}

	sort.Strings(applicationNamespaces)

	input.Values.Set("istio.internal.applicationNamespaces", applicationNamespaces)

	return nil
}
