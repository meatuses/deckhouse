# Copyright 2022 Flant JSC
# Licensed under the Deckhouse Platform Enterprise Edition (EE) license. See https://github.com/deckhouse/deckhouse/blob/main/ee/LICENSE.

{{- $kubernetesCniVersion := "1.2.0" | replace "." "" }}
bb-rp-install "kubeadm:{{ index .images.registrypackages (printf "kubeadm%s" $kubernetesVersion) }}" "kubelet:{{ index .images.registrypackages (printf "kubeletAlteros%s" $kubernetesVersion) }}" "kubectl:{{ index .images.registrypackages (printf "kubectlAlteros%s" $kubernetesVersion) }}" "crictl:{{ index .images.registrypackages (printf "crictl%s" $kubernetesMajorVersion) }}" "kubernetes-cni:{{ index .images.registrypackages (printf "kubernetesCni%s" $kubernetesCniVersion) }}"
