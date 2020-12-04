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

func EncodeIotPolicy(r IotPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIotPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeIotPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeIotPolicy_Policy(r.Spec.ForProvider, ctyVal)
	EncodeIotPolicy_Arn(r.Status.AtProvider, ctyVal)
	EncodeIotPolicy_DefaultVersionId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeIotPolicy_Id(p IotPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIotPolicy_Name(p IotPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIotPolicy_Policy(p IotPolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeIotPolicy_Arn(p IotPolicyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeIotPolicy_DefaultVersionId(p IotPolicyObservation, vals map[string]cty.Value) {
	vals["default_version_id"] = cty.StringVal(p.DefaultVersionId)
}