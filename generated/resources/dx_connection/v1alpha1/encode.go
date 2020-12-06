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

func EncodeDxConnection(r DxConnection) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxConnection_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_Location(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_Bandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxConnection_AwsDevice(r.Status.AtProvider, ctyVal)
	EncodeDxConnection_Arn(r.Status.AtProvider, ctyVal)
	EncodeDxConnection_HasLogicalRedundancy(r.Status.AtProvider, ctyVal)
	EncodeDxConnection_JumboFrameCapable(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxConnection_Id(p DxConnectionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxConnection_Location(p DxConnectionParameters, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeDxConnection_Tags(p DxConnectionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDxConnection_Bandwidth(p DxConnectionParameters, vals map[string]cty.Value) {
	vals["bandwidth"] = cty.StringVal(p.Bandwidth)
}

func EncodeDxConnection_Name(p DxConnectionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxConnection_AwsDevice(p DxConnectionObservation, vals map[string]cty.Value) {
	vals["aws_device"] = cty.StringVal(p.AwsDevice)
}

func EncodeDxConnection_Arn(p DxConnectionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDxConnection_HasLogicalRedundancy(p DxConnectionObservation, vals map[string]cty.Value) {
	vals["has_logical_redundancy"] = cty.StringVal(p.HasLogicalRedundancy)
}

func EncodeDxConnection_JumboFrameCapable(p DxConnectionObservation, vals map[string]cty.Value) {
	vals["jumbo_frame_capable"] = cty.BoolVal(p.JumboFrameCapable)
}