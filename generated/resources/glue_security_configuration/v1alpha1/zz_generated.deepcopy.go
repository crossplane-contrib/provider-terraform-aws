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
func (in *CloudwatchEncryption) DeepCopyInto(out *CloudwatchEncryption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchEncryption.
func (in *CloudwatchEncryption) DeepCopy() *CloudwatchEncryption {
	if in == nil {
		return nil
	}
	out := new(CloudwatchEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfiguration) DeepCopyInto(out *EncryptionConfiguration) {
	*out = *in
	out.CloudwatchEncryption = in.CloudwatchEncryption
	out.JobBookmarksEncryption = in.JobBookmarksEncryption
	out.S3Encryption = in.S3Encryption
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
func (in *GlueSecurityConfiguration) DeepCopyInto(out *GlueSecurityConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueSecurityConfiguration.
func (in *GlueSecurityConfiguration) DeepCopy() *GlueSecurityConfiguration {
	if in == nil {
		return nil
	}
	out := new(GlueSecurityConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueSecurityConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueSecurityConfigurationList) DeepCopyInto(out *GlueSecurityConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlueSecurityConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueSecurityConfigurationList.
func (in *GlueSecurityConfigurationList) DeepCopy() *GlueSecurityConfigurationList {
	if in == nil {
		return nil
	}
	out := new(GlueSecurityConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueSecurityConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueSecurityConfigurationObservation) DeepCopyInto(out *GlueSecurityConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueSecurityConfigurationObservation.
func (in *GlueSecurityConfigurationObservation) DeepCopy() *GlueSecurityConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(GlueSecurityConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueSecurityConfigurationParameters) DeepCopyInto(out *GlueSecurityConfigurationParameters) {
	*out = *in
	out.EncryptionConfiguration = in.EncryptionConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueSecurityConfigurationParameters.
func (in *GlueSecurityConfigurationParameters) DeepCopy() *GlueSecurityConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(GlueSecurityConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueSecurityConfigurationSpec) DeepCopyInto(out *GlueSecurityConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueSecurityConfigurationSpec.
func (in *GlueSecurityConfigurationSpec) DeepCopy() *GlueSecurityConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(GlueSecurityConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueSecurityConfigurationStatus) DeepCopyInto(out *GlueSecurityConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueSecurityConfigurationStatus.
func (in *GlueSecurityConfigurationStatus) DeepCopy() *GlueSecurityConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(GlueSecurityConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobBookmarksEncryption) DeepCopyInto(out *JobBookmarksEncryption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobBookmarksEncryption.
func (in *JobBookmarksEncryption) DeepCopy() *JobBookmarksEncryption {
	if in == nil {
		return nil
	}
	out := new(JobBookmarksEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Encryption) DeepCopyInto(out *S3Encryption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Encryption.
func (in *S3Encryption) DeepCopy() *S3Encryption {
	if in == nil {
		return nil
	}
	out := new(S3Encryption)
	in.DeepCopyInto(out)
	return out
}