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
	k := kube.(*DxPrivateVirtualInterface)
	p := prov.(*DxPrivateVirtualInterface)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDxPrivateVirtualInterface_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_VpnGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_AddressFamily(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_AmazonAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_ConnectionId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_DxGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_Mtu(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_Vlan(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_BgpAsn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_BgpAuthKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_CustomerAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_AmazonSideAsn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_AwsDevice(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_JumboFrameCapable(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeDxPrivateVirtualInterface_Tags(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_VpnGatewayId(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.VpnGatewayId != p.VpnGatewayId {
		p.VpnGatewayId = k.VpnGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_AddressFamily(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.AddressFamily != p.AddressFamily {
		p.AddressFamily = k.AddressFamily
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_AmazonAddress(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.AmazonAddress != p.AmazonAddress {
		p.AmazonAddress = k.AmazonAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_ConnectionId(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.ConnectionId != p.ConnectionId {
		p.ConnectionId = k.ConnectionId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_DxGatewayId(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.DxGatewayId != p.DxGatewayId {
		p.DxGatewayId = k.DxGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_Mtu(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Mtu != p.Mtu {
		p.Mtu = k.Mtu
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_Vlan(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Vlan != p.Vlan {
		p.Vlan = k.Vlan
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_BgpAsn(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.BgpAsn != p.BgpAsn {
		p.BgpAsn = k.BgpAsn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_BgpAuthKey(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.BgpAuthKey != p.BgpAuthKey {
		p.BgpAuthKey = k.BgpAuthKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_CustomerAddress(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.CustomerAddress != p.CustomerAddress {
		p.CustomerAddress = k.CustomerAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_Name(k *DxPrivateVirtualInterfaceParameters, p *DxPrivateVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDxPrivateVirtualInterface_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDxPrivateVirtualInterface_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPrivateVirtualInterface_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPrivateVirtualInterface_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxPrivateVirtualInterface_AmazonSideAsn(k *DxPrivateVirtualInterfaceObservation, p *DxPrivateVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.AmazonSideAsn != p.AmazonSideAsn {
		k.AmazonSideAsn = p.AmazonSideAsn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxPrivateVirtualInterface_Arn(k *DxPrivateVirtualInterfaceObservation, p *DxPrivateVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxPrivateVirtualInterface_AwsDevice(k *DxPrivateVirtualInterfaceObservation, p *DxPrivateVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.AwsDevice != p.AwsDevice {
		k.AwsDevice = p.AwsDevice
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxPrivateVirtualInterface_JumboFrameCapable(k *DxPrivateVirtualInterfaceObservation, p *DxPrivateVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.JumboFrameCapable != p.JumboFrameCapable {
		k.JumboFrameCapable = p.JumboFrameCapable
		md.StatusUpdated = true
		return true
	}
	return false
}