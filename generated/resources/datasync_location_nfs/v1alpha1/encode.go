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
	r, ok := mr.(*DatasyncLocationNfs)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DatasyncLocationNfs.")
	}
	return EncodeDatasyncLocationNfs(*r), nil
}

func EncodeDatasyncLocationNfs(r DatasyncLocationNfs) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncLocationNfs_ServerHostname(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationNfs_Subdirectory(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationNfs_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationNfs_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationNfs_OnPremConfig(r.Spec.ForProvider.OnPremConfig, ctyVal)
	EncodeDatasyncLocationNfs_Uri(r.Status.AtProvider, ctyVal)
	EncodeDatasyncLocationNfs_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDatasyncLocationNfs_ServerHostname(p DatasyncLocationNfsParameters, vals map[string]cty.Value) {
	vals["server_hostname"] = cty.StringVal(p.ServerHostname)
}

func EncodeDatasyncLocationNfs_Subdirectory(p DatasyncLocationNfsParameters, vals map[string]cty.Value) {
	vals["subdirectory"] = cty.StringVal(p.Subdirectory)
}

func EncodeDatasyncLocationNfs_Tags(p DatasyncLocationNfsParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDatasyncLocationNfs_Id(p DatasyncLocationNfsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncLocationNfs_OnPremConfig(p OnPremConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncLocationNfs_OnPremConfig_AgentArns(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["on_prem_config"] = cty.ListVal(valsForCollection)
}

func EncodeDatasyncLocationNfs_OnPremConfig_AgentArns(p OnPremConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AgentArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["agent_arns"] = cty.SetVal(colVals)
}

func EncodeDatasyncLocationNfs_Uri(p DatasyncLocationNfsObservation, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeDatasyncLocationNfs_Arn(p DatasyncLocationNfsObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}