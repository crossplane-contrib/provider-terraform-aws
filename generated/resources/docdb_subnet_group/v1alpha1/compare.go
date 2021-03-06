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
	k := kube.(*DocdbSubnetGroup)
	p := prov.(*DocdbSubnetGroup)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDocdbSubnetGroup_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbSubnetGroup_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbSubnetGroup_SubnetIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbSubnetGroup_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbSubnetGroup_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbSubnetGroup_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDocdbSubnetGroup_Name(k *DocdbSubnetGroupParameters, p *DocdbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbSubnetGroup_NamePrefix(k *DocdbSubnetGroupParameters, p *DocdbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDocdbSubnetGroup_SubnetIds(k *DocdbSubnetGroupParameters, p *DocdbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SubnetIds, p.SubnetIds) {
		p.SubnetIds = k.SubnetIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDocdbSubnetGroup_Tags(k *DocdbSubnetGroupParameters, p *DocdbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbSubnetGroup_Description(k *DocdbSubnetGroupParameters, p *DocdbSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbSubnetGroup_Arn(k *DocdbSubnetGroupObservation, p *DocdbSubnetGroupObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}