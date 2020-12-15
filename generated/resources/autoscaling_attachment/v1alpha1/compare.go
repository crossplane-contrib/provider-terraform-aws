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
	k := kube.(*AutoscalingAttachment)
	p := prov.(*AutoscalingAttachment)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeAutoscalingAttachment_AlbTargetGroupArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingAttachment_AutoscalingGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingAttachment_Elb(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingAttachment_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeAutoscalingAttachment_AlbTargetGroupArn(k *AutoscalingAttachmentParameters, p *AutoscalingAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.AlbTargetGroupArn != p.AlbTargetGroupArn {
		p.AlbTargetGroupArn = k.AlbTargetGroupArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingAttachment_AutoscalingGroupName(k *AutoscalingAttachmentParameters, p *AutoscalingAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.AutoscalingGroupName != p.AutoscalingGroupName {
		p.AutoscalingGroupName = k.AutoscalingGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingAttachment_Elb(k *AutoscalingAttachmentParameters, p *AutoscalingAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.Elb != p.Elb {
		p.Elb = k.Elb
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingAttachment_Id(k *AutoscalingAttachmentParameters, p *AutoscalingAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}