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
func (in *CatalogTarget) DeepCopyInto(out *CatalogTarget) {
	*out = *in
	if in.Tables != nil {
		in, out := &in.Tables, &out.Tables
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogTarget.
func (in *CatalogTarget) DeepCopy() *CatalogTarget {
	if in == nil {
		return nil
	}
	out := new(CatalogTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamodbTarget) DeepCopyInto(out *DynamodbTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamodbTarget.
func (in *DynamodbTarget) DeepCopy() *DynamodbTarget {
	if in == nil {
		return nil
	}
	out := new(DynamodbTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueCrawler) DeepCopyInto(out *GlueCrawler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueCrawler.
func (in *GlueCrawler) DeepCopy() *GlueCrawler {
	if in == nil {
		return nil
	}
	out := new(GlueCrawler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueCrawler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueCrawlerList) DeepCopyInto(out *GlueCrawlerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlueCrawler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueCrawlerList.
func (in *GlueCrawlerList) DeepCopy() *GlueCrawlerList {
	if in == nil {
		return nil
	}
	out := new(GlueCrawlerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueCrawlerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueCrawlerObservation) DeepCopyInto(out *GlueCrawlerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueCrawlerObservation.
func (in *GlueCrawlerObservation) DeepCopy() *GlueCrawlerObservation {
	if in == nil {
		return nil
	}
	out := new(GlueCrawlerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueCrawlerParameters) DeepCopyInto(out *GlueCrawlerParameters) {
	*out = *in
	if in.Classifiers != nil {
		in, out := &in.Classifiers, &out.Classifiers
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
	in.CatalogTarget.DeepCopyInto(&out.CatalogTarget)
	out.DynamodbTarget = in.DynamodbTarget
	in.JdbcTarget.DeepCopyInto(&out.JdbcTarget)
	in.S3Target.DeepCopyInto(&out.S3Target)
	out.SchemaChangePolicy = in.SchemaChangePolicy
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueCrawlerParameters.
func (in *GlueCrawlerParameters) DeepCopy() *GlueCrawlerParameters {
	if in == nil {
		return nil
	}
	out := new(GlueCrawlerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueCrawlerSpec) DeepCopyInto(out *GlueCrawlerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueCrawlerSpec.
func (in *GlueCrawlerSpec) DeepCopy() *GlueCrawlerSpec {
	if in == nil {
		return nil
	}
	out := new(GlueCrawlerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueCrawlerStatus) DeepCopyInto(out *GlueCrawlerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueCrawlerStatus.
func (in *GlueCrawlerStatus) DeepCopy() *GlueCrawlerStatus {
	if in == nil {
		return nil
	}
	out := new(GlueCrawlerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JdbcTarget) DeepCopyInto(out *JdbcTarget) {
	*out = *in
	if in.Exclusions != nil {
		in, out := &in.Exclusions, &out.Exclusions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JdbcTarget.
func (in *JdbcTarget) DeepCopy() *JdbcTarget {
	if in == nil {
		return nil
	}
	out := new(JdbcTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Target) DeepCopyInto(out *S3Target) {
	*out = *in
	if in.Exclusions != nil {
		in, out := &in.Exclusions, &out.Exclusions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Target.
func (in *S3Target) DeepCopy() *S3Target {
	if in == nil {
		return nil
	}
	out := new(S3Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaChangePolicy) DeepCopyInto(out *SchemaChangePolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaChangePolicy.
func (in *SchemaChangePolicy) DeepCopy() *SchemaChangePolicy {
	if in == nil {
		return nil
	}
	out := new(SchemaChangePolicy)
	in.DeepCopyInto(out)
	return out
}