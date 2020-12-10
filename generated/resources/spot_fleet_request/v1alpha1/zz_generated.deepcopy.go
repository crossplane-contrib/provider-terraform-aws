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
func (in *EbsBlockDevice) DeepCopyInto(out *EbsBlockDevice) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsBlockDevice.
func (in *EbsBlockDevice) DeepCopy() *EbsBlockDevice {
	if in == nil {
		return nil
	}
	out := new(EbsBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralBlockDevice) DeepCopyInto(out *EphemeralBlockDevice) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralBlockDevice.
func (in *EphemeralBlockDevice) DeepCopy() *EphemeralBlockDevice {
	if in == nil {
		return nil
	}
	out := new(EphemeralBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchSpecification) DeepCopyInto(out *LaunchSpecification) {
	*out = *in
	if in.VpcSecurityGroupIds != nil {
		in, out := &in.VpcSecurityGroupIds, &out.VpcSecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.EphemeralBlockDevice = in.EphemeralBlockDevice
	out.RootBlockDevice = in.RootBlockDevice
	out.EbsBlockDevice = in.EbsBlockDevice
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchSpecification.
func (in *LaunchSpecification) DeepCopy() *LaunchSpecification {
	if in == nil {
		return nil
	}
	out := new(LaunchSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateConfig) DeepCopyInto(out *LaunchTemplateConfig) {
	*out = *in
	out.Overrides = in.Overrides
	out.LaunchTemplateSpecification = in.LaunchTemplateSpecification
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateConfig.
func (in *LaunchTemplateConfig) DeepCopy() *LaunchTemplateConfig {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateSpecification) DeepCopyInto(out *LaunchTemplateSpecification) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateSpecification.
func (in *LaunchTemplateSpecification) DeepCopy() *LaunchTemplateSpecification {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Overrides) DeepCopyInto(out *Overrides) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Overrides.
func (in *Overrides) DeepCopy() *Overrides {
	if in == nil {
		return nil
	}
	out := new(Overrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootBlockDevice) DeepCopyInto(out *RootBlockDevice) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootBlockDevice.
func (in *RootBlockDevice) DeepCopy() *RootBlockDevice {
	if in == nil {
		return nil
	}
	out := new(RootBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotFleetRequest) DeepCopyInto(out *SpotFleetRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotFleetRequest.
func (in *SpotFleetRequest) DeepCopy() *SpotFleetRequest {
	if in == nil {
		return nil
	}
	out := new(SpotFleetRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpotFleetRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotFleetRequestList) DeepCopyInto(out *SpotFleetRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpotFleetRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotFleetRequestList.
func (in *SpotFleetRequestList) DeepCopy() *SpotFleetRequestList {
	if in == nil {
		return nil
	}
	out := new(SpotFleetRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpotFleetRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotFleetRequestObservation) DeepCopyInto(out *SpotFleetRequestObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotFleetRequestObservation.
func (in *SpotFleetRequestObservation) DeepCopy() *SpotFleetRequestObservation {
	if in == nil {
		return nil
	}
	out := new(SpotFleetRequestObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotFleetRequestParameters) DeepCopyInto(out *SpotFleetRequestParameters) {
	*out = *in
	if in.LoadBalancers != nil {
		in, out := &in.LoadBalancers, &out.LoadBalancers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetGroupArns != nil {
		in, out := &in.TargetGroupArns, &out.TargetGroupArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.LaunchSpecification.DeepCopyInto(&out.LaunchSpecification)
	out.LaunchTemplateConfig = in.LaunchTemplateConfig
	out.Timeouts = in.Timeouts
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotFleetRequestParameters.
func (in *SpotFleetRequestParameters) DeepCopy() *SpotFleetRequestParameters {
	if in == nil {
		return nil
	}
	out := new(SpotFleetRequestParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotFleetRequestSpec) DeepCopyInto(out *SpotFleetRequestSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotFleetRequestSpec.
func (in *SpotFleetRequestSpec) DeepCopy() *SpotFleetRequestSpec {
	if in == nil {
		return nil
	}
	out := new(SpotFleetRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotFleetRequestStatus) DeepCopyInto(out *SpotFleetRequestStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotFleetRequestStatus.
func (in *SpotFleetRequestStatus) DeepCopy() *SpotFleetRequestStatus {
	if in == nil {
		return nil
	}
	out := new(SpotFleetRequestStatus)
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
