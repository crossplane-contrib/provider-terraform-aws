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
func (in *AllSettings) DeepCopyInto(out *AllSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllSettings.
func (in *AllSettings) DeepCopy() *AllSettings {
	if in == nil {
		return nil
	}
	out := new(AllSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticBeanstalkEnvironment) DeepCopyInto(out *ElasticBeanstalkEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticBeanstalkEnvironment.
func (in *ElasticBeanstalkEnvironment) DeepCopy() *ElasticBeanstalkEnvironment {
	if in == nil {
		return nil
	}
	out := new(ElasticBeanstalkEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticBeanstalkEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticBeanstalkEnvironmentList) DeepCopyInto(out *ElasticBeanstalkEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticBeanstalkEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticBeanstalkEnvironmentList.
func (in *ElasticBeanstalkEnvironmentList) DeepCopy() *ElasticBeanstalkEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(ElasticBeanstalkEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticBeanstalkEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticBeanstalkEnvironmentObservation) DeepCopyInto(out *ElasticBeanstalkEnvironmentObservation) {
	*out = *in
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AutoscalingGroups != nil {
		in, out := &in.AutoscalingGroups, &out.AutoscalingGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Queues != nil {
		in, out := &in.Queues, &out.Queues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllSettings != nil {
		in, out := &in.AllSettings, &out.AllSettings
		*out = make([]AllSettings, len(*in))
		copy(*out, *in)
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LoadBalancers != nil {
		in, out := &in.LoadBalancers, &out.LoadBalancers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LaunchConfigurations != nil {
		in, out := &in.LaunchConfigurations, &out.LaunchConfigurations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticBeanstalkEnvironmentObservation.
func (in *ElasticBeanstalkEnvironmentObservation) DeepCopy() *ElasticBeanstalkEnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(ElasticBeanstalkEnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticBeanstalkEnvironmentParameters) DeepCopyInto(out *ElasticBeanstalkEnvironmentParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Setting = in.Setting
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticBeanstalkEnvironmentParameters.
func (in *ElasticBeanstalkEnvironmentParameters) DeepCopy() *ElasticBeanstalkEnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(ElasticBeanstalkEnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticBeanstalkEnvironmentSpec) DeepCopyInto(out *ElasticBeanstalkEnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticBeanstalkEnvironmentSpec.
func (in *ElasticBeanstalkEnvironmentSpec) DeepCopy() *ElasticBeanstalkEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticBeanstalkEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticBeanstalkEnvironmentStatus) DeepCopyInto(out *ElasticBeanstalkEnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticBeanstalkEnvironmentStatus.
func (in *ElasticBeanstalkEnvironmentStatus) DeepCopy() *ElasticBeanstalkEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticBeanstalkEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Setting) DeepCopyInto(out *Setting) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Setting.
func (in *Setting) DeepCopy() *Setting {
	if in == nil {
		return nil
	}
	out := new(Setting)
	in.DeepCopyInto(out)
	return out
}
