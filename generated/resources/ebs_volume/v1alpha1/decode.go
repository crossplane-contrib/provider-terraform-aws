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
	r, ok := mr.(*EbsVolume)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEbsVolume(r, ctyValue)
}

func DecodeEbsVolume(prev *EbsVolume, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEbsVolume_Tags(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_Type(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_Iops(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_MultiAttachEnabled(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_Size(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_SnapshotId(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_Encrypted(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_KmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_OutpostArn(&new.Spec.ForProvider, valMap)
	DecodeEbsVolume_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeEbsVolume_Tags(p *EbsVolumeParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_Type(p *EbsVolumeParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_Iops(p *EbsVolumeParameters, vals map[string]cty.Value) {
	p.Iops = ctwhy.ValueAsInt64(vals["iops"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_MultiAttachEnabled(p *EbsVolumeParameters, vals map[string]cty.Value) {
	p.MultiAttachEnabled = ctwhy.ValueAsBool(vals["multi_attach_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_Size(p *EbsVolumeParameters, vals map[string]cty.Value) {
	p.Size = ctwhy.ValueAsInt64(vals["size"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_SnapshotId(p *EbsVolumeParameters, vals map[string]cty.Value) {
	p.SnapshotId = ctwhy.ValueAsString(vals["snapshot_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_AvailabilityZone(p *EbsVolumeParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_Encrypted(p *EbsVolumeParameters, vals map[string]cty.Value) {
	p.Encrypted = ctwhy.ValueAsBool(vals["encrypted"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_KmsKeyId(p *EbsVolumeParameters, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_OutpostArn(p *EbsVolumeParameters, vals map[string]cty.Value) {
	p.OutpostArn = ctwhy.ValueAsString(vals["outpost_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsVolume_Arn(p *EbsVolumeObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}