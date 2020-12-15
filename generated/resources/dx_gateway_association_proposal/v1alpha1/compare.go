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
	k := kube.(*DxGatewayAssociationProposal)
	p := prov.(*DxGatewayAssociationProposal)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDxGatewayAssociationProposal_AllowedPrefixes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociationProposal_AssociatedGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociationProposal_DxGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociationProposal_DxGatewayOwnerAccountId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociationProposal_AssociatedGatewayOwnerAccountId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxGatewayAssociationProposal_AssociatedGatewayType(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDxGatewayAssociationProposal_AllowedPrefixes(k *DxGatewayAssociationProposalParameters, p *DxGatewayAssociationProposalParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.AllowedPrefixes, p.AllowedPrefixes) {
		p.AllowedPrefixes = k.AllowedPrefixes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociationProposal_AssociatedGatewayId(k *DxGatewayAssociationProposalParameters, p *DxGatewayAssociationProposalParameters, md *plugin.MergeDescription) bool {
	if k.AssociatedGatewayId != p.AssociatedGatewayId {
		p.AssociatedGatewayId = k.AssociatedGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociationProposal_DxGatewayId(k *DxGatewayAssociationProposalParameters, p *DxGatewayAssociationProposalParameters, md *plugin.MergeDescription) bool {
	if k.DxGatewayId != p.DxGatewayId {
		p.DxGatewayId = k.DxGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxGatewayAssociationProposal_DxGatewayOwnerAccountId(k *DxGatewayAssociationProposalParameters, p *DxGatewayAssociationProposalParameters, md *plugin.MergeDescription) bool {
	if k.DxGatewayOwnerAccountId != p.DxGatewayOwnerAccountId {
		p.DxGatewayOwnerAccountId = k.DxGatewayOwnerAccountId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxGatewayAssociationProposal_AssociatedGatewayOwnerAccountId(k *DxGatewayAssociationProposalObservation, p *DxGatewayAssociationProposalObservation, md *plugin.MergeDescription) bool {
	if k.AssociatedGatewayOwnerAccountId != p.AssociatedGatewayOwnerAccountId {
		k.AssociatedGatewayOwnerAccountId = p.AssociatedGatewayOwnerAccountId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxGatewayAssociationProposal_AssociatedGatewayType(k *DxGatewayAssociationProposalObservation, p *DxGatewayAssociationProposalObservation, md *plugin.MergeDescription) bool {
	if k.AssociatedGatewayType != p.AssociatedGatewayType {
		k.AssociatedGatewayType = p.AssociatedGatewayType
		md.StatusUpdated = true
		return true
	}
	return false
}