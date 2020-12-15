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
	k := kube.(*DxGatewayAssociation)
	p := prov.(*DxGatewayAssociation)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDxGatewayAssociation_AllowedPrefixes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_AssociatedGatewayOwnerAccountId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_AssociatedGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_DxGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_ProposalId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_AssociatedGatewayType(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_DxGatewayAssociationId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_DxGatewayOwnerAccountId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDxGatewayAssociation_AllowedPrefixes(k *DxGatewayAssociationParameters, p *DxGatewayAssociationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.AllowedPrefixes, p.AllowedPrefixes) {
		p.AllowedPrefixes = k.AllowedPrefixes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociation_AssociatedGatewayOwnerAccountId(k *DxGatewayAssociationParameters, p *DxGatewayAssociationParameters, md *plugin.MergeDescription) bool {
	if k.AssociatedGatewayOwnerAccountId != p.AssociatedGatewayOwnerAccountId {
		p.AssociatedGatewayOwnerAccountId = k.AssociatedGatewayOwnerAccountId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociation_AssociatedGatewayId(k *DxGatewayAssociationParameters, p *DxGatewayAssociationParameters, md *plugin.MergeDescription) bool {
	if k.AssociatedGatewayId != p.AssociatedGatewayId {
		p.AssociatedGatewayId = k.AssociatedGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociation_DxGatewayId(k *DxGatewayAssociationParameters, p *DxGatewayAssociationParameters, md *plugin.MergeDescription) bool {
	if k.DxGatewayId != p.DxGatewayId {
		p.DxGatewayId = k.DxGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociation_ProposalId(k *DxGatewayAssociationParameters, p *DxGatewayAssociationParameters, md *plugin.MergeDescription) bool {
	if k.ProposalId != p.ProposalId {
		p.ProposalId = k.ProposalId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDxGatewayAssociation_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDxGatewayAssociation_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociation_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociation_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociation_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociation_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxGatewayAssociation_AssociatedGatewayType(k *DxGatewayAssociationObservation, p *DxGatewayAssociationObservation, md *plugin.MergeDescription) bool {
	if k.AssociatedGatewayType != p.AssociatedGatewayType {
		k.AssociatedGatewayType = p.AssociatedGatewayType
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxGatewayAssociation_DxGatewayAssociationId(k *DxGatewayAssociationObservation, p *DxGatewayAssociationObservation, md *plugin.MergeDescription) bool {
	if k.DxGatewayAssociationId != p.DxGatewayAssociationId {
		k.DxGatewayAssociationId = p.DxGatewayAssociationId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxGatewayAssociation_DxGatewayOwnerAccountId(k *DxGatewayAssociationObservation, p *DxGatewayAssociationObservation, md *plugin.MergeDescription) bool {
	if k.DxGatewayOwnerAccountId != p.DxGatewayOwnerAccountId {
		k.DxGatewayOwnerAccountId = p.DxGatewayOwnerAccountId
		md.StatusUpdated = true
		return true
	}
	return false
}