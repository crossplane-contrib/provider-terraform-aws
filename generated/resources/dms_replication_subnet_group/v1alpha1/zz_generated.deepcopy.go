// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DmsReplicationSubnetGroup) DeepCopyInto(out *DmsReplicationSubnetGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DmsReplicationSubnetGroup.
func (in *DmsReplicationSubnetGroup) DeepCopy() *DmsReplicationSubnetGroup {
	if in == nil {
		return nil
	}
	out := new(DmsReplicationSubnetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DmsReplicationSubnetGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DmsReplicationSubnetGroupList) DeepCopyInto(out *DmsReplicationSubnetGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DmsReplicationSubnetGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DmsReplicationSubnetGroupList.
func (in *DmsReplicationSubnetGroupList) DeepCopy() *DmsReplicationSubnetGroupList {
	if in == nil {
		return nil
	}
	out := new(DmsReplicationSubnetGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DmsReplicationSubnetGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DmsReplicationSubnetGroupObservation) DeepCopyInto(out *DmsReplicationSubnetGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DmsReplicationSubnetGroupObservation.
func (in *DmsReplicationSubnetGroupObservation) DeepCopy() *DmsReplicationSubnetGroupObservation {
	if in == nil {
		return nil
	}
	out := new(DmsReplicationSubnetGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DmsReplicationSubnetGroupParameters) DeepCopyInto(out *DmsReplicationSubnetGroupParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DmsReplicationSubnetGroupParameters.
func (in *DmsReplicationSubnetGroupParameters) DeepCopy() *DmsReplicationSubnetGroupParameters {
	if in == nil {
		return nil
	}
	out := new(DmsReplicationSubnetGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DmsReplicationSubnetGroupSpec) DeepCopyInto(out *DmsReplicationSubnetGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DmsReplicationSubnetGroupSpec.
func (in *DmsReplicationSubnetGroupSpec) DeepCopy() *DmsReplicationSubnetGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DmsReplicationSubnetGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DmsReplicationSubnetGroupStatus) DeepCopyInto(out *DmsReplicationSubnetGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DmsReplicationSubnetGroupStatus.
func (in *DmsReplicationSubnetGroupStatus) DeepCopy() *DmsReplicationSubnetGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DmsReplicationSubnetGroupStatus)
	in.DeepCopyInto(out)
	return out
}
