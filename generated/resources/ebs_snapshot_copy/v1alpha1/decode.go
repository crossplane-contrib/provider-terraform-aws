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
	r, ok := mr.(*EbsSnapshotCopy)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEbsSnapshotCopy(r, ctyValue)
}

func DecodeEbsSnapshotCopy(prev *EbsSnapshotCopy, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEbsSnapshotCopy_SourceSnapshotId(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshotCopy_Tags(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshotCopy_Id(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshotCopy_KmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshotCopy_SourceRegion(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshotCopy_Description(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshotCopy_Encrypted(&new.Spec.ForProvider, valMap)
	DecodeEbsSnapshotCopy_VolumeId(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshotCopy_VolumeSize(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshotCopy_OwnerAlias(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshotCopy_OwnerId(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshotCopy_Arn(&new.Status.AtProvider, valMap)
	DecodeEbsSnapshotCopy_DataEncryptionKeyId(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeEbsSnapshotCopy_SourceSnapshotId(p *EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	p.SourceSnapshotId = ctwhy.ValueAsString(vals["source_snapshot_id"])
}

func DecodeEbsSnapshotCopy_Tags(p *EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeEbsSnapshotCopy_Id(p *EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeEbsSnapshotCopy_KmsKeyId(p *EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

func DecodeEbsSnapshotCopy_SourceRegion(p *EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	p.SourceRegion = ctwhy.ValueAsString(vals["source_region"])
}

func DecodeEbsSnapshotCopy_Description(p *EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeEbsSnapshotCopy_Encrypted(p *EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	p.Encrypted = ctwhy.ValueAsBool(vals["encrypted"])
}

func DecodeEbsSnapshotCopy_VolumeId(p *EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	p.VolumeId = ctwhy.ValueAsString(vals["volume_id"])
}

func DecodeEbsSnapshotCopy_VolumeSize(p *EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	p.VolumeSize = ctwhy.ValueAsInt64(vals["volume_size"])
}

func DecodeEbsSnapshotCopy_OwnerAlias(p *EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	p.OwnerAlias = ctwhy.ValueAsString(vals["owner_alias"])
}

func DecodeEbsSnapshotCopy_OwnerId(p *EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

func DecodeEbsSnapshotCopy_Arn(p *EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeEbsSnapshotCopy_DataEncryptionKeyId(p *EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	p.DataEncryptionKeyId = ctwhy.ValueAsString(vals["data_encryption_key_id"])
}