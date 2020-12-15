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
	r, ok := mr.(*EbsSnapshot)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEbsSnapshot(r, ctyValue)
}

func DecodeEbsSnapshot(prev *EbsSnapshot, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEbsSnapshot_Tags(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshot_Description(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshot_VolumeId(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshot_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeEbsSnapshot_Arn(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshot_KmsKeyId(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshot_OwnerAlias(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshot_OwnerId(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshot_VolumeSize(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshot_DataEncryptionKeyId(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshot_Encrypted(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeEbsSnapshot_Tags(p *EbsSnapshotParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_Description(p *EbsSnapshotParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_VolumeId(p *EbsSnapshotParameters, vals map[string]cty.Value) {
	p.VolumeId = ctwhy.ValueAsString(vals["volume_id"])
}

//containerTypeDecodeTemplate
func DecodeEbsSnapshot_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeEbsSnapshot_Timeouts_Create(p, valMap)
	DecodeEbsSnapshot_Timeouts_Delete(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_Arn(p *EbsSnapshotObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_KmsKeyId(p *EbsSnapshotObservation, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_OwnerAlias(p *EbsSnapshotObservation, vals map[string]cty.Value) {
	p.OwnerAlias = ctwhy.ValueAsString(vals["owner_alias"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_OwnerId(p *EbsSnapshotObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_VolumeSize(p *EbsSnapshotObservation, vals map[string]cty.Value) {
	p.VolumeSize = ctwhy.ValueAsInt64(vals["volume_size"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_DataEncryptionKeyId(p *EbsSnapshotObservation, vals map[string]cty.Value) {
	p.DataEncryptionKeyId = ctwhy.ValueAsString(vals["data_encryption_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEbsSnapshot_Encrypted(p *EbsSnapshotObservation, vals map[string]cty.Value) {
	p.Encrypted = ctwhy.ValueAsBool(vals["encrypted"])
}