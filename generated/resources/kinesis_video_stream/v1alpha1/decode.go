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
	r, ok := mr.(*KinesisVideoStream)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeKinesisVideoStream(r, ctyValue)
}

func DecodeKinesisVideoStream(prev *KinesisVideoStream, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeKinesisVideoStream_KmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeKinesisVideoStream_Tags(&new.Spec.ForProvider, valMap)
	DecodeKinesisVideoStream_DataRetentionInHours(&new.Spec.ForProvider, valMap)
	DecodeKinesisVideoStream_DeviceName(&new.Spec.ForProvider, valMap)
	DecodeKinesisVideoStream_MediaType(&new.Spec.ForProvider, valMap)
	DecodeKinesisVideoStream_Name(&new.Spec.ForProvider, valMap)
	DecodeKinesisVideoStream_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeKinesisVideoStream_Arn(&new.Status.AtProvider, valMap)
	DecodeKinesisVideoStream_CreationTime(&new.Status.AtProvider, valMap)
	DecodeKinesisVideoStream_Version(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_KmsKeyId(p *KinesisVideoStreamParameters, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeKinesisVideoStream_Tags(p *KinesisVideoStreamParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_DataRetentionInHours(p *KinesisVideoStreamParameters, vals map[string]cty.Value) {
	p.DataRetentionInHours = ctwhy.ValueAsInt64(vals["data_retention_in_hours"])
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_DeviceName(p *KinesisVideoStreamParameters, vals map[string]cty.Value) {
	p.DeviceName = ctwhy.ValueAsString(vals["device_name"])
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_MediaType(p *KinesisVideoStreamParameters, vals map[string]cty.Value) {
	p.MediaType = ctwhy.ValueAsString(vals["media_type"])
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_Name(p *KinesisVideoStreamParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//containerTypeDecodeTemplate
func DecodeKinesisVideoStream_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeKinesisVideoStream_Timeouts_Create(p, valMap)
	DecodeKinesisVideoStream_Timeouts_Delete(p, valMap)
	DecodeKinesisVideoStream_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_Arn(p *KinesisVideoStreamObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_CreationTime(p *KinesisVideoStreamObservation, vals map[string]cty.Value) {
	p.CreationTime = ctwhy.ValueAsString(vals["creation_time"])
}

//primitiveTypeDecodeTemplate
func DecodeKinesisVideoStream_Version(p *KinesisVideoStreamObservation, vals map[string]cty.Value) {
	p.Version = ctwhy.ValueAsString(vals["version"])
}