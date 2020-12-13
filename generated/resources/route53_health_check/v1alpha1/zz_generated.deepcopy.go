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
func (in *Route53HealthCheck) DeepCopyInto(out *Route53HealthCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53HealthCheck.
func (in *Route53HealthCheck) DeepCopy() *Route53HealthCheck {
	if in == nil {
		return nil
	}
	out := new(Route53HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route53HealthCheck) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53HealthCheckList) DeepCopyInto(out *Route53HealthCheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route53HealthCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53HealthCheckList.
func (in *Route53HealthCheckList) DeepCopy() *Route53HealthCheckList {
	if in == nil {
		return nil
	}
	out := new(Route53HealthCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route53HealthCheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53HealthCheckObservation) DeepCopyInto(out *Route53HealthCheckObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53HealthCheckObservation.
func (in *Route53HealthCheckObservation) DeepCopy() *Route53HealthCheckObservation {
	if in == nil {
		return nil
	}
	out := new(Route53HealthCheckObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53HealthCheckParameters) DeepCopyInto(out *Route53HealthCheckParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ChildHealthchecks != nil {
		in, out := &in.ChildHealthchecks, &out.ChildHealthchecks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53HealthCheckParameters.
func (in *Route53HealthCheckParameters) DeepCopy() *Route53HealthCheckParameters {
	if in == nil {
		return nil
	}
	out := new(Route53HealthCheckParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53HealthCheckSpec) DeepCopyInto(out *Route53HealthCheckSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53HealthCheckSpec.
func (in *Route53HealthCheckSpec) DeepCopy() *Route53HealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(Route53HealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53HealthCheckStatus) DeepCopyInto(out *Route53HealthCheckStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53HealthCheckStatus.
func (in *Route53HealthCheckStatus) DeepCopy() *Route53HealthCheckStatus {
	if in == nil {
		return nil
	}
	out := new(Route53HealthCheckStatus)
	in.DeepCopyInto(out)
	return out
}
