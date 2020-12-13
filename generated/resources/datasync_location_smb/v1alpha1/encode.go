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
	r, ok := mr.(*DatasyncLocationSmb)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DatasyncLocationSmb.")
	}
	return EncodeDatasyncLocationSmb(*r), nil
}

func EncodeDatasyncLocationSmb(r DatasyncLocationSmb) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncLocationSmb_User(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_AgentArns(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_Domain(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_Subdirectory(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_Password(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_ServerHostname(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationSmb_MountOptions(r.Spec.ForProvider.MountOptions, ctyVal)
	EncodeDatasyncLocationSmb_Uri(r.Status.AtProvider, ctyVal)
	EncodeDatasyncLocationSmb_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDatasyncLocationSmb_User(p DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["user"] = cty.StringVal(p.User)
}

func EncodeDatasyncLocationSmb_AgentArns(p DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AgentArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["agent_arns"] = cty.SetVal(colVals)
}

func EncodeDatasyncLocationSmb_Domain(p DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeDatasyncLocationSmb_Id(p DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncLocationSmb_Subdirectory(p DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["subdirectory"] = cty.StringVal(p.Subdirectory)
}

func EncodeDatasyncLocationSmb_Tags(p DatasyncLocationSmbParameters, vals map[string]cty.Value) {
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

func EncodeDatasyncLocationSmb_Password(p DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeDatasyncLocationSmb_ServerHostname(p DatasyncLocationSmbParameters, vals map[string]cty.Value) {
	vals["server_hostname"] = cty.StringVal(p.ServerHostname)
}

func EncodeDatasyncLocationSmb_MountOptions(p MountOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncLocationSmb_MountOptions_Version(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["mount_options"] = cty.ListVal(valsForCollection)
}

func EncodeDatasyncLocationSmb_MountOptions_Version(p MountOptions, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeDatasyncLocationSmb_Uri(p DatasyncLocationSmbObservation, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeDatasyncLocationSmb_Arn(p DatasyncLocationSmbObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}