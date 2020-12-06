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

func EncodeDxLag(r DxLag) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxLag_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_ConnectionsBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_Location(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_Arn(r.Status.AtProvider, ctyVal)
	EncodeDxLag_HasLogicalRedundancy(r.Status.AtProvider, ctyVal)
	EncodeDxLag_JumboFrameCapable(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxLag_ForceDestroy(p DxLagParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeDxLag_Id(p DxLagParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxLag_Tags(p DxLagParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDxLag_ConnectionsBandwidth(p DxLagParameters, vals map[string]cty.Value) {
	vals["connections_bandwidth"] = cty.StringVal(p.ConnectionsBandwidth)
}

func EncodeDxLag_Location(p DxLagParameters, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeDxLag_Name(p DxLagParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxLag_Arn(p DxLagObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDxLag_HasLogicalRedundancy(p DxLagObservation, vals map[string]cty.Value) {
	vals["has_logical_redundancy"] = cty.StringVal(p.HasLogicalRedundancy)
}

func EncodeDxLag_JumboFrameCapable(p DxLagObservation, vals map[string]cty.Value) {
	vals["jumbo_frame_capable"] = cty.BoolVal(p.JumboFrameCapable)
}