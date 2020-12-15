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
	k := kube.(*Ec2TransitGatewayVpcAttachment)
	p := prov.(*Ec2TransitGatewayVpcAttachment)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2TransitGatewayVpcAttachment_TransitGatewayDefaultRouteTableAssociation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachment_TransitGatewayDefaultRouteTablePropagation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachment_TransitGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachment_Ipv6Support(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachment_SubnetIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachment_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachment_VpcId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachment_DnsSupport(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachment_VpcOwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEc2TransitGatewayVpcAttachment_TransitGatewayDefaultRouteTableAssociation(k *Ec2TransitGatewayVpcAttachmentParameters, p *Ec2TransitGatewayVpcAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayDefaultRouteTableAssociation != p.TransitGatewayDefaultRouteTableAssociation {
		p.TransitGatewayDefaultRouteTableAssociation = k.TransitGatewayDefaultRouteTableAssociation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayVpcAttachment_TransitGatewayDefaultRouteTablePropagation(k *Ec2TransitGatewayVpcAttachmentParameters, p *Ec2TransitGatewayVpcAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayDefaultRouteTablePropagation != p.TransitGatewayDefaultRouteTablePropagation {
		p.TransitGatewayDefaultRouteTablePropagation = k.TransitGatewayDefaultRouteTablePropagation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayVpcAttachment_TransitGatewayId(k *Ec2TransitGatewayVpcAttachmentParameters, p *Ec2TransitGatewayVpcAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayId != p.TransitGatewayId {
		p.TransitGatewayId = k.TransitGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayVpcAttachment_Ipv6Support(k *Ec2TransitGatewayVpcAttachmentParameters, p *Ec2TransitGatewayVpcAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.Ipv6Support != p.Ipv6Support {
		p.Ipv6Support = k.Ipv6Support
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEc2TransitGatewayVpcAttachment_SubnetIds(k *Ec2TransitGatewayVpcAttachmentParameters, p *Ec2TransitGatewayVpcAttachmentParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SubnetIds, p.SubnetIds) {
		p.SubnetIds = k.SubnetIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEc2TransitGatewayVpcAttachment_Tags(k *Ec2TransitGatewayVpcAttachmentParameters, p *Ec2TransitGatewayVpcAttachmentParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayVpcAttachment_VpcId(k *Ec2TransitGatewayVpcAttachmentParameters, p *Ec2TransitGatewayVpcAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		p.VpcId = k.VpcId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayVpcAttachment_DnsSupport(k *Ec2TransitGatewayVpcAttachmentParameters, p *Ec2TransitGatewayVpcAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.DnsSupport != p.DnsSupport {
		p.DnsSupport = k.DnsSupport
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2TransitGatewayVpcAttachment_VpcOwnerId(k *Ec2TransitGatewayVpcAttachmentObservation, p *Ec2TransitGatewayVpcAttachmentObservation, md *plugin.MergeDescription) bool {
	if k.VpcOwnerId != p.VpcOwnerId {
		k.VpcOwnerId = p.VpcOwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}