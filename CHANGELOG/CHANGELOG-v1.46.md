# Changelog v1.46

## [MALFORMED]


 - #4424 unknown section "global"
 - #4482 unknown section "e2e"

## Know before update


 - An alert will be generated for each instance of an object with a deprecated `extended-monitoring.flant.com`. Change it to `extended-monitoring.deckhouse.io` ASAP.
 - Control plane components and kubelet will restart.

## Features


 - **[admission-policy-engine]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[ceph-csi]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cert-manager]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[chrony]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cloud-provider-aws]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cloud-provider-azure]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cloud-provider-gcp]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cloud-provider-openstack]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cloud-provider-openstack]** Add cloud data discoverer service which get information about available instance types for node groups. [#4187](https://github.com/deckhouse/deckhouse/pull/4187)
 - **[cloud-provider-vsphere]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cloud-provider-yandex]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cni-cilium]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cni-cilium]** Enable external access to ClusterIP services [#4302](https://github.com/deckhouse/deckhouse/pull/4302)
    cilium pods should be restarted
 - **[cni-flannel]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[cni-simple-bridge]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[containerized-data-importer]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[dashboard]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[deckhouse]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[delivery]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[descheduler]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[flant-integration]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[ingress-nginx]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[istio]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[keepalived]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[linstor]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[local-path-provisioner]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[log-shipper]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[metallb]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[network-gateway]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[node-manager]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[node-manager]** Add annotation `update.node.deckhouse.io/draining=user` for starting node drain process [#4310](https://github.com/deckhouse/deckhouse/pull/4310)
 - **[okmeter]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[openvpn]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[operator-prometheus]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[operator-trivy]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[operator-trivy]** Added CIS Benchmark reports and dashboard. [#3995](https://github.com/deckhouse/deckhouse/pull/3995)
 - **[operator-trivy]** Added `NodeRestriction` admission plugin and turned on `RotateKubeletServerCertificate` via feature gate. [#3995](https://github.com/deckhouse/deckhouse/pull/3995)
    Control plane components and kubelet will restart.
 - **[pod-reloader]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[prometheus]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[runtime-audit-engine]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[snapshot-controller]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[upmeter]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[user-authn]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[user-authz]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)
 - **[virtualization]** Add support for `PrometheusRule`. [#4407](https://github.com/deckhouse/deckhouse/pull/4407)

## Fixes


 - **[node-manager]** Fix error node group condition [#4367](https://github.com/deckhouse/deckhouse/pull/4367)
 - **[operator-trivy]** Add support for kubernetes.io/dockercfg secrets in imagePullSecrets pods field for scan jobs. [#4469](https://github.com/deckhouse/deckhouse/pull/4469)
 - **[operator-trivy]** Fixed k8s file permissions. [#3995](https://github.com/deckhouse/deckhouse/pull/3995)
 - **[prometheus]** Fixed creation of multiple CustomAlertmanager resources. [#4402](https://github.com/deckhouse/deckhouse/pull/4402)
 - **[prometheus]** Update Prometheus to `2.43.0` (bug and security fixes, performance improvements). [#4269](https://github.com/deckhouse/deckhouse/pull/4269)

## Chore


 - **[candi]** refactoring control-plane-manager image in golang [#4237](https://github.com/deckhouse/deckhouse/pull/4237)
    control-plane-manager will restart
 - **[ceph-csi]** volumeBindingMode changed from default to WaitForFirstConsumer [#3974](https://github.com/deckhouse/deckhouse/pull/3974)
 - **[deckhouse-controller]** Bump addon-operator version. [#4425](https://github.com/deckhouse/deckhouse/pull/4425)
 - **[docs]** Added the guide about how to go to the production environment with Deckhouse. [#3831](https://github.com/deckhouse/deckhouse/pull/3831)
 - **[extended-monitoring]** Starting the process of migration from extended-monitoring annotations to labels [#4356](https://github.com/deckhouse/deckhouse/pull/4356)
    An alert will be generated for each instance of an object with a deprecated `extended-monitoring.flant.com`. Change it to `extended-monitoring.deckhouse.io` ASAP.
 - **[linstor]** Adjust timers and timeouts for more stability in non-stable networks [#4463](https://github.com/deckhouse/deckhouse/pull/4463)
 - **[operator-trivy]** Update operator-trivy version to `v0.13.1` and trivy version to `v0.40.0`. [#4465](https://github.com/deckhouse/deckhouse/pull/4465)

