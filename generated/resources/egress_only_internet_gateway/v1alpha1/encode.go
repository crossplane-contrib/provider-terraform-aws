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

func EncodeEgressOnlyInternetGateway(r EgressOnlyInternetGateway) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEgressOnlyInternetGateway_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeEgressOnlyInternetGateway_Id(r.Spec.ForProvider, ctyVal)
	EncodeEgressOnlyInternetGateway_Tags(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeEgressOnlyInternetGateway_VpcId(p EgressOnlyInternetGatewayParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeEgressOnlyInternetGateway_Id(p EgressOnlyInternetGatewayParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEgressOnlyInternetGateway_Tags(p EgressOnlyInternetGatewayParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}