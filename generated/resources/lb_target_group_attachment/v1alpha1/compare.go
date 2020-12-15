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
	k := kube.(*LbTargetGroupAttachment)
	p := prov.(*LbTargetGroupAttachment)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeLbTargetGroupAttachment_AvailabilityZone(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLbTargetGroupAttachment_Port(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLbTargetGroupAttachment_TargetGroupArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLbTargetGroupAttachment_TargetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeLbTargetGroupAttachment_AvailabilityZone(k *LbTargetGroupAttachmentParameters, p *LbTargetGroupAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.AvailabilityZone != p.AvailabilityZone {
		p.AvailabilityZone = k.AvailabilityZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLbTargetGroupAttachment_Port(k *LbTargetGroupAttachmentParameters, p *LbTargetGroupAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		p.Port = k.Port
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLbTargetGroupAttachment_TargetGroupArn(k *LbTargetGroupAttachmentParameters, p *LbTargetGroupAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.TargetGroupArn != p.TargetGroupArn {
		p.TargetGroupArn = k.TargetGroupArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLbTargetGroupAttachment_TargetId(k *LbTargetGroupAttachmentParameters, p *LbTargetGroupAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.TargetId != p.TargetId {
		p.TargetId = k.TargetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}