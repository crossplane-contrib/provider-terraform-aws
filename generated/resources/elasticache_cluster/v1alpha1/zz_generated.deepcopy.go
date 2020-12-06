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
func (in *CacheNodes) DeepCopyInto(out *CacheNodes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheNodes.
func (in *CacheNodes) DeepCopy() *CacheNodes {
	if in == nil {
		return nil
	}
	out := new(CacheNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheCluster) DeepCopyInto(out *ElasticacheCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheCluster.
func (in *ElasticacheCluster) DeepCopy() *ElasticacheCluster {
	if in == nil {
		return nil
	}
	out := new(ElasticacheCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticacheCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterList) DeepCopyInto(out *ElasticacheClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticacheCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterList.
func (in *ElasticacheClusterList) DeepCopy() *ElasticacheClusterList {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticacheClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterObservation) DeepCopyInto(out *ElasticacheClusterObservation) {
	*out = *in
	if in.CacheNodes != nil {
		in, out := &in.CacheNodes, &out.CacheNodes
		*out = make([]CacheNodes, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterObservation.
func (in *ElasticacheClusterObservation) DeepCopy() *ElasticacheClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterParameters) DeepCopyInto(out *ElasticacheClusterParameters) {
	*out = *in
	if in.PreferredAvailabilityZones != nil {
		in, out := &in.PreferredAvailabilityZones, &out.PreferredAvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupNames != nil {
		in, out := &in.SecurityGroupNames, &out.SecurityGroupNames
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
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotArns != nil {
		in, out := &in.SnapshotArns, &out.SnapshotArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterParameters.
func (in *ElasticacheClusterParameters) DeepCopy() *ElasticacheClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterSpec) DeepCopyInto(out *ElasticacheClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterSpec.
func (in *ElasticacheClusterSpec) DeepCopy() *ElasticacheClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticacheClusterStatus) DeepCopyInto(out *ElasticacheClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticacheClusterStatus.
func (in *ElasticacheClusterStatus) DeepCopy() *ElasticacheClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticacheClusterStatus)
	in.DeepCopyInto(out)
	return out
}
