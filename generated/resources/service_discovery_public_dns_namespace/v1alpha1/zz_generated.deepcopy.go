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
func (in *ServiceDiscoveryPublicDnsNamespace) DeepCopyInto(out *ServiceDiscoveryPublicDnsNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryPublicDnsNamespace.
func (in *ServiceDiscoveryPublicDnsNamespace) DeepCopy() *ServiceDiscoveryPublicDnsNamespace {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryPublicDnsNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceDiscoveryPublicDnsNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryPublicDnsNamespaceList) DeepCopyInto(out *ServiceDiscoveryPublicDnsNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceDiscoveryPublicDnsNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryPublicDnsNamespaceList.
func (in *ServiceDiscoveryPublicDnsNamespaceList) DeepCopy() *ServiceDiscoveryPublicDnsNamespaceList {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryPublicDnsNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceDiscoveryPublicDnsNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryPublicDnsNamespaceObservation) DeepCopyInto(out *ServiceDiscoveryPublicDnsNamespaceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryPublicDnsNamespaceObservation.
func (in *ServiceDiscoveryPublicDnsNamespaceObservation) DeepCopy() *ServiceDiscoveryPublicDnsNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryPublicDnsNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryPublicDnsNamespaceParameters) DeepCopyInto(out *ServiceDiscoveryPublicDnsNamespaceParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryPublicDnsNamespaceParameters.
func (in *ServiceDiscoveryPublicDnsNamespaceParameters) DeepCopy() *ServiceDiscoveryPublicDnsNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryPublicDnsNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryPublicDnsNamespaceSpec) DeepCopyInto(out *ServiceDiscoveryPublicDnsNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryPublicDnsNamespaceSpec.
func (in *ServiceDiscoveryPublicDnsNamespaceSpec) DeepCopy() *ServiceDiscoveryPublicDnsNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryPublicDnsNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryPublicDnsNamespaceStatus) DeepCopyInto(out *ServiceDiscoveryPublicDnsNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryPublicDnsNamespaceStatus.
func (in *ServiceDiscoveryPublicDnsNamespaceStatus) DeepCopy() *ServiceDiscoveryPublicDnsNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryPublicDnsNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}
