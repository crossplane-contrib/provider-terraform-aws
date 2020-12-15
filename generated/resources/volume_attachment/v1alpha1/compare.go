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
	k := kube.(*VolumeAttachment)
	p := prov.(*VolumeAttachment)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVolumeAttachment_InstanceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVolumeAttachment_SkipDestroy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVolumeAttachment_VolumeId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVolumeAttachment_DeviceName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVolumeAttachment_ForceDetach(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVolumeAttachment_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeVolumeAttachment_InstanceId(k *VolumeAttachmentParameters, p *VolumeAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.InstanceId != p.InstanceId {
		p.InstanceId = k.InstanceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVolumeAttachment_SkipDestroy(k *VolumeAttachmentParameters, p *VolumeAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.SkipDestroy != p.SkipDestroy {
		p.SkipDestroy = k.SkipDestroy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVolumeAttachment_VolumeId(k *VolumeAttachmentParameters, p *VolumeAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.VolumeId != p.VolumeId {
		p.VolumeId = k.VolumeId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVolumeAttachment_DeviceName(k *VolumeAttachmentParameters, p *VolumeAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.DeviceName != p.DeviceName {
		p.DeviceName = k.DeviceName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVolumeAttachment_ForceDetach(k *VolumeAttachmentParameters, p *VolumeAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.ForceDetach != p.ForceDetach {
		p.ForceDetach = k.ForceDetach
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVolumeAttachment_Id(k *VolumeAttachmentParameters, p *VolumeAttachmentParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}