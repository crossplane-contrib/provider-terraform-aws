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
	k := kube.(*Vpc)
	p := prov.(*Vpc)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVpc_EnableClassiclink(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_InstanceTenancy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_EnableDnsSupport(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_AssignGeneratedIpv6CidrBlock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_CidrBlock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_EnableClassiclinkDnsSupport(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_EnableDnsHostnames(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_DefaultRouteTableId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_Ipv6AssociationId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_DhcpOptionsId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_DefaultSecurityGroupId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_Ipv6CidrBlock(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_OwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_DefaultNetworkAclId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpc_MainRouteTableId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeVpc_EnableClassiclink(k *VpcParameters, p *VpcParameters, md *plugin.MergeDescription) bool {
	if k.EnableClassiclink != p.EnableClassiclink {
		p.EnableClassiclink = k.EnableClassiclink
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpc_Id(k *VpcParameters, p *VpcParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVpc_Tags(k *VpcParameters, p *VpcParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpc_InstanceTenancy(k *VpcParameters, p *VpcParameters, md *plugin.MergeDescription) bool {
	if k.InstanceTenancy != p.InstanceTenancy {
		p.InstanceTenancy = k.InstanceTenancy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpc_EnableDnsSupport(k *VpcParameters, p *VpcParameters, md *plugin.MergeDescription) bool {
	if k.EnableDnsSupport != p.EnableDnsSupport {
		p.EnableDnsSupport = k.EnableDnsSupport
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpc_AssignGeneratedIpv6CidrBlock(k *VpcParameters, p *VpcParameters, md *plugin.MergeDescription) bool {
	if k.AssignGeneratedIpv6CidrBlock != p.AssignGeneratedIpv6CidrBlock {
		p.AssignGeneratedIpv6CidrBlock = k.AssignGeneratedIpv6CidrBlock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpc_CidrBlock(k *VpcParameters, p *VpcParameters, md *plugin.MergeDescription) bool {
	if k.CidrBlock != p.CidrBlock {
		p.CidrBlock = k.CidrBlock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpc_EnableClassiclinkDnsSupport(k *VpcParameters, p *VpcParameters, md *plugin.MergeDescription) bool {
	if k.EnableClassiclinkDnsSupport != p.EnableClassiclinkDnsSupport {
		p.EnableClassiclinkDnsSupport = k.EnableClassiclinkDnsSupport
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpc_EnableDnsHostnames(k *VpcParameters, p *VpcParameters, md *plugin.MergeDescription) bool {
	if k.EnableDnsHostnames != p.EnableDnsHostnames {
		p.EnableDnsHostnames = k.EnableDnsHostnames
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpc_DefaultRouteTableId(k *VpcObservation, p *VpcObservation, md *plugin.MergeDescription) bool {
	if k.DefaultRouteTableId != p.DefaultRouteTableId {
		k.DefaultRouteTableId = p.DefaultRouteTableId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpc_Ipv6AssociationId(k *VpcObservation, p *VpcObservation, md *plugin.MergeDescription) bool {
	if k.Ipv6AssociationId != p.Ipv6AssociationId {
		k.Ipv6AssociationId = p.Ipv6AssociationId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpc_Arn(k *VpcObservation, p *VpcObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpc_DhcpOptionsId(k *VpcObservation, p *VpcObservation, md *plugin.MergeDescription) bool {
	if k.DhcpOptionsId != p.DhcpOptionsId {
		k.DhcpOptionsId = p.DhcpOptionsId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpc_DefaultSecurityGroupId(k *VpcObservation, p *VpcObservation, md *plugin.MergeDescription) bool {
	if k.DefaultSecurityGroupId != p.DefaultSecurityGroupId {
		k.DefaultSecurityGroupId = p.DefaultSecurityGroupId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpc_Ipv6CidrBlock(k *VpcObservation, p *VpcObservation, md *plugin.MergeDescription) bool {
	if k.Ipv6CidrBlock != p.Ipv6CidrBlock {
		k.Ipv6CidrBlock = p.Ipv6CidrBlock
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpc_OwnerId(k *VpcObservation, p *VpcObservation, md *plugin.MergeDescription) bool {
	if k.OwnerId != p.OwnerId {
		k.OwnerId = p.OwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpc_DefaultNetworkAclId(k *VpcObservation, p *VpcObservation, md *plugin.MergeDescription) bool {
	if k.DefaultNetworkAclId != p.DefaultNetworkAclId {
		k.DefaultNetworkAclId = p.DefaultNetworkAclId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpc_MainRouteTableId(k *VpcObservation, p *VpcObservation, md *plugin.MergeDescription) bool {
	if k.MainRouteTableId != p.MainRouteTableId {
		k.MainRouteTableId = p.MainRouteTableId
		md.StatusUpdated = true
		return true
	}
	return false
}