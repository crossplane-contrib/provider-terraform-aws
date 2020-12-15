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
	k := kube.(*DxBgpPeer)
	p := prov.(*DxBgpPeer)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDxBgpPeer_BgpAuthKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_VirtualInterfaceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_AddressFamily(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_AmazonAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_BgpAsn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_CustomerAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_BgpStatus(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_AwsDevice(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_BgpPeerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDxBgpPeer_BgpAuthKey(k *DxBgpPeerParameters, p *DxBgpPeerParameters, md *plugin.MergeDescription) bool {
	if k.BgpAuthKey != p.BgpAuthKey {
		p.BgpAuthKey = k.BgpAuthKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxBgpPeer_VirtualInterfaceId(k *DxBgpPeerParameters, p *DxBgpPeerParameters, md *plugin.MergeDescription) bool {
	if k.VirtualInterfaceId != p.VirtualInterfaceId {
		p.VirtualInterfaceId = k.VirtualInterfaceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxBgpPeer_AddressFamily(k *DxBgpPeerParameters, p *DxBgpPeerParameters, md *plugin.MergeDescription) bool {
	if k.AddressFamily != p.AddressFamily {
		p.AddressFamily = k.AddressFamily
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxBgpPeer_AmazonAddress(k *DxBgpPeerParameters, p *DxBgpPeerParameters, md *plugin.MergeDescription) bool {
	if k.AmazonAddress != p.AmazonAddress {
		p.AmazonAddress = k.AmazonAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxBgpPeer_BgpAsn(k *DxBgpPeerParameters, p *DxBgpPeerParameters, md *plugin.MergeDescription) bool {
	if k.BgpAsn != p.BgpAsn {
		p.BgpAsn = k.BgpAsn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxBgpPeer_CustomerAddress(k *DxBgpPeerParameters, p *DxBgpPeerParameters, md *plugin.MergeDescription) bool {
	if k.CustomerAddress != p.CustomerAddress {
		p.CustomerAddress = k.CustomerAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxBgpPeer_Id(k *DxBgpPeerParameters, p *DxBgpPeerParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDxBgpPeer_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDxBgpPeer_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxBgpPeer_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDxBgpPeer_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxBgpPeer_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxBgpPeer_BgpStatus(k *DxBgpPeerObservation, p *DxBgpPeerObservation, md *plugin.MergeDescription) bool {
	if k.BgpStatus != p.BgpStatus {
		k.BgpStatus = p.BgpStatus
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxBgpPeer_AwsDevice(k *DxBgpPeerObservation, p *DxBgpPeerObservation, md *plugin.MergeDescription) bool {
	if k.AwsDevice != p.AwsDevice {
		k.AwsDevice = p.AwsDevice
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxBgpPeer_BgpPeerId(k *DxBgpPeerObservation, p *DxBgpPeerObservation, md *plugin.MergeDescription) bool {
	if k.BgpPeerId != p.BgpPeerId {
		k.BgpPeerId = p.BgpPeerId
		md.StatusUpdated = true
		return true
	}
	return false
}