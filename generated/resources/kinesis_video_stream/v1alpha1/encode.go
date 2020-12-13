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
	r, ok := mr.(*KinesisVideoStream)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a KinesisVideoStream.")
	}
	return EncodeKinesisVideoStream(*r), nil
}

func EncodeKinesisVideoStream(r KinesisVideoStream) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKinesisVideoStream_DataRetentionInHours(r.Spec.ForProvider, ctyVal)
	EncodeKinesisVideoStream_Id(r.Spec.ForProvider, ctyVal)
	EncodeKinesisVideoStream_Name(r.Spec.ForProvider, ctyVal)
	EncodeKinesisVideoStream_DeviceName(r.Spec.ForProvider, ctyVal)
	EncodeKinesisVideoStream_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeKinesisVideoStream_MediaType(r.Spec.ForProvider, ctyVal)
	EncodeKinesisVideoStream_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKinesisVideoStream_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeKinesisVideoStream_Version(r.Status.AtProvider, ctyVal)
	EncodeKinesisVideoStream_Arn(r.Status.AtProvider, ctyVal)
	EncodeKinesisVideoStream_CreationTime(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeKinesisVideoStream_DataRetentionInHours(p KinesisVideoStreamParameters, vals map[string]cty.Value) {
	vals["data_retention_in_hours"] = cty.NumberIntVal(p.DataRetentionInHours)
}

func EncodeKinesisVideoStream_Id(p KinesisVideoStreamParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKinesisVideoStream_Name(p KinesisVideoStreamParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKinesisVideoStream_DeviceName(p KinesisVideoStreamParameters, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeKinesisVideoStream_KmsKeyId(p KinesisVideoStreamParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeKinesisVideoStream_MediaType(p KinesisVideoStreamParameters, vals map[string]cty.Value) {
	vals["media_type"] = cty.StringVal(p.MediaType)
}

func EncodeKinesisVideoStream_Tags(p KinesisVideoStreamParameters, vals map[string]cty.Value) {
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

func EncodeKinesisVideoStream_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeKinesisVideoStream_Timeouts_Create(p, ctyVal)
	EncodeKinesisVideoStream_Timeouts_Delete(p, ctyVal)
	EncodeKinesisVideoStream_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeKinesisVideoStream_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeKinesisVideoStream_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeKinesisVideoStream_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeKinesisVideoStream_Version(p KinesisVideoStreamObservation, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeKinesisVideoStream_Arn(p KinesisVideoStreamObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeKinesisVideoStream_CreationTime(p KinesisVideoStreamObservation, vals map[string]cty.Value) {
	vals["creation_time"] = cty.StringVal(p.CreationTime)
}