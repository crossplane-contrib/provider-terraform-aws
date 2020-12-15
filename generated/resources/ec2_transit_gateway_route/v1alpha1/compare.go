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
	k := kube.(*Ec2TransitGatewayRoute)
	p := prov.(*Ec2TransitGatewayRoute)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2TransitGatewayRoute_Blackhole(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayRoute_DestinationCidrBlock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayRoute_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayRoute_TransitGatewayAttachmentId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayRoute_TransitGatewayRouteTableId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeEc2TransitGatewayRoute_Blackhole(k *Ec2TransitGatewayRouteParameters, p *Ec2TransitGatewayRouteParameters, md *plugin.MergeDescription) bool {
	if k.Blackhole != p.Blackhole {
		p.Blackhole = k.Blackhole
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayRoute_DestinationCidrBlock(k *Ec2TransitGatewayRouteParameters, p *Ec2TransitGatewayRouteParameters, md *plugin.MergeDescription) bool {
	if k.DestinationCidrBlock != p.DestinationCidrBlock {
		p.DestinationCidrBlock = k.DestinationCidrBlock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayRoute_Id(k *Ec2TransitGatewayRouteParameters, p *Ec2TransitGatewayRouteParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayRoute_TransitGatewayAttachmentId(k *Ec2TransitGatewayRouteParameters, p *Ec2TransitGatewayRouteParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayAttachmentId != p.TransitGatewayAttachmentId {
		p.TransitGatewayAttachmentId = k.TransitGatewayAttachmentId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayRoute_TransitGatewayRouteTableId(k *Ec2TransitGatewayRouteParameters, p *Ec2TransitGatewayRouteParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayRouteTableId != p.TransitGatewayRouteTableId {
		p.TransitGatewayRouteTableId = k.TransitGatewayRouteTableId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}