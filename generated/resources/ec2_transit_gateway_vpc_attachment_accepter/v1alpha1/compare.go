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
	k := kube.(*Ec2TransitGatewayVpcAttachmentAccepter)
	p := prov.(*Ec2TransitGatewayVpcAttachmentAccepter)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTablePropagation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTableAssociation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayAttachmentId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_VpcId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_VpcOwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_Ipv6Support(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_DnsSupport(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_SubnetIds(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTablePropagation(k *Ec2TransitGatewayVpcAttachmentAccepterParameters, p *Ec2TransitGatewayVpcAttachmentAccepterParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayDefaultRouteTablePropagation != p.TransitGatewayDefaultRouteTablePropagation {
		p.TransitGatewayDefaultRouteTablePropagation = k.TransitGatewayDefaultRouteTablePropagation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEc2TransitGatewayVpcAttachmentAccepter_Tags(k *Ec2TransitGatewayVpcAttachmentAccepterParameters, p *Ec2TransitGatewayVpcAttachmentAccepterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTableAssociation(k *Ec2TransitGatewayVpcAttachmentAccepterParameters, p *Ec2TransitGatewayVpcAttachmentAccepterParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayDefaultRouteTableAssociation != p.TransitGatewayDefaultRouteTableAssociation {
		p.TransitGatewayDefaultRouteTableAssociation = k.TransitGatewayDefaultRouteTableAssociation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayAttachmentId(k *Ec2TransitGatewayVpcAttachmentAccepterParameters, p *Ec2TransitGatewayVpcAttachmentAccepterParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayAttachmentId != p.TransitGatewayAttachmentId {
		p.TransitGatewayAttachmentId = k.TransitGatewayAttachmentId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2TransitGatewayVpcAttachmentAccepter_VpcId(k *Ec2TransitGatewayVpcAttachmentAccepterObservation, p *Ec2TransitGatewayVpcAttachmentAccepterObservation, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		k.VpcId = p.VpcId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2TransitGatewayVpcAttachmentAccepter_VpcOwnerId(k *Ec2TransitGatewayVpcAttachmentAccepterObservation, p *Ec2TransitGatewayVpcAttachmentAccepterObservation, md *plugin.MergeDescription) bool {
	if k.VpcOwnerId != p.VpcOwnerId {
		k.VpcOwnerId = p.VpcOwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2TransitGatewayVpcAttachmentAccepter_Ipv6Support(k *Ec2TransitGatewayVpcAttachmentAccepterObservation, p *Ec2TransitGatewayVpcAttachmentAccepterObservation, md *plugin.MergeDescription) bool {
	if k.Ipv6Support != p.Ipv6Support {
		k.Ipv6Support = p.Ipv6Support
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2TransitGatewayVpcAttachmentAccepter_DnsSupport(k *Ec2TransitGatewayVpcAttachmentAccepterObservation, p *Ec2TransitGatewayVpcAttachmentAccepterObservation, md *plugin.MergeDescription) bool {
	if k.DnsSupport != p.DnsSupport {
		k.DnsSupport = p.DnsSupport
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeEc2TransitGatewayVpcAttachmentAccepter_SubnetIds(k *Ec2TransitGatewayVpcAttachmentAccepterObservation, p *Ec2TransitGatewayVpcAttachmentAccepterObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SubnetIds, p.SubnetIds) {
		k.SubnetIds = p.SubnetIds
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayId(k *Ec2TransitGatewayVpcAttachmentAccepterObservation, p *Ec2TransitGatewayVpcAttachmentAccepterObservation, md *plugin.MergeDescription) bool {
	if k.TransitGatewayId != p.TransitGatewayId {
		k.TransitGatewayId = p.TransitGatewayId
		md.StatusUpdated = true
		return true
	}
	return false
}