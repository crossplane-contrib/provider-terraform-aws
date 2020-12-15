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
	k := kube.(*LightsailStaticIpAttachment)
	p := prov.(*LightsailStaticIpAttachment)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeLightsailStaticIpAttachment_InstanceName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailStaticIpAttachment_StaticIpName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailStaticIpAttachment_IpAddress(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeLightsailStaticIpAttachment_InstanceName(k *LightsailStaticIpAttachmentParameters, p *LightsailStaticIpAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.InstanceName != p.InstanceName {
		p.InstanceName = k.InstanceName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLightsailStaticIpAttachment_StaticIpName(k *LightsailStaticIpAttachmentParameters, p *LightsailStaticIpAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.StaticIpName != p.StaticIpName {
		p.StaticIpName = k.StaticIpName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailStaticIpAttachment_IpAddress(k *LightsailStaticIpAttachmentObservation, p *LightsailStaticIpAttachmentObservation, md *plugin.MergeDescription) bool {
	if k.IpAddress != p.IpAddress {
		k.IpAddress = p.IpAddress
		md.StatusUpdated = true
		return true
	}
	return false
}