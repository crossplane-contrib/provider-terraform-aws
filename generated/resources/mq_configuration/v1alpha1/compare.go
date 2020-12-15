/*
	Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*MqConfiguration)
	p := prov.(*MqConfiguration)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeMqConfiguration_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMqConfiguration_Data(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMqConfiguration_EngineType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMqConfiguration_EngineVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMqConfiguration_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMqConfiguration_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMqConfiguration_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMqConfiguration_LatestRevision(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveTemplateSpec
func MergeMqConfiguration_Description(k *MqConfigurationParameters, p *MqConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeMqConfiguration_Data(k *MqConfigurationParameters, p *MqConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.Data != p.Data {
		p.Data = k.Data
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeMqConfiguration_EngineType(k *MqConfigurationParameters, p *MqConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.EngineType != p.EngineType {
		p.EngineType = k.EngineType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeMqConfiguration_EngineVersion(k *MqConfigurationParameters, p *MqConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		p.EngineVersion = k.EngineVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeMqConfiguration_Name(k *MqConfigurationParameters, p *MqConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeMqConfiguration_Tags(k *MqConfigurationParameters, p *MqConfigurationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeMqConfiguration_Arn(k *MqConfigurationObservation, p *MqConfigurationObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeMqConfiguration_LatestRevision(k *MqConfigurationObservation, p *MqConfigurationObservation, md *plugin.MergeDescription) bool {
	if k.LatestRevision != p.LatestRevision {
		k.LatestRevision = p.LatestRevision
		md.StatusUpdated = true
		return true
	}
	return false
}