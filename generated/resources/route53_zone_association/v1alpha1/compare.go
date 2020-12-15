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
	k := kube.(*Route53ZoneAssociation)
	p := prov.(*Route53ZoneAssociation)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeRoute53ZoneAssociation_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53ZoneAssociation_VpcId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53ZoneAssociation_VpcRegion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53ZoneAssociation_ZoneId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53ZoneAssociation_OwningAccount(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeRoute53ZoneAssociation_Id(k *Route53ZoneAssociationParameters, p *Route53ZoneAssociationParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53ZoneAssociation_VpcId(k *Route53ZoneAssociationParameters, p *Route53ZoneAssociationParameters, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		p.VpcId = k.VpcId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53ZoneAssociation_VpcRegion(k *Route53ZoneAssociationParameters, p *Route53ZoneAssociationParameters, md *plugin.MergeDescription) bool {
	if k.VpcRegion != p.VpcRegion {
		p.VpcRegion = k.VpcRegion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53ZoneAssociation_ZoneId(k *Route53ZoneAssociationParameters, p *Route53ZoneAssociationParameters, md *plugin.MergeDescription) bool {
	if k.ZoneId != p.ZoneId {
		p.ZoneId = k.ZoneId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRoute53ZoneAssociation_OwningAccount(k *Route53ZoneAssociationObservation, p *Route53ZoneAssociationObservation, md *plugin.MergeDescription) bool {
	if k.OwningAccount != p.OwningAccount {
		k.OwningAccount = p.OwningAccount
		md.StatusUpdated = true
		return true
	}
	return false
}