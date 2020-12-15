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
	k := kube.(*DbSubnetGroup)
	p := prov.(*DbSubnetGroup)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDbSubnetGroup_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSubnetGroup_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSubnetGroup_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSubnetGroup_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSubnetGroup_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSubnetGroup_SubnetIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbSubnetGroup_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDbSubnetGroup_Tags(k *DbSubnetGroupParameters, p *DbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDbSubnetGroup_Description(k *DbSubnetGroupParameters, p *DbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDbSubnetGroup_Id(k *DbSubnetGroupParameters, p *DbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDbSubnetGroup_Name(k *DbSubnetGroupParameters, p *DbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDbSubnetGroup_NamePrefix(k *DbSubnetGroupParameters, p *DbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDbSubnetGroup_SubnetIds(k *DbSubnetGroupParameters, p *DbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.SubnetIds, p.SubnetIds) {
		p.SubnetIds = k.SubnetIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbSubnetGroup_Arn(k *DbSubnetGroupObservation, p *DbSubnetGroupObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}