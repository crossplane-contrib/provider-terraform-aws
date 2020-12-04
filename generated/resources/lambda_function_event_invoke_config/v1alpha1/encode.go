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

func EncodeLambdaFunctionEventInvokeConfig(r LambdaFunctionEventInvokeConfig) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunctionEventInvokeConfig_Id(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunctionEventInvokeConfig_MaximumEventAgeInSeconds(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunctionEventInvokeConfig_MaximumRetryAttempts(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunctionEventInvokeConfig_Qualifier(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunctionEventInvokeConfig_FunctionName(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunctionEventInvokeConfig_DestinationConfig(r.Spec.ForProvider.DestinationConfig, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeLambdaFunctionEventInvokeConfig_Id(p LambdaFunctionEventInvokeConfigParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLambdaFunctionEventInvokeConfig_MaximumEventAgeInSeconds(p LambdaFunctionEventInvokeConfigParameters, vals map[string]cty.Value) {
	vals["maximum_event_age_in_seconds"] = cty.NumberIntVal(p.MaximumEventAgeInSeconds)
}

func EncodeLambdaFunctionEventInvokeConfig_MaximumRetryAttempts(p LambdaFunctionEventInvokeConfigParameters, vals map[string]cty.Value) {
	vals["maximum_retry_attempts"] = cty.NumberIntVal(p.MaximumRetryAttempts)
}

func EncodeLambdaFunctionEventInvokeConfig_Qualifier(p LambdaFunctionEventInvokeConfigParameters, vals map[string]cty.Value) {
	vals["qualifier"] = cty.StringVal(p.Qualifier)
}

func EncodeLambdaFunctionEventInvokeConfig_FunctionName(p LambdaFunctionEventInvokeConfigParameters, vals map[string]cty.Value) {
	vals["function_name"] = cty.StringVal(p.FunctionName)
}

func EncodeLambdaFunctionEventInvokeConfig_DestinationConfig(p DestinationConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunctionEventInvokeConfig_DestinationConfig_OnFailure(p.OnFailure, ctyVal)
	EncodeLambdaFunctionEventInvokeConfig_DestinationConfig_OnSuccess(p.OnSuccess, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["destination_config"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaFunctionEventInvokeConfig_DestinationConfig_OnFailure(p OnFailure, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunctionEventInvokeConfig_DestinationConfig_OnFailure_Destination(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["on_failure"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaFunctionEventInvokeConfig_DestinationConfig_OnFailure_Destination(p OnFailure, vals map[string]cty.Value) {
	vals["destination"] = cty.StringVal(p.Destination)
}

func EncodeLambdaFunctionEventInvokeConfig_DestinationConfig_OnSuccess(p OnSuccess, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunctionEventInvokeConfig_DestinationConfig_OnSuccess_Destination(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["on_success"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaFunctionEventInvokeConfig_DestinationConfig_OnSuccess_Destination(p OnSuccess, vals map[string]cty.Value) {
	vals["destination"] = cty.StringVal(p.Destination)
}