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
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*VolumeAttachment)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVolumeAttachment(r, ctyValue)
}

func DecodeVolumeAttachment(prev *VolumeAttachment, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVolumeAttachment_InstanceId(&new.Spec.ForProvider, valMap)
	DecodeVolumeAttachment_SkipDestroy(&new.Spec.ForProvider, valMap)
	DecodeVolumeAttachment_VolumeId(&new.Spec.ForProvider, valMap)
	DecodeVolumeAttachment_DeviceName(&new.Spec.ForProvider, valMap)
	DecodeVolumeAttachment_ForceDetach(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVolumeAttachment_InstanceId(p *VolumeAttachmentParameters, vals map[string]cty.Value) {
	p.InstanceId = ctwhy.ValueAsString(vals["instance_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVolumeAttachment_SkipDestroy(p *VolumeAttachmentParameters, vals map[string]cty.Value) {
	p.SkipDestroy = ctwhy.ValueAsBool(vals["skip_destroy"])
}

//primitiveTypeDecodeTemplate
func DecodeVolumeAttachment_VolumeId(p *VolumeAttachmentParameters, vals map[string]cty.Value) {
	p.VolumeId = ctwhy.ValueAsString(vals["volume_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVolumeAttachment_DeviceName(p *VolumeAttachmentParameters, vals map[string]cty.Value) {
	p.DeviceName = ctwhy.ValueAsString(vals["device_name"])
}

//primitiveTypeDecodeTemplate
func DecodeVolumeAttachment_ForceDetach(p *VolumeAttachmentParameters, vals map[string]cty.Value) {
	p.ForceDetach = ctwhy.ValueAsBool(vals["force_detach"])
}