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
func (in *Columns) DeepCopyInto(out *Columns) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Columns.
func (in *Columns) DeepCopy() *Columns {
	if in == nil {
		return nil
	}
	out := new(Columns)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GluePartition) DeepCopyInto(out *GluePartition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GluePartition.
func (in *GluePartition) DeepCopy() *GluePartition {
	if in == nil {
		return nil
	}
	out := new(GluePartition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GluePartition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GluePartitionList) DeepCopyInto(out *GluePartitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GluePartition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GluePartitionList.
func (in *GluePartitionList) DeepCopy() *GluePartitionList {
	if in == nil {
		return nil
	}
	out := new(GluePartitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GluePartitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GluePartitionObservation) DeepCopyInto(out *GluePartitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GluePartitionObservation.
func (in *GluePartitionObservation) DeepCopy() *GluePartitionObservation {
	if in == nil {
		return nil
	}
	out := new(GluePartitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GluePartitionParameters) DeepCopyInto(out *GluePartitionParameters) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PartitionValues != nil {
		in, out := &in.PartitionValues, &out.PartitionValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.StorageDescriptor.DeepCopyInto(&out.StorageDescriptor)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GluePartitionParameters.
func (in *GluePartitionParameters) DeepCopy() *GluePartitionParameters {
	if in == nil {
		return nil
	}
	out := new(GluePartitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GluePartitionSpec) DeepCopyInto(out *GluePartitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GluePartitionSpec.
func (in *GluePartitionSpec) DeepCopy() *GluePartitionSpec {
	if in == nil {
		return nil
	}
	out := new(GluePartitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GluePartitionStatus) DeepCopyInto(out *GluePartitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GluePartitionStatus.
func (in *GluePartitionStatus) DeepCopy() *GluePartitionStatus {
	if in == nil {
		return nil
	}
	out := new(GluePartitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SerDeInfo) DeepCopyInto(out *SerDeInfo) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SerDeInfo.
func (in *SerDeInfo) DeepCopy() *SerDeInfo {
	if in == nil {
		return nil
	}
	out := new(SerDeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkewedInfo) DeepCopyInto(out *SkewedInfo) {
	*out = *in
	if in.SkewedColumnNames != nil {
		in, out := &in.SkewedColumnNames, &out.SkewedColumnNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SkewedColumnValueLocationMaps != nil {
		in, out := &in.SkewedColumnValueLocationMaps, &out.SkewedColumnValueLocationMaps
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SkewedColumnValues != nil {
		in, out := &in.SkewedColumnValues, &out.SkewedColumnValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkewedInfo.
func (in *SkewedInfo) DeepCopy() *SkewedInfo {
	if in == nil {
		return nil
	}
	out := new(SkewedInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SortColumns) DeepCopyInto(out *SortColumns) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SortColumns.
func (in *SortColumns) DeepCopy() *SortColumns {
	if in == nil {
		return nil
	}
	out := new(SortColumns)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDescriptor) DeepCopyInto(out *StorageDescriptor) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.BucketColumns != nil {
		in, out := &in.BucketColumns, &out.BucketColumns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Columns = in.Columns
	in.SerDeInfo.DeepCopyInto(&out.SerDeInfo)
	in.SkewedInfo.DeepCopyInto(&out.SkewedInfo)
	out.SortColumns = in.SortColumns
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDescriptor.
func (in *StorageDescriptor) DeepCopy() *StorageDescriptor {
	if in == nil {
		return nil
	}
	out := new(StorageDescriptor)
	in.DeepCopyInto(out)
	return out
}
