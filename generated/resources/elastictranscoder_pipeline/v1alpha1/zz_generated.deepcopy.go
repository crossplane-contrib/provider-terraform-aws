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
func (in *ContentConfig) DeepCopyInto(out *ContentConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfig.
func (in *ContentConfig) DeepCopy() *ContentConfig {
	if in == nil {
		return nil
	}
	out := new(ContentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentConfigPermissions) DeepCopyInto(out *ContentConfigPermissions) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentConfigPermissions.
func (in *ContentConfigPermissions) DeepCopy() *ContentConfigPermissions {
	if in == nil {
		return nil
	}
	out := new(ContentConfigPermissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipeline) DeepCopyInto(out *ElastictranscoderPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipeline.
func (in *ElastictranscoderPipeline) DeepCopy() *ElastictranscoderPipeline {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElastictranscoderPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineList) DeepCopyInto(out *ElastictranscoderPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElastictranscoderPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineList.
func (in *ElastictranscoderPipelineList) DeepCopy() *ElastictranscoderPipelineList {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElastictranscoderPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineObservation) DeepCopyInto(out *ElastictranscoderPipelineObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineObservation.
func (in *ElastictranscoderPipelineObservation) DeepCopy() *ElastictranscoderPipelineObservation {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineParameters) DeepCopyInto(out *ElastictranscoderPipelineParameters) {
	*out = *in
	out.ContentConfig = in.ContentConfig
	in.ContentConfigPermissions.DeepCopyInto(&out.ContentConfigPermissions)
	out.Notifications = in.Notifications
	out.ThumbnailConfig = in.ThumbnailConfig
	in.ThumbnailConfigPermissions.DeepCopyInto(&out.ThumbnailConfigPermissions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineParameters.
func (in *ElastictranscoderPipelineParameters) DeepCopy() *ElastictranscoderPipelineParameters {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineSpec) DeepCopyInto(out *ElastictranscoderPipelineSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineSpec.
func (in *ElastictranscoderPipelineSpec) DeepCopy() *ElastictranscoderPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastictranscoderPipelineStatus) DeepCopyInto(out *ElastictranscoderPipelineStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastictranscoderPipelineStatus.
func (in *ElastictranscoderPipelineStatus) DeepCopy() *ElastictranscoderPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(ElastictranscoderPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notifications) DeepCopyInto(out *Notifications) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notifications.
func (in *Notifications) DeepCopy() *Notifications {
	if in == nil {
		return nil
	}
	out := new(Notifications)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfig) DeepCopyInto(out *ThumbnailConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfig.
func (in *ThumbnailConfig) DeepCopy() *ThumbnailConfig {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThumbnailConfigPermissions) DeepCopyInto(out *ThumbnailConfigPermissions) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThumbnailConfigPermissions.
func (in *ThumbnailConfigPermissions) DeepCopy() *ThumbnailConfigPermissions {
	if in == nil {
		return nil
	}
	out := new(ThumbnailConfigPermissions)
	in.DeepCopyInto(out)
	return out
}