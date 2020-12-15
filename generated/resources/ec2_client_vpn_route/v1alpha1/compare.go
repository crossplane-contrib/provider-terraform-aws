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
	k := kube.(*Ec2ClientVpnRoute)
	p := prov.(*Ec2ClientVpnRoute)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2ClientVpnRoute_TargetVpcSubnetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnRoute_ClientVpnEndpointId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnRoute_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnRoute_DestinationCidrBlock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnRoute_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnRoute_Origin(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2ClientVpnRoute_Type(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEc2ClientVpnRoute_TargetVpcSubnetId(k *Ec2ClientVpnRouteParameters, p *Ec2ClientVpnRouteParameters, md *plugin.MergeDescription) bool {
	if k.TargetVpcSubnetId != p.TargetVpcSubnetId {
		p.TargetVpcSubnetId = k.TargetVpcSubnetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2ClientVpnRoute_ClientVpnEndpointId(k *Ec2ClientVpnRouteParameters, p *Ec2ClientVpnRouteParameters, md *plugin.MergeDescription) bool {
	if k.ClientVpnEndpointId != p.ClientVpnEndpointId {
		p.ClientVpnEndpointId = k.ClientVpnEndpointId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2ClientVpnRoute_Description(k *Ec2ClientVpnRouteParameters, p *Ec2ClientVpnRouteParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2ClientVpnRoute_DestinationCidrBlock(k *Ec2ClientVpnRouteParameters, p *Ec2ClientVpnRouteParameters, md *plugin.MergeDescription) bool {
	if k.DestinationCidrBlock != p.DestinationCidrBlock {
		p.DestinationCidrBlock = k.DestinationCidrBlock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2ClientVpnRoute_Id(k *Ec2ClientVpnRouteParameters, p *Ec2ClientVpnRouteParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2ClientVpnRoute_Origin(k *Ec2ClientVpnRouteObservation, p *Ec2ClientVpnRouteObservation, md *plugin.MergeDescription) bool {
	if k.Origin != p.Origin {
		k.Origin = p.Origin
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2ClientVpnRoute_Type(k *Ec2ClientVpnRouteObservation, p *Ec2ClientVpnRouteObservation, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		k.Type = p.Type
		md.StatusUpdated = true
		return true
	}
	return false
}