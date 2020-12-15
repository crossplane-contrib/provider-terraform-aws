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
	k := kube.(*DxHostedPublicVirtualInterface)
	p := prov.(*DxHostedPublicVirtualInterface)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDxHostedPublicVirtualInterface_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_Vlan(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_RouteFilterPrefixes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_BgpAuthKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_ConnectionId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_AddressFamily(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_BgpAsn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_CustomerAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_OwnerAccountId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_AmazonAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_AwsDevice(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_AmazonSideAsn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDxHostedPublicVirtualInterface_Name(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_Vlan(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Vlan != p.Vlan {
		p.Vlan = k.Vlan
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDxHostedPublicVirtualInterface_RouteFilterPrefixes(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.RouteFilterPrefixes, p.RouteFilterPrefixes) {
		p.RouteFilterPrefixes = k.RouteFilterPrefixes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_BgpAuthKey(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.BgpAuthKey != p.BgpAuthKey {
		p.BgpAuthKey = k.BgpAuthKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_ConnectionId(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.ConnectionId != p.ConnectionId {
		p.ConnectionId = k.ConnectionId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_AddressFamily(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.AddressFamily != p.AddressFamily {
		p.AddressFamily = k.AddressFamily
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_BgpAsn(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.BgpAsn != p.BgpAsn {
		p.BgpAsn = k.BgpAsn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_CustomerAddress(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.CustomerAddress != p.CustomerAddress {
		p.CustomerAddress = k.CustomerAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_OwnerAccountId(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.OwnerAccountId != p.OwnerAccountId {
		p.OwnerAccountId = k.OwnerAccountId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_AmazonAddress(k *DxHostedPublicVirtualInterfaceParameters, p *DxHostedPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.AmazonAddress != p.AmazonAddress {
		p.AmazonAddress = k.AmazonAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDxHostedPublicVirtualInterface_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDxHostedPublicVirtualInterface_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPublicVirtualInterface_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPublicVirtualInterface_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxHostedPublicVirtualInterface_AwsDevice(k *DxHostedPublicVirtualInterfaceObservation, p *DxHostedPublicVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.AwsDevice != p.AwsDevice {
		k.AwsDevice = p.AwsDevice
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxHostedPublicVirtualInterface_AmazonSideAsn(k *DxHostedPublicVirtualInterfaceObservation, p *DxHostedPublicVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.AmazonSideAsn != p.AmazonSideAsn {
		k.AmazonSideAsn = p.AmazonSideAsn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxHostedPublicVirtualInterface_Arn(k *DxHostedPublicVirtualInterfaceObservation, p *DxHostedPublicVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}