//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 Flant JSC

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomaticDisruptions) DeepCopyInto(out *AutomaticDisruptions) {
	*out = *in
	if in.DrainBeforeApproval != nil {
		in, out := &in.DrainBeforeApproval, &out.DrainBeforeApproval
		*out = new(bool)
		**out = **in
	}
	out.Windows = in.Windows.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomaticDisruptions.
func (in *AutomaticDisruptions) DeepCopy() *AutomaticDisruptions {
	if in == nil {
		return nil
	}
	out := new(AutomaticDisruptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRI) DeepCopyInto(out *CRI) {
	*out = *in
	if in.Containerd != nil {
		in, out := &in.Containerd, &out.Containerd
		*out = new(Containerd)
		(*in).DeepCopyInto(*out)
	}
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = new(Docker)
		(*in).DeepCopyInto(*out)
	}
	if in.NotManaged != nil {
		in, out := &in.NotManaged, &out.NotManaged
		*out = new(NotManaged)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRI.
func (in *CRI) DeepCopy() *CRI {
	if in == nil {
		return nil
	}
	out := new(CRI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chaos) DeepCopyInto(out *Chaos) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chaos.
func (in *Chaos) DeepCopy() *Chaos {
	if in == nil {
		return nil
	}
	out := new(Chaos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClassReference) DeepCopyInto(out *ClassReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClassReference.
func (in *ClassReference) DeepCopy() *ClassReference {
	if in == nil {
		return nil
	}
	out := new(ClassReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudInstances) DeepCopyInto(out *CloudInstances) {
	*out = *in
	if in.QuickShutdown != nil {
		in, out := &in.QuickShutdown, &out.QuickShutdown
		*out = new(bool)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MinPerZone != nil {
		in, out := &in.MinPerZone, &out.MinPerZone
		*out = new(int32)
		**out = **in
	}
	if in.MaxPerZone != nil {
		in, out := &in.MaxPerZone, &out.MaxPerZone
		*out = new(int32)
		**out = **in
	}
	if in.MaxUnavailablePerZone != nil {
		in, out := &in.MaxUnavailablePerZone, &out.MaxUnavailablePerZone
		*out = new(int32)
		**out = **in
	}
	if in.MaxSurgePerZone != nil {
		in, out := &in.MaxSurgePerZone, &out.MaxSurgePerZone
		*out = new(int32)
		**out = **in
	}
	if in.Standby != nil {
		in, out := &in.Standby, &out.Standby
		*out = new(intstr.IntOrString)
		**out = **in
	}
	in.StandbyHolder.DeepCopyInto(&out.StandbyHolder)
	out.ClassReference = in.ClassReference
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudInstances.
func (in *CloudInstances) DeepCopy() *CloudInstances {
	if in == nil {
		return nil
	}
	out := new(CloudInstances)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionSummary) DeepCopyInto(out *ConditionSummary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionSummary.
func (in *ConditionSummary) DeepCopy() *ConditionSummary {
	if in == nil {
		return nil
	}
	out := new(ConditionSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Containerd) DeepCopyInto(out *Containerd) {
	*out = *in
	if in.MaxConcurrentDownloads != nil {
		in, out := &in.MaxConcurrentDownloads, &out.MaxConcurrentDownloads
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Containerd.
func (in *Containerd) DeepCopy() *Containerd {
	if in == nil {
		return nil
	}
	out := new(Containerd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disruptions) DeepCopyInto(out *Disruptions) {
	*out = *in
	in.Automatic.DeepCopyInto(&out.Automatic)
	in.RollingUpdate.DeepCopyInto(&out.RollingUpdate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disruptions.
func (in *Disruptions) DeepCopy() *Disruptions {
	if in == nil {
		return nil
	}
	out := new(Disruptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Docker) DeepCopyInto(out *Docker) {
	*out = *in
	if in.MaxConcurrentDownloads != nil {
		in, out := &in.MaxConcurrentDownloads, &out.MaxConcurrentDownloads
		*out = new(int32)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Docker.
func (in *Docker) DeepCopy() *Docker {
	if in == nil {
		return nil
	}
	out := new(Docker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubelet) DeepCopyInto(out *Kubelet) {
	*out = *in
	if in.MaxPods != nil {
		in, out := &in.MaxPods, &out.MaxPods
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubelet.
func (in *Kubelet) DeepCopy() *Kubelet {
	if in == nil {
		return nil
	}
	out := new(Kubelet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineFailure) DeepCopyInto(out *MachineFailure) {
	*out = *in
	out.LastOperation = in.LastOperation
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineFailure.
func (in *MachineFailure) DeepCopy() *MachineFailure {
	if in == nil {
		return nil
	}
	out := new(MachineFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOperation) DeepCopyInto(out *MachineOperation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOperation.
func (in *MachineOperation) DeepCopy() *MachineOperation {
	if in == nil {
		return nil
	}
	out := new(MachineOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroup) DeepCopyInto(out *NodeGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroup.
func (in *NodeGroup) DeepCopy() *NodeGroup {
	if in == nil {
		return nil
	}
	out := new(NodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupCondition) DeepCopyInto(out *NodeGroupCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupCondition.
func (in *NodeGroupCondition) DeepCopy() *NodeGroupCondition {
	if in == nil {
		return nil
	}
	out := new(NodeGroupCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupSpec) DeepCopyInto(out *NodeGroupSpec) {
	*out = *in
	in.CRI.DeepCopyInto(&out.CRI)
	in.CloudInstances.DeepCopyInto(&out.CloudInstances)
	in.NodeTemplate.DeepCopyInto(&out.NodeTemplate)
	out.Chaos = in.Chaos
	in.OperatingSystem.DeepCopyInto(&out.OperatingSystem)
	in.Disruptions.DeepCopyInto(&out.Disruptions)
	in.Update.DeepCopyInto(&out.Update)
	in.Kubelet.DeepCopyInto(&out.Kubelet)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupSpec.
func (in *NodeGroupSpec) DeepCopy() *NodeGroupSpec {
	if in == nil {
		return nil
	}
	out := new(NodeGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupStatus) DeepCopyInto(out *NodeGroupStatus) {
	*out = *in
	if in.LastMachineFailures != nil {
		in, out := &in.LastMachineFailures, &out.LastMachineFailures
		*out = make([]MachineFailure, len(*in))
		copy(*out, *in)
	}
	out.ConditionSummary = in.ConditionSummary
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]NodeGroupCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupStatus.
func (in *NodeGroupStatus) DeepCopy() *NodeGroupStatus {
	if in == nil {
		return nil
	}
	out := new(NodeGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotManaged) DeepCopyInto(out *NotManaged) {
	*out = *in
	if in.CriSocketPath != nil {
		in, out := &in.CriSocketPath, &out.CriSocketPath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotManaged.
func (in *NotManaged) DeepCopy() *NotManaged {
	if in == nil {
		return nil
	}
	out := new(NotManaged)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatingSystem) DeepCopyInto(out *OperatingSystem) {
	*out = *in
	if in.ManageKernel != nil {
		in, out := &in.ManageKernel, &out.ManageKernel
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatingSystem.
func (in *OperatingSystem) DeepCopy() *OperatingSystem {
	if in == nil {
		return nil
	}
	out := new(OperatingSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.CPU = in.CPU
	out.Memory = in.Memory
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateDisruptions) DeepCopyInto(out *RollingUpdateDisruptions) {
	*out = *in
	out.Windows = in.Windows.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateDisruptions.
func (in *RollingUpdateDisruptions) DeepCopy() *RollingUpdateDisruptions {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateDisruptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyHolder) DeepCopyInto(out *StandbyHolder) {
	*out = *in
	if in.OverprovisioningRate != nil {
		in, out := &in.OverprovisioningRate, &out.OverprovisioningRate
		*out = new(int64)
		**out = **in
	}
	out.NotHeldResources = in.NotHeldResources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyHolder.
func (in *StandbyHolder) DeepCopy() *StandbyHolder {
	if in == nil {
		return nil
	}
	out := new(StandbyHolder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Update) DeepCopyInto(out *Update) {
	*out = *in
	if in.MaxConcurrent != nil {
		in, out := &in.MaxConcurrent, &out.MaxConcurrent
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Update.
func (in *Update) DeepCopy() *Update {
	if in == nil {
		return nil
	}
	out := new(Update)
	in.DeepCopyInto(out)
	return out
}
