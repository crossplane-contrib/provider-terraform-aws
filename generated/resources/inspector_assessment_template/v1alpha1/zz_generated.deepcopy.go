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
func (in *InspectorAssessmentTemplate) DeepCopyInto(out *InspectorAssessmentTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectorAssessmentTemplate.
func (in *InspectorAssessmentTemplate) DeepCopy() *InspectorAssessmentTemplate {
	if in == nil {
		return nil
	}
	out := new(InspectorAssessmentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectorAssessmentTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectorAssessmentTemplateList) DeepCopyInto(out *InspectorAssessmentTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InspectorAssessmentTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectorAssessmentTemplateList.
func (in *InspectorAssessmentTemplateList) DeepCopy() *InspectorAssessmentTemplateList {
	if in == nil {
		return nil
	}
	out := new(InspectorAssessmentTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectorAssessmentTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectorAssessmentTemplateObservation) DeepCopyInto(out *InspectorAssessmentTemplateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectorAssessmentTemplateObservation.
func (in *InspectorAssessmentTemplateObservation) DeepCopy() *InspectorAssessmentTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(InspectorAssessmentTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectorAssessmentTemplateParameters) DeepCopyInto(out *InspectorAssessmentTemplateParameters) {
	*out = *in
	if in.RulesPackageArns != nil {
		in, out := &in.RulesPackageArns, &out.RulesPackageArns
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectorAssessmentTemplateParameters.
func (in *InspectorAssessmentTemplateParameters) DeepCopy() *InspectorAssessmentTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(InspectorAssessmentTemplateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectorAssessmentTemplateSpec) DeepCopyInto(out *InspectorAssessmentTemplateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectorAssessmentTemplateSpec.
func (in *InspectorAssessmentTemplateSpec) DeepCopy() *InspectorAssessmentTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(InspectorAssessmentTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectorAssessmentTemplateStatus) DeepCopyInto(out *InspectorAssessmentTemplateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectorAssessmentTemplateStatus.
func (in *InspectorAssessmentTemplateStatus) DeepCopy() *InspectorAssessmentTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(InspectorAssessmentTemplateStatus)
	in.DeepCopyInto(out)
	return out
}
