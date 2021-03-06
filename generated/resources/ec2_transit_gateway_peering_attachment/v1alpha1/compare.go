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
	k := kube.(*Ec2TransitGatewayPeeringAttachment)
	p := prov.(*Ec2TransitGatewayPeeringAttachment)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2TransitGatewayPeeringAttachment_PeerAccountId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayPeeringAttachment_PeerRegion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayPeeringAttachment_PeerTransitGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayPeeringAttachment_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TransitGatewayPeeringAttachment_TransitGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeEc2TransitGatewayPeeringAttachment_PeerAccountId(k *Ec2TransitGatewayPeeringAttachmentParameters, p *Ec2TransitGatewayPeeringAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.PeerAccountId != p.PeerAccountId {
		p.PeerAccountId = k.PeerAccountId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayPeeringAttachment_PeerRegion(k *Ec2TransitGatewayPeeringAttachmentParameters, p *Ec2TransitGatewayPeeringAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.PeerRegion != p.PeerRegion {
		p.PeerRegion = k.PeerRegion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayPeeringAttachment_PeerTransitGatewayId(k *Ec2TransitGatewayPeeringAttachmentParameters, p *Ec2TransitGatewayPeeringAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.PeerTransitGatewayId != p.PeerTransitGatewayId {
		p.PeerTransitGatewayId = k.PeerTransitGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEc2TransitGatewayPeeringAttachment_Tags(k *Ec2TransitGatewayPeeringAttachmentParameters, p *Ec2TransitGatewayPeeringAttachmentParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TransitGatewayPeeringAttachment_TransitGatewayId(k *Ec2TransitGatewayPeeringAttachmentParameters, p *Ec2TransitGatewayPeeringAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.TransitGatewayId != p.TransitGatewayId {
		p.TransitGatewayId = k.TransitGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}