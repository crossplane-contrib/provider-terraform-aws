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

func EncodeSnsTopicPolicy(r SnsTopicPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSnsTopicPolicy_Arn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicPolicy_Policy(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeSnsTopicPolicy_Arn(p SnsTopicPolicyParameters, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeSnsTopicPolicy_Id(p SnsTopicPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSnsTopicPolicy_Policy(p SnsTopicPolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}