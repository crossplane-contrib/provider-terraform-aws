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
func (in *ConnectionPasswordEncryption) DeepCopyInto(out *ConnectionPasswordEncryption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPasswordEncryption.
func (in *ConnectionPasswordEncryption) DeepCopy() *ConnectionPasswordEncryption {
	if in == nil {
		return nil
	}
	out := new(ConnectionPasswordEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogEncryptionSettings) DeepCopyInto(out *DataCatalogEncryptionSettings) {
	*out = *in
	out.ConnectionPasswordEncryption = in.ConnectionPasswordEncryption
	out.EncryptionAtRest = in.EncryptionAtRest
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogEncryptionSettings.
func (in *DataCatalogEncryptionSettings) DeepCopy() *DataCatalogEncryptionSettings {
	if in == nil {
		return nil
	}
	out := new(DataCatalogEncryptionSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionAtRest) DeepCopyInto(out *EncryptionAtRest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionAtRest.
func (in *EncryptionAtRest) DeepCopy() *EncryptionAtRest {
	if in == nil {
		return nil
	}
	out := new(EncryptionAtRest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettings) DeepCopyInto(out *GlueDataCatalogEncryptionSettings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettings.
func (in *GlueDataCatalogEncryptionSettings) DeepCopy() *GlueDataCatalogEncryptionSettings {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueDataCatalogEncryptionSettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsList) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlueDataCatalogEncryptionSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsList.
func (in *GlueDataCatalogEncryptionSettingsList) DeepCopy() *GlueDataCatalogEncryptionSettingsList {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueDataCatalogEncryptionSettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsObservation) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsObservation.
func (in *GlueDataCatalogEncryptionSettingsObservation) DeepCopy() *GlueDataCatalogEncryptionSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsParameters) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsParameters) {
	*out = *in
	out.DataCatalogEncryptionSettings = in.DataCatalogEncryptionSettings
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsParameters.
func (in *GlueDataCatalogEncryptionSettingsParameters) DeepCopy() *GlueDataCatalogEncryptionSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsSpec) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsSpec.
func (in *GlueDataCatalogEncryptionSettingsSpec) DeepCopy() *GlueDataCatalogEncryptionSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsStatus) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsStatus.
func (in *GlueDataCatalogEncryptionSettingsStatus) DeepCopy() *GlueDataCatalogEncryptionSettingsStatus {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsStatus)
	in.DeepCopyInto(out)
	return out
}
