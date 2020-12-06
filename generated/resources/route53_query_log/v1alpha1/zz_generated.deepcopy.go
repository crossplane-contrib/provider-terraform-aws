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
func (in *Route53QueryLog) DeepCopyInto(out *Route53QueryLog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53QueryLog.
func (in *Route53QueryLog) DeepCopy() *Route53QueryLog {
	if in == nil {
		return nil
	}
	out := new(Route53QueryLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route53QueryLog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53QueryLogList) DeepCopyInto(out *Route53QueryLogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route53QueryLog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53QueryLogList.
func (in *Route53QueryLogList) DeepCopy() *Route53QueryLogList {
	if in == nil {
		return nil
	}
	out := new(Route53QueryLogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route53QueryLogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53QueryLogObservation) DeepCopyInto(out *Route53QueryLogObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53QueryLogObservation.
func (in *Route53QueryLogObservation) DeepCopy() *Route53QueryLogObservation {
	if in == nil {
		return nil
	}
	out := new(Route53QueryLogObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53QueryLogParameters) DeepCopyInto(out *Route53QueryLogParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53QueryLogParameters.
func (in *Route53QueryLogParameters) DeepCopy() *Route53QueryLogParameters {
	if in == nil {
		return nil
	}
	out := new(Route53QueryLogParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53QueryLogSpec) DeepCopyInto(out *Route53QueryLogSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53QueryLogSpec.
func (in *Route53QueryLogSpec) DeepCopy() *Route53QueryLogSpec {
	if in == nil {
		return nil
	}
	out := new(Route53QueryLogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53QueryLogStatus) DeepCopyInto(out *Route53QueryLogStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53QueryLogStatus.
func (in *Route53QueryLogStatus) DeepCopy() *Route53QueryLogStatus {
	if in == nil {
		return nil
	}
	out := new(Route53QueryLogStatus)
	in.DeepCopyInto(out)
	return out
}
