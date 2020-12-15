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
	k := kube.(*Route)
	p := prov.(*Route)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeRoute_NetworkInterfaceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_VpcPeeringConnectionId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_InstanceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_NatGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_DestinationCidrBlock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_EgressOnlyGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_GatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_RouteTableId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_TransitGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_DestinationIpv6CidrBlock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_LocalGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_State(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_Origin(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_DestinationPrefixListId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_InstanceOwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeRoute_NetworkInterfaceId(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.NetworkInterfaceId != p.NetworkInterfaceId {
		p.NetworkInterfaceId = k.NetworkInterfaceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_VpcPeeringConnectionId(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.VpcPeeringConnectionId != p.VpcPeeringConnectionId {
		p.VpcPeeringConnectionId = k.VpcPeeringConnectionId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_InstanceId(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.InstanceId != p.InstanceId {
		p.InstanceId = k.InstanceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_NatGatewayId(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.NatGatewayId != p.NatGatewayId {
		p.NatGatewayId = k.NatGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_DestinationCidrBlock(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.DestinationCidrBlock != p.DestinationCidrBlock {
		p.DestinationCidrBlock = k.DestinationCidrBlock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_EgressOnlyGatewayId(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.EgressOnlyGatewayId != p.EgressOnlyGatewayId {
		p.EgressOnlyGatewayId = k.EgressOnlyGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_GatewayId(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.GatewayId != p.GatewayId {
		p.GatewayId = k.GatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_RouteTableId(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.RouteTableId != p.RouteTableId {
		p.RouteTableId = k.RouteTableId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_TransitGatewayId(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayId != p.TransitGatewayId {
		p.TransitGatewayId = k.TransitGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_DestinationIpv6CidrBlock(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.DestinationIpv6CidrBlock != p.DestinationIpv6CidrBlock {
		p.DestinationIpv6CidrBlock = k.DestinationIpv6CidrBlock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_LocalGatewayId(k *RouteParameters, p *RouteParameters, md *plugin.MergeDescription) bool {
	if k.LocalGatewayId != p.LocalGatewayId {
		p.LocalGatewayId = k.LocalGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeRoute_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeRoute_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeRoute_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRoute_State(k *RouteObservation, p *RouteObservation, md *plugin.MergeDescription) bool {
	if k.State != p.State {
		k.State = p.State
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRoute_Origin(k *RouteObservation, p *RouteObservation, md *plugin.MergeDescription) bool {
	if k.Origin != p.Origin {
		k.Origin = p.Origin
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRoute_DestinationPrefixListId(k *RouteObservation, p *RouteObservation, md *plugin.MergeDescription) bool {
	if k.DestinationPrefixListId != p.DestinationPrefixListId {
		k.DestinationPrefixListId = p.DestinationPrefixListId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRoute_InstanceOwnerId(k *RouteObservation, p *RouteObservation, md *plugin.MergeDescription) bool {
	if k.InstanceOwnerId != p.InstanceOwnerId {
		k.InstanceOwnerId = p.InstanceOwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}