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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*EbsSnapshotCopy)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EbsSnapshotCopy.")
	}
	return EncodeEbsSnapshotCopy(*r), nil
}

func EncodeEbsSnapshotCopy(r EbsSnapshotCopy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEbsSnapshotCopy_SourceSnapshotId(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_Id(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_SourceRegion(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_Description(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_Encrypted(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_VolumeId(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_VolumeSize(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_OwnerAlias(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_Arn(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_DataEncryptionKeyId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeEbsSnapshotCopy_SourceSnapshotId(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["source_snapshot_id"] = cty.StringVal(p.SourceSnapshotId)
}

func EncodeEbsSnapshotCopy_Tags(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEbsSnapshotCopy_Id(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEbsSnapshotCopy_KmsKeyId(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeEbsSnapshotCopy_SourceRegion(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["source_region"] = cty.StringVal(p.SourceRegion)
}

func EncodeEbsSnapshotCopy_Description(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEbsSnapshotCopy_Encrypted(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeEbsSnapshotCopy_VolumeId(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["volume_id"] = cty.StringVal(p.VolumeId)
}

func EncodeEbsSnapshotCopy_VolumeSize(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeEbsSnapshotCopy_OwnerAlias(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["owner_alias"] = cty.StringVal(p.OwnerAlias)
}

func EncodeEbsSnapshotCopy_OwnerId(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeEbsSnapshotCopy_Arn(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEbsSnapshotCopy_DataEncryptionKeyId(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["data_encryption_key_id"] = cty.StringVal(p.DataEncryptionKeyId)
}