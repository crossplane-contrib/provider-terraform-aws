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
	k := kube.(*LicensemanagerLicenseConfiguration)
	p := prov.(*LicensemanagerLicenseConfiguration)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeLicensemanagerLicenseConfiguration_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicensemanagerLicenseConfiguration_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicensemanagerLicenseConfiguration_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicensemanagerLicenseConfiguration_LicenseCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicensemanagerLicenseConfiguration_LicenseCountHardLimit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicensemanagerLicenseConfiguration_LicenseCountingType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicensemanagerLicenseConfiguration_LicenseRules(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicensemanagerLicenseConfiguration_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeLicensemanagerLicenseConfiguration_Tags(k *LicensemanagerLicenseConfigurationParameters, p *LicensemanagerLicenseConfigurationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLicensemanagerLicenseConfiguration_Description(k *LicensemanagerLicenseConfigurationParameters, p *LicensemanagerLicenseConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLicensemanagerLicenseConfiguration_Id(k *LicensemanagerLicenseConfigurationParameters, p *LicensemanagerLicenseConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLicensemanagerLicenseConfiguration_LicenseCount(k *LicensemanagerLicenseConfigurationParameters, p *LicensemanagerLicenseConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.LicenseCount != p.LicenseCount {
		p.LicenseCount = k.LicenseCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLicensemanagerLicenseConfiguration_LicenseCountHardLimit(k *LicensemanagerLicenseConfigurationParameters, p *LicensemanagerLicenseConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.LicenseCountHardLimit != p.LicenseCountHardLimit {
		p.LicenseCountHardLimit = k.LicenseCountHardLimit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLicensemanagerLicenseConfiguration_LicenseCountingType(k *LicensemanagerLicenseConfigurationParameters, p *LicensemanagerLicenseConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.LicenseCountingType != p.LicenseCountingType {
		p.LicenseCountingType = k.LicenseCountingType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeLicensemanagerLicenseConfiguration_LicenseRules(k *LicensemanagerLicenseConfigurationParameters, p *LicensemanagerLicenseConfigurationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.LicenseRules, p.LicenseRules) {
		p.LicenseRules = k.LicenseRules
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLicensemanagerLicenseConfiguration_Name(k *LicensemanagerLicenseConfigurationParameters, p *LicensemanagerLicenseConfigurationParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}