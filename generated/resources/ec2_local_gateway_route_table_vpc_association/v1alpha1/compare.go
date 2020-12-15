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
	k := kube.(*Ec2LocalGatewayRouteTableVpcAssociation)
	p := prov.(*Ec2LocalGatewayRouteTableVpcAssociation)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2LocalGatewayRouteTableVpcAssociation_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayRouteTableId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2LocalGatewayRouteTableVpcAssociation_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2LocalGatewayRouteTableVpcAssociation_VpcId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEc2LocalGatewayRouteTableVpcAssociation_Id(k *Ec2LocalGatewayRouteTableVpcAssociationParameters, p *Ec2LocalGatewayRouteTableVpcAssociationParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayRouteTableId(k *Ec2LocalGatewayRouteTableVpcAssociationParameters, p *Ec2LocalGatewayRouteTableVpcAssociationParameters, md *plugin.MergeDescription) bool {
	if k.LocalGatewayRouteTableId != p.LocalGatewayRouteTableId {
		p.LocalGatewayRouteTableId = k.LocalGatewayRouteTableId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEc2LocalGatewayRouteTableVpcAssociation_Tags(k *Ec2LocalGatewayRouteTableVpcAssociationParameters, p *Ec2LocalGatewayRouteTableVpcAssociationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2LocalGatewayRouteTableVpcAssociation_VpcId(k *Ec2LocalGatewayRouteTableVpcAssociationParameters, p *Ec2LocalGatewayRouteTableVpcAssociationParameters, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		p.VpcId = k.VpcId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2LocalGatewayRouteTableVpcAssociation_LocalGatewayId(k *Ec2LocalGatewayRouteTableVpcAssociationObservation, p *Ec2LocalGatewayRouteTableVpcAssociationObservation, md *plugin.MergeDescription) bool {
	if k.LocalGatewayId != p.LocalGatewayId {
		k.LocalGatewayId = p.LocalGatewayId
		md.StatusUpdated = true
		return true
	}
	return false
}