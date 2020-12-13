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
	r, ok := mr.(*Cloud9EnvironmentEc2)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Cloud9EnvironmentEc2.")
	}
	return EncodeCloud9EnvironmentEc2(*r), nil
}

func EncodeCloud9EnvironmentEc2(r Cloud9EnvironmentEc2) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloud9EnvironmentEc2_OwnerArn(r.Spec.ForProvider, ctyVal)
	EncodeCloud9EnvironmentEc2_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloud9EnvironmentEc2_Description(r.Spec.ForProvider, ctyVal)
	EncodeCloud9EnvironmentEc2_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloud9EnvironmentEc2_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeCloud9EnvironmentEc2_AutomaticStopTimeMinutes(r.Spec.ForProvider, ctyVal)
	EncodeCloud9EnvironmentEc2_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloud9EnvironmentEc2_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeCloud9EnvironmentEc2_Arn(r.Status.AtProvider, ctyVal)
	EncodeCloud9EnvironmentEc2_Type(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCloud9EnvironmentEc2_OwnerArn(p Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	vals["owner_arn"] = cty.StringVal(p.OwnerArn)
}

func EncodeCloud9EnvironmentEc2_Tags(p Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
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

func EncodeCloud9EnvironmentEc2_Description(p Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeCloud9EnvironmentEc2_Name(p Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloud9EnvironmentEc2_SubnetId(p Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeCloud9EnvironmentEc2_AutomaticStopTimeMinutes(p Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	vals["automatic_stop_time_minutes"] = cty.NumberIntVal(p.AutomaticStopTimeMinutes)
}

func EncodeCloud9EnvironmentEc2_Id(p Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloud9EnvironmentEc2_InstanceType(p Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeCloud9EnvironmentEc2_Arn(p Cloud9EnvironmentEc2Observation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCloud9EnvironmentEc2_Type(p Cloud9EnvironmentEc2Observation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}