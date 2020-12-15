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
	k := kube.(*DxTransitVirtualInterface)
	p := prov.(*DxTransitVirtualInterface)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDxTransitVirtualInterface_Mtu(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_AddressFamily(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_AmazonAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_BgpAsn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_CustomerAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_DxGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_BgpAuthKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_ConnectionId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_Vlan(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_AmazonSideAsn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_AwsDevice(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_JumboFrameCapable(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDxTransitVirtualInterface_Mtu(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Mtu != p.Mtu {
		p.Mtu = k.Mtu
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_AddressFamily(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.AddressFamily != p.AddressFamily {
		p.AddressFamily = k.AddressFamily
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_AmazonAddress(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.AmazonAddress != p.AmazonAddress {
		p.AmazonAddress = k.AmazonAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_BgpAsn(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.BgpAsn != p.BgpAsn {
		p.BgpAsn = k.BgpAsn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDxTransitVirtualInterface_Tags(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_CustomerAddress(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.CustomerAddress != p.CustomerAddress {
		p.CustomerAddress = k.CustomerAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_DxGatewayId(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.DxGatewayId != p.DxGatewayId {
		p.DxGatewayId = k.DxGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_Name(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_BgpAuthKey(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.BgpAuthKey != p.BgpAuthKey {
		p.BgpAuthKey = k.BgpAuthKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_ConnectionId(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.ConnectionId != p.ConnectionId {
		p.ConnectionId = k.ConnectionId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_Vlan(k *DxTransitVirtualInterfaceParameters, p *DxTransitVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Vlan != p.Vlan {
		p.Vlan = k.Vlan
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDxTransitVirtualInterface_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDxTransitVirtualInterface_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxTransitVirtualInterface_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxTransitVirtualInterface_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxTransitVirtualInterface_Arn(k *DxTransitVirtualInterfaceObservation, p *DxTransitVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxTransitVirtualInterface_AmazonSideAsn(k *DxTransitVirtualInterfaceObservation, p *DxTransitVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.AmazonSideAsn != p.AmazonSideAsn {
		k.AmazonSideAsn = p.AmazonSideAsn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxTransitVirtualInterface_AwsDevice(k *DxTransitVirtualInterfaceObservation, p *DxTransitVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.AwsDevice != p.AwsDevice {
		k.AwsDevice = p.AwsDevice
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxTransitVirtualInterface_JumboFrameCapable(k *DxTransitVirtualInterfaceObservation, p *DxTransitVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.JumboFrameCapable != p.JumboFrameCapable {
		k.JumboFrameCapable = p.JumboFrameCapable
		md.StatusUpdated = true
		return true
	}
	return false
}