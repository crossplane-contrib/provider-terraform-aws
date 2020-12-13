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
	r, ok := mr.(*DatasyncAgent)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DatasyncAgent.")
	}
	return EncodeDatasyncAgent(*r), nil
}

func EncodeDatasyncAgent(r DatasyncAgent) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncAgent_ActivationKey(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_IpAddress(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_Name(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncAgent_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDatasyncAgent_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDatasyncAgent_ActivationKey(p DatasyncAgentParameters, vals map[string]cty.Value) {
	vals["activation_key"] = cty.StringVal(p.ActivationKey)
}

func EncodeDatasyncAgent_Id(p DatasyncAgentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncAgent_IpAddress(p DatasyncAgentParameters, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}

func EncodeDatasyncAgent_Name(p DatasyncAgentParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDatasyncAgent_Tags(p DatasyncAgentParameters, vals map[string]cty.Value) {
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

func EncodeDatasyncAgent_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncAgent_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDatasyncAgent_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDatasyncAgent_Arn(p DatasyncAgentObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}