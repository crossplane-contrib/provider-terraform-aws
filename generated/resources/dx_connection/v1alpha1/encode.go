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
	r, ok := mr.(*DxConnection)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DxConnection.")
	}
	return EncodeDxConnection(*r), nil
}

func EncodeDxConnection(r DxConnection) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxConnection_Location(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_Bandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_Arn(r.Status.AtProvider, ctyVal)
	EncodeDxConnection_AwsDevice(r.Status.AtProvider, ctyVal)
	EncodeDxConnection_HasLogicalRedundancy(r.Status.AtProvider, ctyVal)
	EncodeDxConnection_JumboFrameCapable(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDxConnection_Location(p DxConnectionParameters, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeDxConnection_Name(p DxConnectionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxConnection_Id(p DxConnectionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxConnection_Bandwidth(p DxConnectionParameters, vals map[string]cty.Value) {
	vals["bandwidth"] = cty.StringVal(p.Bandwidth)
}

func EncodeDxConnection_Tags(p DxConnectionParameters, vals map[string]cty.Value) {
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

func EncodeDxConnection_Arn(p DxConnectionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDxConnection_AwsDevice(p DxConnectionObservation, vals map[string]cty.Value) {
	vals["aws_device"] = cty.StringVal(p.AwsDevice)
}

func EncodeDxConnection_HasLogicalRedundancy(p DxConnectionObservation, vals map[string]cty.Value) {
	vals["has_logical_redundancy"] = cty.StringVal(p.HasLogicalRedundancy)
}

func EncodeDxConnection_JumboFrameCapable(p DxConnectionObservation, vals map[string]cty.Value) {
	vals["jumbo_frame_capable"] = cty.BoolVal(p.JumboFrameCapable)
}