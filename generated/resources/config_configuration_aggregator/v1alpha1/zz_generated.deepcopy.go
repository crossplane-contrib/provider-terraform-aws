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
func (in *AccountAggregationSource) DeepCopyInto(out *AccountAggregationSource) {
	*out = *in
	if in.AccountIds != nil {
		in, out := &in.AccountIds, &out.AccountIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountAggregationSource.
func (in *AccountAggregationSource) DeepCopy() *AccountAggregationSource {
	if in == nil {
		return nil
	}
	out := new(AccountAggregationSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigurationAggregator) DeepCopyInto(out *ConfigConfigurationAggregator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigurationAggregator.
func (in *ConfigConfigurationAggregator) DeepCopy() *ConfigConfigurationAggregator {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigurationAggregator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigConfigurationAggregator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigurationAggregatorList) DeepCopyInto(out *ConfigConfigurationAggregatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigConfigurationAggregator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigurationAggregatorList.
func (in *ConfigConfigurationAggregatorList) DeepCopy() *ConfigConfigurationAggregatorList {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigurationAggregatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigConfigurationAggregatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigurationAggregatorObservation) DeepCopyInto(out *ConfigConfigurationAggregatorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigurationAggregatorObservation.
func (in *ConfigConfigurationAggregatorObservation) DeepCopy() *ConfigConfigurationAggregatorObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigurationAggregatorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigurationAggregatorParameters) DeepCopyInto(out *ConfigConfigurationAggregatorParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AccountAggregationSource.DeepCopyInto(&out.AccountAggregationSource)
	in.OrganizationAggregationSource.DeepCopyInto(&out.OrganizationAggregationSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigurationAggregatorParameters.
func (in *ConfigConfigurationAggregatorParameters) DeepCopy() *ConfigConfigurationAggregatorParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigurationAggregatorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigurationAggregatorSpec) DeepCopyInto(out *ConfigConfigurationAggregatorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigurationAggregatorSpec.
func (in *ConfigConfigurationAggregatorSpec) DeepCopy() *ConfigConfigurationAggregatorSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigurationAggregatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConfigurationAggregatorStatus) DeepCopyInto(out *ConfigConfigurationAggregatorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConfigurationAggregatorStatus.
func (in *ConfigConfigurationAggregatorStatus) DeepCopy() *ConfigConfigurationAggregatorStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigConfigurationAggregatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationAggregationSource) DeepCopyInto(out *OrganizationAggregationSource) {
	*out = *in
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationAggregationSource.
func (in *OrganizationAggregationSource) DeepCopy() *OrganizationAggregationSource {
	if in == nil {
		return nil
	}
	out := new(OrganizationAggregationSource)
	in.DeepCopyInto(out)
	return out
}
