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
	k := kube.(*DxPublicVirtualInterface)
	p := prov.(*DxPublicVirtualInterface)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDxPublicVirtualInterface_CustomerAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_Vlan(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_RouteFilterPrefixes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_AmazonAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_BgpAsn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_ConnectionId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_AddressFamily(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_BgpAuthKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_AwsDevice(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_AmazonSideAsn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDxPublicVirtualInterface_CustomerAddress(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.CustomerAddress != p.CustomerAddress {
		p.CustomerAddress = k.CustomerAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_Name(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDxPublicVirtualInterface_Tags(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_Vlan(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Vlan != p.Vlan {
		p.Vlan = k.Vlan
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_Id(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDxPublicVirtualInterface_RouteFilterPrefixes(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.RouteFilterPrefixes, p.RouteFilterPrefixes) {
		p.RouteFilterPrefixes = k.RouteFilterPrefixes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_AmazonAddress(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.AmazonAddress != p.AmazonAddress {
		p.AmazonAddress = k.AmazonAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_BgpAsn(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.BgpAsn != p.BgpAsn {
		p.BgpAsn = k.BgpAsn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_ConnectionId(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.ConnectionId != p.ConnectionId {
		p.ConnectionId = k.ConnectionId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_AddressFamily(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.AddressFamily != p.AddressFamily {
		p.AddressFamily = k.AddressFamily
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_BgpAuthKey(k *DxPublicVirtualInterfaceParameters, p *DxPublicVirtualInterfaceParameters, md *plugin.MergeDescription) bool {
	if k.BgpAuthKey != p.BgpAuthKey {
		p.BgpAuthKey = k.BgpAuthKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDxPublicVirtualInterface_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDxPublicVirtualInterface_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxPublicVirtualInterface_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxPublicVirtualInterface_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxPublicVirtualInterface_AwsDevice(k *DxPublicVirtualInterfaceObservation, p *DxPublicVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.AwsDevice != p.AwsDevice {
		k.AwsDevice = p.AwsDevice
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxPublicVirtualInterface_AmazonSideAsn(k *DxPublicVirtualInterfaceObservation, p *DxPublicVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.AmazonSideAsn != p.AmazonSideAsn {
		k.AmazonSideAsn = p.AmazonSideAsn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxPublicVirtualInterface_Arn(k *DxPublicVirtualInterfaceObservation, p *DxPublicVirtualInterfaceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}