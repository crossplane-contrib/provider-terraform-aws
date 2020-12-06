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
func (in *CacheAttributes) DeepCopyInto(out *CacheAttributes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheAttributes.
func (in *CacheAttributes) DeepCopy() *CacheAttributes {
	if in == nil {
		return nil
	}
	out := new(CacheAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewaySmbFileShare) DeepCopyInto(out *StoragegatewaySmbFileShare) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewaySmbFileShare.
func (in *StoragegatewaySmbFileShare) DeepCopy() *StoragegatewaySmbFileShare {
	if in == nil {
		return nil
	}
	out := new(StoragegatewaySmbFileShare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragegatewaySmbFileShare) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewaySmbFileShareList) DeepCopyInto(out *StoragegatewaySmbFileShareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoragegatewaySmbFileShare, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewaySmbFileShareList.
func (in *StoragegatewaySmbFileShareList) DeepCopy() *StoragegatewaySmbFileShareList {
	if in == nil {
		return nil
	}
	out := new(StoragegatewaySmbFileShareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragegatewaySmbFileShareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewaySmbFileShareObservation) DeepCopyInto(out *StoragegatewaySmbFileShareObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewaySmbFileShareObservation.
func (in *StoragegatewaySmbFileShareObservation) DeepCopy() *StoragegatewaySmbFileShareObservation {
	if in == nil {
		return nil
	}
	out := new(StoragegatewaySmbFileShareObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewaySmbFileShareParameters) DeepCopyInto(out *StoragegatewaySmbFileShareParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ValidUserList != nil {
		in, out := &in.ValidUserList, &out.ValidUserList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdminUserList != nil {
		in, out := &in.AdminUserList, &out.AdminUserList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InvalidUserList != nil {
		in, out := &in.InvalidUserList, &out.InvalidUserList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Timeouts = in.Timeouts
	out.CacheAttributes = in.CacheAttributes
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewaySmbFileShareParameters.
func (in *StoragegatewaySmbFileShareParameters) DeepCopy() *StoragegatewaySmbFileShareParameters {
	if in == nil {
		return nil
	}
	out := new(StoragegatewaySmbFileShareParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewaySmbFileShareSpec) DeepCopyInto(out *StoragegatewaySmbFileShareSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewaySmbFileShareSpec.
func (in *StoragegatewaySmbFileShareSpec) DeepCopy() *StoragegatewaySmbFileShareSpec {
	if in == nil {
		return nil
	}
	out := new(StoragegatewaySmbFileShareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewaySmbFileShareStatus) DeepCopyInto(out *StoragegatewaySmbFileShareStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewaySmbFileShareStatus.
func (in *StoragegatewaySmbFileShareStatus) DeepCopy() *StoragegatewaySmbFileShareStatus {
	if in == nil {
		return nil
	}
	out := new(StoragegatewaySmbFileShareStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeouts) DeepCopyInto(out *Timeouts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeouts.
func (in *Timeouts) DeepCopy() *Timeouts {
	if in == nil {
		return nil
	}
	out := new(Timeouts)
	in.DeepCopyInto(out)
	return out
}
