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
	r, ok := mr.(*EbsSnapshot)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EbsSnapshot.")
	}
	return EncodeEbsSnapshot(*r), nil
}

func EncodeEbsSnapshot(r EbsSnapshot) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEbsSnapshot_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshot_VolumeId(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshot_Description(r.Spec.ForProvider, ctyVal)
	EncodeEbsSnapshot_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeEbsSnapshot_VolumeSize(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshot_DataEncryptionKeyId(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshot_KmsKeyId(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshot_OwnerAlias(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshot_Arn(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshot_Encrypted(r.Status.AtProvider, ctyVal)
	EncodeEbsSnapshot_OwnerId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeEbsSnapshot_Tags(p EbsSnapshotParameters, vals map[string]cty.Value) {
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

func EncodeEbsSnapshot_VolumeId(p EbsSnapshotParameters, vals map[string]cty.Value) {
	vals["volume_id"] = cty.StringVal(p.VolumeId)
}

func EncodeEbsSnapshot_Description(p EbsSnapshotParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEbsSnapshot_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeEbsSnapshot_Timeouts_Create(p, ctyVal)
	EncodeEbsSnapshot_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeEbsSnapshot_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeEbsSnapshot_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeEbsSnapshot_VolumeSize(p EbsSnapshotObservation, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeEbsSnapshot_DataEncryptionKeyId(p EbsSnapshotObservation, vals map[string]cty.Value) {
	vals["data_encryption_key_id"] = cty.StringVal(p.DataEncryptionKeyId)
}

func EncodeEbsSnapshot_KmsKeyId(p EbsSnapshotObservation, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeEbsSnapshot_OwnerAlias(p EbsSnapshotObservation, vals map[string]cty.Value) {
	vals["owner_alias"] = cty.StringVal(p.OwnerAlias)
}

func EncodeEbsSnapshot_Arn(p EbsSnapshotObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEbsSnapshot_Encrypted(p EbsSnapshotObservation, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeEbsSnapshot_OwnerId(p EbsSnapshotObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}