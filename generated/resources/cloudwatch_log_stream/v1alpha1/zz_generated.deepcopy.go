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
func (in *CloudwatchLogStream) DeepCopyInto(out *CloudwatchLogStream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogStream.
func (in *CloudwatchLogStream) DeepCopy() *CloudwatchLogStream {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogStream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchLogStream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogStreamList) DeepCopyInto(out *CloudwatchLogStreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudwatchLogStream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogStreamList.
func (in *CloudwatchLogStreamList) DeepCopy() *CloudwatchLogStreamList {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogStreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchLogStreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogStreamObservation) DeepCopyInto(out *CloudwatchLogStreamObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogStreamObservation.
func (in *CloudwatchLogStreamObservation) DeepCopy() *CloudwatchLogStreamObservation {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogStreamObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogStreamParameters) DeepCopyInto(out *CloudwatchLogStreamParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogStreamParameters.
func (in *CloudwatchLogStreamParameters) DeepCopy() *CloudwatchLogStreamParameters {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogStreamParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogStreamSpec) DeepCopyInto(out *CloudwatchLogStreamSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogStreamSpec.
func (in *CloudwatchLogStreamSpec) DeepCopy() *CloudwatchLogStreamSpec {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogStreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogStreamStatus) DeepCopyInto(out *CloudwatchLogStreamStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogStreamStatus.
func (in *CloudwatchLogStreamStatus) DeepCopy() *CloudwatchLogStreamStatus {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogStreamStatus)
	in.DeepCopyInto(out)
	return out
}
