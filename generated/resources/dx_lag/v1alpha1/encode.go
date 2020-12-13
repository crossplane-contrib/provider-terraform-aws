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
	r, ok := mr.(*DxLag)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DxLag.")
	}
	return EncodeDxLag(*r), nil
}

func EncodeDxLag(r DxLag) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxLag_Location(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_ConnectionsBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeDxLag_JumboFrameCapable(r.Status.AtProvider, ctyVal)
	EncodeDxLag_HasLogicalRedundancy(r.Status.AtProvider, ctyVal)
	EncodeDxLag_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDxLag_Location(p DxLagParameters, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeDxLag_Id(p DxLagParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxLag_Name(p DxLagParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxLag_Tags(p DxLagParameters, vals map[string]cty.Value) {
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

func EncodeDxLag_ConnectionsBandwidth(p DxLagParameters, vals map[string]cty.Value) {
	vals["connections_bandwidth"] = cty.StringVal(p.ConnectionsBandwidth)
}

func EncodeDxLag_ForceDestroy(p DxLagParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeDxLag_JumboFrameCapable(p DxLagObservation, vals map[string]cty.Value) {
	vals["jumbo_frame_capable"] = cty.BoolVal(p.JumboFrameCapable)
}

func EncodeDxLag_HasLogicalRedundancy(p DxLagObservation, vals map[string]cty.Value) {
	vals["has_logical_redundancy"] = cty.StringVal(p.HasLogicalRedundancy)
}

func EncodeDxLag_Arn(p DxLagObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}