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
func (in *Configurations) DeepCopyInto(out *Configurations) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configurations.
func (in *Configurations) DeepCopy() *Configurations {
	if in == nil {
		return nil
	}
	out := new(Configurations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbsConfig) DeepCopyInto(out *EbsConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsConfig.
func (in *EbsConfig) DeepCopy() *EbsConfig {
	if in == nil {
		return nil
	}
	out := new(EbsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleet) DeepCopyInto(out *EmrInstanceFleet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleet.
func (in *EmrInstanceFleet) DeepCopy() *EmrInstanceFleet {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmrInstanceFleet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetList) DeepCopyInto(out *EmrInstanceFleetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EmrInstanceFleet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetList.
func (in *EmrInstanceFleetList) DeepCopy() *EmrInstanceFleetList {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmrInstanceFleetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetObservation) DeepCopyInto(out *EmrInstanceFleetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetObservation.
func (in *EmrInstanceFleetObservation) DeepCopy() *EmrInstanceFleetObservation {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetParameters) DeepCopyInto(out *EmrInstanceFleetParameters) {
	*out = *in
	in.InstanceTypeConfigs.DeepCopyInto(&out.InstanceTypeConfigs)
	out.LaunchSpecifications = in.LaunchSpecifications
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetParameters.
func (in *EmrInstanceFleetParameters) DeepCopy() *EmrInstanceFleetParameters {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetSpec) DeepCopyInto(out *EmrInstanceFleetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetSpec.
func (in *EmrInstanceFleetSpec) DeepCopy() *EmrInstanceFleetSpec {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetStatus) DeepCopyInto(out *EmrInstanceFleetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetStatus.
func (in *EmrInstanceFleetStatus) DeepCopy() *EmrInstanceFleetStatus {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceTypeConfigs) DeepCopyInto(out *InstanceTypeConfigs) {
	*out = *in
	in.Configurations.DeepCopyInto(&out.Configurations)
	out.EbsConfig = in.EbsConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceTypeConfigs.
func (in *InstanceTypeConfigs) DeepCopy() *InstanceTypeConfigs {
	if in == nil {
		return nil
	}
	out := new(InstanceTypeConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchSpecifications) DeepCopyInto(out *LaunchSpecifications) {
	*out = *in
	out.SpotSpecification = in.SpotSpecification
	out.OnDemandSpecification = in.OnDemandSpecification
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchSpecifications.
func (in *LaunchSpecifications) DeepCopy() *LaunchSpecifications {
	if in == nil {
		return nil
	}
	out := new(LaunchSpecifications)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnDemandSpecification) DeepCopyInto(out *OnDemandSpecification) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnDemandSpecification.
func (in *OnDemandSpecification) DeepCopy() *OnDemandSpecification {
	if in == nil {
		return nil
	}
	out := new(OnDemandSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotSpecification) DeepCopyInto(out *SpotSpecification) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotSpecification.
func (in *SpotSpecification) DeepCopy() *SpotSpecification {
	if in == nil {
		return nil
	}
	out := new(SpotSpecification)
	in.DeepCopyInto(out)
	return out
}