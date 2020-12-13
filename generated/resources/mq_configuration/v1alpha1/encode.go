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
	r, ok := mr.(*MqConfiguration)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a MqConfiguration.")
	}
	return EncodeMqConfiguration(*r), nil
}

func EncodeMqConfiguration(r MqConfiguration) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeMqConfiguration_EngineType(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Data(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Description(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Tags(r.Spec.ForProvider, ctyVal)
	EncodeMqConfiguration_Arn(r.Status.AtProvider, ctyVal)
	EncodeMqConfiguration_LatestRevision(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeMqConfiguration_EngineType(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["engine_type"] = cty.StringVal(p.EngineType)
}

func EncodeMqConfiguration_EngineVersion(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeMqConfiguration_Id(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMqConfiguration_Name(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeMqConfiguration_Data(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeMqConfiguration_Description(p MqConfigurationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeMqConfiguration_Tags(p MqConfigurationParameters, vals map[string]cty.Value) {
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

func EncodeMqConfiguration_Arn(p MqConfigurationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeMqConfiguration_LatestRevision(p MqConfigurationObservation, vals map[string]cty.Value) {
	vals["latest_revision"] = cty.NumberIntVal(p.LatestRevision)
}