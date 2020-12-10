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
	r, ok := mr.(*ApiGatewayAccount)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayAccount.")
	}
	return EncodeApiGatewayAccount(*r), nil
}

func EncodeApiGatewayAccount(r ApiGatewayAccount) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayAccount_CloudwatchRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAccount_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAccount_ThrottleSettings(r.Status.AtProvider.ThrottleSettings, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayAccount_CloudwatchRoleArn(p ApiGatewayAccountParameters, vals map[string]cty.Value) {
	vals["cloudwatch_role_arn"] = cty.StringVal(p.CloudwatchRoleArn)
}

func EncodeApiGatewayAccount_Id(p ApiGatewayAccountParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayAccount_ThrottleSettings(p []ThrottleSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeApiGatewayAccount_ThrottleSettings_RateLimit(v, ctyVal)
		EncodeApiGatewayAccount_ThrottleSettings_BurstLimit(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["throttle_settings"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayAccount_ThrottleSettings_RateLimit(p ThrottleSettings, vals map[string]cty.Value) {
	vals["rate_limit"] = cty.NumberIntVal(p.RateLimit)
}

func EncodeApiGatewayAccount_ThrottleSettings_BurstLimit(p ThrottleSettings, vals map[string]cty.Value) {
	vals["burst_limit"] = cty.NumberIntVal(p.BurstLimit)
}