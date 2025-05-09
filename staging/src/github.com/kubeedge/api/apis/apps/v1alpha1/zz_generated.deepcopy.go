//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KubeEdge Authors.

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandArgsOverrider) DeepCopyInto(out *CommandArgsOverrider) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandArgsOverrider.
func (in *CommandArgsOverrider) DeepCopy() *CommandArgsOverrider {
	if in == nil {
		return nil
	}
	out := new(CommandArgsOverrider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeApplication) DeepCopyInto(out *EdgeApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeApplication.
func (in *EdgeApplication) DeepCopy() *EdgeApplication {
	if in == nil {
		return nil
	}
	out := new(EdgeApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeApplicationList) DeepCopyInto(out *EdgeApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeApplicationList.
func (in *EdgeApplicationList) DeepCopy() *EdgeApplicationList {
	if in == nil {
		return nil
	}
	out := new(EdgeApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeApplicationSpec) DeepCopyInto(out *EdgeApplicationSpec) {
	*out = *in
	in.WorkloadTemplate.DeepCopyInto(&out.WorkloadTemplate)
	in.WorkloadScope.DeepCopyInto(&out.WorkloadScope)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeApplicationSpec.
func (in *EdgeApplicationSpec) DeepCopy() *EdgeApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeApplicationStatus) DeepCopyInto(out *EdgeApplicationStatus) {
	*out = *in
	if in.WorkloadStatus != nil {
		in, out := &in.WorkloadStatus, &out.WorkloadStatus
		*out = make([]ManifestStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeApplicationStatus.
func (in *EdgeApplicationStatus) DeepCopy() *EdgeApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvOverrider) DeepCopyInto(out *EnvOverrider) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvOverrider.
func (in *EnvOverrider) DeepCopy() *EnvOverrider {
	if in == nil {
		return nil
	}
	out := new(EnvOverrider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageOverrider) DeepCopyInto(out *ImageOverrider) {
	*out = *in
	if in.Predicate != nil {
		in, out := &in.Predicate, &out.Predicate
		*out = new(ImagePredicate)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOverrider.
func (in *ImageOverrider) DeepCopy() *ImageOverrider {
	if in == nil {
		return nil
	}
	out := new(ImageOverrider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePredicate) DeepCopyInto(out *ImagePredicate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePredicate.
func (in *ImagePredicate) DeepCopy() *ImagePredicate {
	if in == nil {
		return nil
	}
	out := new(ImagePredicate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Manifest) DeepCopyInto(out *Manifest) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Manifest.
func (in *Manifest) DeepCopy() *Manifest {
	if in == nil {
		return nil
	}
	out := new(Manifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestStatus) DeepCopyInto(out *ManifestStatus) {
	*out = *in
	out.Identifier = in.Identifier
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestStatus.
func (in *ManifestStatus) DeepCopy() *ManifestStatus {
	if in == nil {
		return nil
	}
	out := new(ManifestStatus)
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

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupList) DeepCopyInto(out *NodeGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupList.
func (in *NodeGroupList) DeepCopy() *NodeGroupList {
	if in == nil {
		return nil
	}
	out := new(NodeGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupSpec) DeepCopyInto(out *NodeGroupSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
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
	if in.NodeStatuses != nil {
		in, out := &in.NodeStatuses, &out.NodeStatuses
		*out = make([]NodeStatus, len(*in))
		copy(*out, *in)
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
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Overriders) DeepCopyInto(out *Overriders) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	if in.ImageOverriders != nil {
		in, out := &in.ImageOverriders, &out.ImageOverriders
		*out = make([]ImageOverrider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvOverriders != nil {
		in, out := &in.EnvOverriders, &out.EnvOverriders
		*out = make([]EnvOverrider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CommandOverriders != nil {
		in, out := &in.CommandOverriders, &out.CommandOverriders
		*out = make([]CommandArgsOverrider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ArgsOverriders != nil {
		in, out := &in.ArgsOverriders, &out.ArgsOverriders
		*out = make([]CommandArgsOverrider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourcesOverriders != nil {
		in, out := &in.ResourcesOverriders, &out.ResourcesOverriders
		*out = make([]ResourcesOverrider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Overriders.
func (in *Overriders) DeepCopy() *Overriders {
	if in == nil {
		return nil
	}
	out := new(Overriders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceIdentifier) DeepCopyInto(out *ResourceIdentifier) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceIdentifier.
func (in *ResourceIdentifier) DeepCopy() *ResourceIdentifier {
	if in == nil {
		return nil
	}
	out := new(ResourceIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTemplate) DeepCopyInto(out *ResourceTemplate) {
	*out = *in
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = make([]Manifest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTemplate.
func (in *ResourceTemplate) DeepCopy() *ResourceTemplate {
	if in == nil {
		return nil
	}
	out := new(ResourceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesOverrider) DeepCopyInto(out *ResourcesOverrider) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesOverrider.
func (in *ResourcesOverrider) DeepCopy() *ResourcesOverrider {
	if in == nil {
		return nil
	}
	out := new(ResourcesOverrider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetNodeGroup) DeepCopyInto(out *TargetNodeGroup) {
	*out = *in
	in.Overriders.DeepCopyInto(&out.Overriders)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetNodeGroup.
func (in *TargetNodeGroup) DeepCopy() *TargetNodeGroup {
	if in == nil {
		return nil
	}
	out := new(TargetNodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetNodeLabel) DeepCopyInto(out *TargetNodeLabel) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	in.Overriders.DeepCopyInto(&out.Overriders)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetNodeLabel.
func (in *TargetNodeLabel) DeepCopy() *TargetNodeLabel {
	if in == nil {
		return nil
	}
	out := new(TargetNodeLabel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadScope) DeepCopyInto(out *WorkloadScope) {
	*out = *in
	if in.TargetNodeGroups != nil {
		in, out := &in.TargetNodeGroups, &out.TargetNodeGroups
		*out = make([]TargetNodeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetNodeLabels != nil {
		in, out := &in.TargetNodeLabels, &out.TargetNodeLabels
		*out = make([]TargetNodeLabel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadScope.
func (in *WorkloadScope) DeepCopy() *WorkloadScope {
	if in == nil {
		return nil
	}
	out := new(WorkloadScope)
	in.DeepCopyInto(out)
	return out
}
