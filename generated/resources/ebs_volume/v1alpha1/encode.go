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
	r, ok := mr.(*EbsVolume)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EbsVolume.")
	}
	return EncodeEbsVolume(*r), nil
}

func EncodeEbsVolume(r EbsVolume) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEbsVolume_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_Id(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_OutpostArn(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_Size(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_SnapshotId(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_Type(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_Encrypted(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_Iops(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_MultiAttachEnabled(r.Spec.ForProvider, ctyVal)
	EncodeEbsVolume_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEbsVolume_AvailabilityZone(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeEbsVolume_Id(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEbsVolume_OutpostArn(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["outpost_arn"] = cty.StringVal(p.OutpostArn)
}

func EncodeEbsVolume_Size(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeEbsVolume_SnapshotId(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeEbsVolume_Tags(p EbsVolumeParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEbsVolume_Type(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEbsVolume_Encrypted(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeEbsVolume_Iops(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeEbsVolume_KmsKeyId(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeEbsVolume_MultiAttachEnabled(p EbsVolumeParameters, vals map[string]cty.Value) {
	vals["multi_attach_enabled"] = cty.BoolVal(p.MultiAttachEnabled)
}

func EncodeEbsVolume_Arn(p EbsVolumeObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}