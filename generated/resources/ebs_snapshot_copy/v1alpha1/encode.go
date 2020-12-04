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
	"github.com/zclconf/go-cty/cty"
)

func EncodeEbsSnapshotCopy(r EbsSnapshotCopy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEbsSnapshotCopy_Description(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_Encrypted(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_SourceRegion(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_SourceSnapshotId(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_Id(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshotCopy_OwnerAlias(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_VolumeSize(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_Arn(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_DataEncryptionKeyId(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshotCopy_VolumeId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEbsSnapshotCopy_Description(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEbsSnapshotCopy_Encrypted(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeEbsSnapshotCopy_KmsKeyId(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeEbsSnapshotCopy_SourceRegion(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["source_region"] = cty.StringVal(p.SourceRegion)
}

func EncodeEbsSnapshotCopy_SourceSnapshotId(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["source_snapshot_id"] = cty.StringVal(p.SourceSnapshotId)
}

func EncodeEbsSnapshotCopy_Id(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEbsSnapshotCopy_Tags(p EbsSnapshotCopyParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEbsSnapshotCopy_OwnerAlias(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["owner_alias"] = cty.StringVal(p.OwnerAlias)
}

func EncodeEbsSnapshotCopy_VolumeSize(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeEbsSnapshotCopy_Arn(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEbsSnapshotCopy_DataEncryptionKeyId(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["data_encryption_key_id"] = cty.StringVal(p.DataEncryptionKeyId)
}

func EncodeEbsSnapshotCopy_OwnerId(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeEbsSnapshotCopy_VolumeId(p EbsSnapshotCopyObservation, vals map[string]cty.Value) {
	vals["volume_id"] = cty.StringVal(p.VolumeId)
}