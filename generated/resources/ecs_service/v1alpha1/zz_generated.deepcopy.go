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
func (in *CapacityProviderStrategy) DeepCopyInto(out *CapacityProviderStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityProviderStrategy.
func (in *CapacityProviderStrategy) DeepCopy() *CapacityProviderStrategy {
	if in == nil {
		return nil
	}
	out := new(CapacityProviderStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentController) DeepCopyInto(out *DeploymentController) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentController.
func (in *DeploymentController) DeepCopy() *DeploymentController {
	if in == nil {
		return nil
	}
	out := new(DeploymentController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcsService) DeepCopyInto(out *EcsService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcsService.
func (in *EcsService) DeepCopy() *EcsService {
	if in == nil {
		return nil
	}
	out := new(EcsService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EcsService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcsServiceList) DeepCopyInto(out *EcsServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EcsService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcsServiceList.
func (in *EcsServiceList) DeepCopy() *EcsServiceList {
	if in == nil {
		return nil
	}
	out := new(EcsServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EcsServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcsServiceObservation) DeepCopyInto(out *EcsServiceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcsServiceObservation.
func (in *EcsServiceObservation) DeepCopy() *EcsServiceObservation {
	if in == nil {
		return nil
	}
	out := new(EcsServiceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcsServiceParameters) DeepCopyInto(out *EcsServiceParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.LoadBalancer = in.LoadBalancer
	in.NetworkConfiguration.DeepCopyInto(&out.NetworkConfiguration)
	if in.OrderedPlacementStrategy != nil {
		in, out := &in.OrderedPlacementStrategy, &out.OrderedPlacementStrategy
		*out = make([]OrderedPlacementStrategy, len(*in))
		copy(*out, *in)
	}
	if in.PlacementConstraints != nil {
		in, out := &in.PlacementConstraints, &out.PlacementConstraints
		*out = make([]PlacementConstraints, len(*in))
		copy(*out, *in)
	}
	out.ServiceRegistries = in.ServiceRegistries
	out.Timeouts = in.Timeouts
	out.CapacityProviderStrategy = in.CapacityProviderStrategy
	out.DeploymentController = in.DeploymentController
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcsServiceParameters.
func (in *EcsServiceParameters) DeepCopy() *EcsServiceParameters {
	if in == nil {
		return nil
	}
	out := new(EcsServiceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcsServiceSpec) DeepCopyInto(out *EcsServiceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcsServiceSpec.
func (in *EcsServiceSpec) DeepCopy() *EcsServiceSpec {
	if in == nil {
		return nil
	}
	out := new(EcsServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcsServiceStatus) DeepCopyInto(out *EcsServiceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcsServiceStatus.
func (in *EcsServiceStatus) DeepCopy() *EcsServiceStatus {
	if in == nil {
		return nil
	}
	out := new(EcsServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfiguration) DeepCopyInto(out *NetworkConfiguration) {
	*out = *in
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfiguration.
func (in *NetworkConfiguration) DeepCopy() *NetworkConfiguration {
	if in == nil {
		return nil
	}
	out := new(NetworkConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderedPlacementStrategy) DeepCopyInto(out *OrderedPlacementStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderedPlacementStrategy.
func (in *OrderedPlacementStrategy) DeepCopy() *OrderedPlacementStrategy {
	if in == nil {
		return nil
	}
	out := new(OrderedPlacementStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementConstraints) DeepCopyInto(out *PlacementConstraints) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementConstraints.
func (in *PlacementConstraints) DeepCopy() *PlacementConstraints {
	if in == nil {
		return nil
	}
	out := new(PlacementConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRegistries) DeepCopyInto(out *ServiceRegistries) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRegistries.
func (in *ServiceRegistries) DeepCopy() *ServiceRegistries {
	if in == nil {
		return nil
	}
	out := new(ServiceRegistries)
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
