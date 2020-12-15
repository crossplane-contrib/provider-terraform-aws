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
	k := kube.(*Ec2ClientVpnNetworkAssociation)
	p := prov.(*Ec2ClientVpnNetworkAssociation)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2ClientVpnNetworkAssociation_SecurityGroups(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnNetworkAssociation_SubnetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnNetworkAssociation_ClientVpnEndpointId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnNetworkAssociation_Status(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnNetworkAssociation_VpcId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnNetworkAssociation_AssociationId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEc2ClientVpnNetworkAssociation_SecurityGroups(k *Ec2ClientVpnNetworkAssociationParameters, p *Ec2ClientVpnNetworkAssociationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SecurityGroups, p.SecurityGroups) {
		p.SecurityGroups = k.SecurityGroups
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2ClientVpnNetworkAssociation_SubnetId(k *Ec2ClientVpnNetworkAssociationParameters, p *Ec2ClientVpnNetworkAssociationParameters, md *plugin.MergeDescription) bool {
	if k.SubnetId != p.SubnetId {
		p.SubnetId = k.SubnetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2ClientVpnNetworkAssociation_ClientVpnEndpointId(k *Ec2ClientVpnNetworkAssociationParameters, p *Ec2ClientVpnNetworkAssociationParameters, md *plugin.MergeDescription) bool {
	if k.ClientVpnEndpointId != p.ClientVpnEndpointId {
		p.ClientVpnEndpointId = k.ClientVpnEndpointId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2ClientVpnNetworkAssociation_Status(k *Ec2ClientVpnNetworkAssociationObservation, p *Ec2ClientVpnNetworkAssociationObservation, md *plugin.MergeDescription) bool {
	if k.Status != p.Status {
		k.Status = p.Status
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2ClientVpnNetworkAssociation_VpcId(k *Ec2ClientVpnNetworkAssociationObservation, p *Ec2ClientVpnNetworkAssociationObservation, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		k.VpcId = p.VpcId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2ClientVpnNetworkAssociation_AssociationId(k *Ec2ClientVpnNetworkAssociationObservation, p *Ec2ClientVpnNetworkAssociationObservation, md *plugin.MergeDescription) bool {
	if k.AssociationId != p.AssociationId {
		k.AssociationId = p.AssociationId
		md.StatusUpdated = true
		return true
	}
	return false
}