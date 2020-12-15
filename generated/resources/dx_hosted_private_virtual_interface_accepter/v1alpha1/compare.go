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
	k := kube.(*DxHostedPrivateVirtualInterfaceAccepter)
	p := prov.(*DxHostedPrivateVirtualInterfaceAccepter)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDxHostedPrivateVirtualInterfaceAccepter_DxGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPrivateVirtualInterfaceAccepter_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPrivateVirtualInterfaceAccepter_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPrivateVirtualInterfaceAccepter_VirtualInterfaceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPrivateVirtualInterfaceAccepter_VpnGatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPrivateVirtualInterfaceAccepter_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPrivateVirtualInterfaceAccepter_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDxHostedPrivateVirtualInterfaceAccepter_DxGatewayId(k *DxHostedPrivateVirtualInterfaceAccepterParameters, p *DxHostedPrivateVirtualInterfaceAccepterParameters, md *plugin.MergeDescription) bool {
	if k.DxGatewayId != p.DxGatewayId {
		p.DxGatewayId = k.DxGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPrivateVirtualInterfaceAccepter_Id(k *DxHostedPrivateVirtualInterfaceAccepterParameters, p *DxHostedPrivateVirtualInterfaceAccepterParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDxHostedPrivateVirtualInterfaceAccepter_Tags(k *DxHostedPrivateVirtualInterfaceAccepterParameters, p *DxHostedPrivateVirtualInterfaceAccepterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPrivateVirtualInterfaceAccepter_VirtualInterfaceId(k *DxHostedPrivateVirtualInterfaceAccepterParameters, p *DxHostedPrivateVirtualInterfaceAccepterParameters, md *plugin.MergeDescription) bool {
	if k.VirtualInterfaceId != p.VirtualInterfaceId {
		p.VirtualInterfaceId = k.VirtualInterfaceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPrivateVirtualInterfaceAccepter_VpnGatewayId(k *DxHostedPrivateVirtualInterfaceAccepterParameters, p *DxHostedPrivateVirtualInterfaceAccepterParameters, md *plugin.MergeDescription) bool {
	if k.VpnGatewayId != p.VpnGatewayId {
		p.VpnGatewayId = k.VpnGatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDxHostedPrivateVirtualInterfaceAccepter_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDxHostedPrivateVirtualInterfaceAccepter_Arn(k *DxHostedPrivateVirtualInterfaceAccepterObservation, p *DxHostedPrivateVirtualInterfaceAccepterObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}