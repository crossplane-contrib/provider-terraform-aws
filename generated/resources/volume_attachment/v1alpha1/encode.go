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
	
	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*VolumeAttachment)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VolumeAttachment.")
	}
	return EncodeVolumeAttachment(*r), nil
}

func EncodeVolumeAttachment(r VolumeAttachment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVolumeAttachment_InstanceId(r.Spec.ForProvider, ctyVal)
	EncodeVolumeAttachment_SkipDestroy(r.Spec.ForProvider, ctyVal)
	EncodeVolumeAttachment_VolumeId(r.Spec.ForProvider, ctyVal)
	EncodeVolumeAttachment_DeviceName(r.Spec.ForProvider, ctyVal)
	EncodeVolumeAttachment_ForceDetach(r.Spec.ForProvider, ctyVal)
	EncodeVolumeAttachment_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeVolumeAttachment_InstanceId(p VolumeAttachmentParameters, vals map[string]cty.Value) {
	vals["instance_id"] = cty.StringVal(p.InstanceId)
}

func EncodeVolumeAttachment_SkipDestroy(p VolumeAttachmentParameters, vals map[string]cty.Value) {
	vals["skip_destroy"] = cty.BoolVal(p.SkipDestroy)
}

func EncodeVolumeAttachment_VolumeId(p VolumeAttachmentParameters, vals map[string]cty.Value) {
	vals["volume_id"] = cty.StringVal(p.VolumeId)
}

func EncodeVolumeAttachment_DeviceName(p VolumeAttachmentParameters, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeVolumeAttachment_ForceDetach(p VolumeAttachmentParameters, vals map[string]cty.Value) {
	vals["force_detach"] = cty.BoolVal(p.ForceDetach)
}

func EncodeVolumeAttachment_Id(p VolumeAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}