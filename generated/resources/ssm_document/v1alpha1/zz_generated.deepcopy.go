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
func (in *AttachmentsSource) DeepCopyInto(out *AttachmentsSource) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentsSource.
func (in *AttachmentsSource) DeepCopy() *AttachmentsSource {
	if in == nil {
		return nil
	}
	out := new(AttachmentsSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmDocument) DeepCopyInto(out *SsmDocument) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmDocument.
func (in *SsmDocument) DeepCopy() *SsmDocument {
	if in == nil {
		return nil
	}
	out := new(SsmDocument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SsmDocument) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmDocumentList) DeepCopyInto(out *SsmDocumentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SsmDocument, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmDocumentList.
func (in *SsmDocumentList) DeepCopy() *SsmDocumentList {
	if in == nil {
		return nil
	}
	out := new(SsmDocumentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SsmDocumentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmDocumentObservation) DeepCopyInto(out *SsmDocumentObservation) {
	*out = *in
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = make([]Parameter, len(*in))
		copy(*out, *in)
	}
	if in.PlatformTypes != nil {
		in, out := &in.PlatformTypes, &out.PlatformTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmDocumentObservation.
func (in *SsmDocumentObservation) DeepCopy() *SsmDocumentObservation {
	if in == nil {
		return nil
	}
	out := new(SsmDocumentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmDocumentParameters) DeepCopyInto(out *SsmDocumentParameters) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AttachmentsSource.DeepCopyInto(&out.AttachmentsSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmDocumentParameters.
func (in *SsmDocumentParameters) DeepCopy() *SsmDocumentParameters {
	if in == nil {
		return nil
	}
	out := new(SsmDocumentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmDocumentSpec) DeepCopyInto(out *SsmDocumentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmDocumentSpec.
func (in *SsmDocumentSpec) DeepCopy() *SsmDocumentSpec {
	if in == nil {
		return nil
	}
	out := new(SsmDocumentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmDocumentStatus) DeepCopyInto(out *SsmDocumentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmDocumentStatus.
func (in *SsmDocumentStatus) DeepCopy() *SsmDocumentStatus {
	if in == nil {
		return nil
	}
	out := new(SsmDocumentStatus)
	in.DeepCopyInto(out)
	return out
}
