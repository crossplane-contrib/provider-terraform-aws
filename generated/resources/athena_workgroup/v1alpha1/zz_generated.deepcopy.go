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
func (in *AthenaWorkgroup) DeepCopyInto(out *AthenaWorkgroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AthenaWorkgroup.
func (in *AthenaWorkgroup) DeepCopy() *AthenaWorkgroup {
	if in == nil {
		return nil
	}
	out := new(AthenaWorkgroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AthenaWorkgroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AthenaWorkgroupList) DeepCopyInto(out *AthenaWorkgroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AthenaWorkgroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AthenaWorkgroupList.
func (in *AthenaWorkgroupList) DeepCopy() *AthenaWorkgroupList {
	if in == nil {
		return nil
	}
	out := new(AthenaWorkgroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AthenaWorkgroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AthenaWorkgroupObservation) DeepCopyInto(out *AthenaWorkgroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AthenaWorkgroupObservation.
func (in *AthenaWorkgroupObservation) DeepCopy() *AthenaWorkgroupObservation {
	if in == nil {
		return nil
	}
	out := new(AthenaWorkgroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AthenaWorkgroupParameters) DeepCopyInto(out *AthenaWorkgroupParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Configuration = in.Configuration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AthenaWorkgroupParameters.
func (in *AthenaWorkgroupParameters) DeepCopy() *AthenaWorkgroupParameters {
	if in == nil {
		return nil
	}
	out := new(AthenaWorkgroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AthenaWorkgroupSpec) DeepCopyInto(out *AthenaWorkgroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AthenaWorkgroupSpec.
func (in *AthenaWorkgroupSpec) DeepCopy() *AthenaWorkgroupSpec {
	if in == nil {
		return nil
	}
	out := new(AthenaWorkgroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AthenaWorkgroupStatus) DeepCopyInto(out *AthenaWorkgroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AthenaWorkgroupStatus.
func (in *AthenaWorkgroupStatus) DeepCopy() *AthenaWorkgroupStatus {
	if in == nil {
		return nil
	}
	out := new(AthenaWorkgroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.ResultConfiguration = in.ResultConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfiguration) DeepCopyInto(out *EncryptionConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfiguration.
func (in *EncryptionConfiguration) DeepCopy() *EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultConfiguration) DeepCopyInto(out *ResultConfiguration) {
	*out = *in
	out.EncryptionConfiguration = in.EncryptionConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultConfiguration.
func (in *ResultConfiguration) DeepCopy() *ResultConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResultConfiguration)
	in.DeepCopyInto(out)
	return out
}