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
func (in *AcmpcaCertificateAuthority) DeepCopyInto(out *AcmpcaCertificateAuthority) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthority.
func (in *AcmpcaCertificateAuthority) DeepCopy() *AcmpcaCertificateAuthority {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmpcaCertificateAuthority) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityList) DeepCopyInto(out *AcmpcaCertificateAuthorityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AcmpcaCertificateAuthority, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityList.
func (in *AcmpcaCertificateAuthorityList) DeepCopy() *AcmpcaCertificateAuthorityList {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcmpcaCertificateAuthorityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityObservation) DeepCopyInto(out *AcmpcaCertificateAuthorityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityObservation.
func (in *AcmpcaCertificateAuthorityObservation) DeepCopy() *AcmpcaCertificateAuthorityObservation {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityParameters) DeepCopyInto(out *AcmpcaCertificateAuthorityParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Timeouts = in.Timeouts
	out.CertificateAuthorityConfiguration = in.CertificateAuthorityConfiguration
	out.RevocationConfiguration = in.RevocationConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityParameters.
func (in *AcmpcaCertificateAuthorityParameters) DeepCopy() *AcmpcaCertificateAuthorityParameters {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthoritySpec) DeepCopyInto(out *AcmpcaCertificateAuthoritySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthoritySpec.
func (in *AcmpcaCertificateAuthoritySpec) DeepCopy() *AcmpcaCertificateAuthoritySpec {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthoritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcmpcaCertificateAuthorityStatus) DeepCopyInto(out *AcmpcaCertificateAuthorityStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcmpcaCertificateAuthorityStatus.
func (in *AcmpcaCertificateAuthorityStatus) DeepCopy() *AcmpcaCertificateAuthorityStatus {
	if in == nil {
		return nil
	}
	out := new(AcmpcaCertificateAuthorityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityConfiguration) DeepCopyInto(out *CertificateAuthorityConfiguration) {
	*out = *in
	out.Subject = in.Subject
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityConfiguration.
func (in *CertificateAuthorityConfiguration) DeepCopy() *CertificateAuthorityConfiguration {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrlConfiguration) DeepCopyInto(out *CrlConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrlConfiguration.
func (in *CrlConfiguration) DeepCopy() *CrlConfiguration {
	if in == nil {
		return nil
	}
	out := new(CrlConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevocationConfiguration) DeepCopyInto(out *RevocationConfiguration) {
	*out = *in
	out.CrlConfiguration = in.CrlConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationConfiguration.
func (in *RevocationConfiguration) DeepCopy() *RevocationConfiguration {
	if in == nil {
		return nil
	}
	out := new(RevocationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subject.
func (in *Subject) DeepCopy() *Subject {
	if in == nil {
		return nil
	}
	out := new(Subject)
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
