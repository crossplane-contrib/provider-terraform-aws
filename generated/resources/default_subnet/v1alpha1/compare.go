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
	k := kube.(*DefaultSubnet)
	p := prov.(*DefaultSubnet)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDefaultSubnet_OutpostArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_AvailabilityZone(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_MapPublicIpOnLaunch(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_Ipv6CidrBlock(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_OwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_VpcId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_CidrBlock(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_AssignIpv6AddressOnCreation(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_AvailabilityZoneId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_Ipv6CidrBlockAssociationId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDefaultSubnet_OutpostArn(k *DefaultSubnetParameters, p *DefaultSubnetParameters, md *plugin.MergeDescription) bool {
	if k.OutpostArn != p.OutpostArn {
		p.OutpostArn = k.OutpostArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDefaultSubnet_AvailabilityZone(k *DefaultSubnetParameters, p *DefaultSubnetParameters, md *plugin.MergeDescription) bool {
	if k.AvailabilityZone != p.AvailabilityZone {
		p.AvailabilityZone = k.AvailabilityZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDefaultSubnet_MapPublicIpOnLaunch(k *DefaultSubnetParameters, p *DefaultSubnetParameters, md *plugin.MergeDescription) bool {
	if k.MapPublicIpOnLaunch != p.MapPublicIpOnLaunch {
		p.MapPublicIpOnLaunch = k.MapPublicIpOnLaunch
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDefaultSubnet_Tags(k *DefaultSubnetParameters, p *DefaultSubnetParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDefaultSubnet_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDefaultSubnet_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDefaultSubnet_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDefaultSubnet_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDefaultSubnet_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDefaultSubnet_Ipv6CidrBlock(k *DefaultSubnetObservation, p *DefaultSubnetObservation, md *plugin.MergeDescription) bool {
	if k.Ipv6CidrBlock != p.Ipv6CidrBlock {
		k.Ipv6CidrBlock = p.Ipv6CidrBlock
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDefaultSubnet_OwnerId(k *DefaultSubnetObservation, p *DefaultSubnetObservation, md *plugin.MergeDescription) bool {
	if k.OwnerId != p.OwnerId {
		k.OwnerId = p.OwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDefaultSubnet_VpcId(k *DefaultSubnetObservation, p *DefaultSubnetObservation, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		k.VpcId = p.VpcId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDefaultSubnet_CidrBlock(k *DefaultSubnetObservation, p *DefaultSubnetObservation, md *plugin.MergeDescription) bool {
	if k.CidrBlock != p.CidrBlock {
		k.CidrBlock = p.CidrBlock
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDefaultSubnet_AssignIpv6AddressOnCreation(k *DefaultSubnetObservation, p *DefaultSubnetObservation, md *plugin.MergeDescription) bool {
	if k.AssignIpv6AddressOnCreation != p.AssignIpv6AddressOnCreation {
		k.AssignIpv6AddressOnCreation = p.AssignIpv6AddressOnCreation
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDefaultSubnet_AvailabilityZoneId(k *DefaultSubnetObservation, p *DefaultSubnetObservation, md *plugin.MergeDescription) bool {
	if k.AvailabilityZoneId != p.AvailabilityZoneId {
		k.AvailabilityZoneId = p.AvailabilityZoneId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDefaultSubnet_Ipv6CidrBlockAssociationId(k *DefaultSubnetObservation, p *DefaultSubnetObservation, md *plugin.MergeDescription) bool {
	if k.Ipv6CidrBlockAssociationId != p.Ipv6CidrBlockAssociationId {
		k.Ipv6CidrBlockAssociationId = p.Ipv6CidrBlockAssociationId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDefaultSubnet_Arn(k *DefaultSubnetObservation, p *DefaultSubnetObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}