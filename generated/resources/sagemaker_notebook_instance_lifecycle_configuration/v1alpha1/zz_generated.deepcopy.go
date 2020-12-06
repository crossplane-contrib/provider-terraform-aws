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
func (in *SagemakerNotebookInstanceLifecycleConfiguration) DeepCopyInto(out *SagemakerNotebookInstanceLifecycleConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceLifecycleConfiguration.
func (in *SagemakerNotebookInstanceLifecycleConfiguration) DeepCopy() *SagemakerNotebookInstanceLifecycleConfiguration {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceLifecycleConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SagemakerNotebookInstanceLifecycleConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceLifecycleConfigurationList) DeepCopyInto(out *SagemakerNotebookInstanceLifecycleConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SagemakerNotebookInstanceLifecycleConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceLifecycleConfigurationList.
func (in *SagemakerNotebookInstanceLifecycleConfigurationList) DeepCopy() *SagemakerNotebookInstanceLifecycleConfigurationList {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceLifecycleConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SagemakerNotebookInstanceLifecycleConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceLifecycleConfigurationObservation) DeepCopyInto(out *SagemakerNotebookInstanceLifecycleConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceLifecycleConfigurationObservation.
func (in *SagemakerNotebookInstanceLifecycleConfigurationObservation) DeepCopy() *SagemakerNotebookInstanceLifecycleConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceLifecycleConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceLifecycleConfigurationParameters) DeepCopyInto(out *SagemakerNotebookInstanceLifecycleConfigurationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceLifecycleConfigurationParameters.
func (in *SagemakerNotebookInstanceLifecycleConfigurationParameters) DeepCopy() *SagemakerNotebookInstanceLifecycleConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceLifecycleConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceLifecycleConfigurationSpec) DeepCopyInto(out *SagemakerNotebookInstanceLifecycleConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceLifecycleConfigurationSpec.
func (in *SagemakerNotebookInstanceLifecycleConfigurationSpec) DeepCopy() *SagemakerNotebookInstanceLifecycleConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceLifecycleConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerNotebookInstanceLifecycleConfigurationStatus) DeepCopyInto(out *SagemakerNotebookInstanceLifecycleConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerNotebookInstanceLifecycleConfigurationStatus.
func (in *SagemakerNotebookInstanceLifecycleConfigurationStatus) DeepCopy() *SagemakerNotebookInstanceLifecycleConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(SagemakerNotebookInstanceLifecycleConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}
