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
func (in *Ec2TrafficMirrorSession) DeepCopyInto(out *Ec2TrafficMirrorSession) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorSession.
func (in *Ec2TrafficMirrorSession) DeepCopy() *Ec2TrafficMirrorSession {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorSession)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TrafficMirrorSession) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorSessionList) DeepCopyInto(out *Ec2TrafficMirrorSessionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ec2TrafficMirrorSession, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorSessionList.
func (in *Ec2TrafficMirrorSessionList) DeepCopy() *Ec2TrafficMirrorSessionList {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorSessionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TrafficMirrorSessionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorSessionObservation) DeepCopyInto(out *Ec2TrafficMirrorSessionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorSessionObservation.
func (in *Ec2TrafficMirrorSessionObservation) DeepCopy() *Ec2TrafficMirrorSessionObservation {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorSessionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorSessionParameters) DeepCopyInto(out *Ec2TrafficMirrorSessionParameters) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorSessionParameters.
func (in *Ec2TrafficMirrorSessionParameters) DeepCopy() *Ec2TrafficMirrorSessionParameters {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorSessionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorSessionSpec) DeepCopyInto(out *Ec2TrafficMirrorSessionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorSessionSpec.
func (in *Ec2TrafficMirrorSessionSpec) DeepCopy() *Ec2TrafficMirrorSessionSpec {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorSessionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorSessionStatus) DeepCopyInto(out *Ec2TrafficMirrorSessionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorSessionStatus.
func (in *Ec2TrafficMirrorSessionStatus) DeepCopy() *Ec2TrafficMirrorSessionStatus {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorSessionStatus)
	in.DeepCopyInto(out)
	return out
}
