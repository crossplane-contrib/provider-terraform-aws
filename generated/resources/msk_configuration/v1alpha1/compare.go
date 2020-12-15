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
	k := kube.(*MskConfiguration)
	p := prov.(*MskConfiguration)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeMskConfiguration_ServerProperties(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMskConfiguration_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMskConfiguration_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMskConfiguration_KafkaVersions(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMskConfiguration_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMskConfiguration_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeMskConfiguration_LatestRevision(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeMskConfiguration_ServerProperties(k *MskConfigurationParameters, p *MskConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.ServerProperties != p.ServerProperties {
		p.ServerProperties = k.ServerProperties
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeMskConfiguration_Description(k *MskConfigurationParameters, p *MskConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeMskConfiguration_Id(k *MskConfigurationParameters, p *MskConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeMskConfiguration_KafkaVersions(k *MskConfigurationParameters, p *MskConfigurationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.KafkaVersions, p.KafkaVersions) {
		p.KafkaVersions = k.KafkaVersions
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeMskConfiguration_Name(k *MskConfigurationParameters, p *MskConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeMskConfiguration_Arn(k *MskConfigurationObservation, p *MskConfigurationObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeMskConfiguration_LatestRevision(k *MskConfigurationObservation, p *MskConfigurationObservation, md *plugin.MergeDescription) bool {
	if k.LatestRevision != p.LatestRevision {
		k.LatestRevision = p.LatestRevision
		md.StatusUpdated = true
		return true
	}
	return false
}