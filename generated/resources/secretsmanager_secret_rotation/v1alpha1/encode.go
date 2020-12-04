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

func EncodeSecretsmanagerSecretRotation(r SecretsmanagerSecretRotation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSecretsmanagerSecretRotation_SecretId(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretRotation_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretRotation_Id(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretRotation_RotationLambdaArn(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretRotation_RotationRules(r.Spec.ForProvider.RotationRules, ctyVal)
	EncodeSecretsmanagerSecretRotation_RotationEnabled(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSecretsmanagerSecretRotation_SecretId(p SecretsmanagerSecretRotationParameters, vals map[string]cty.Value) {
	vals["secret_id"] = cty.StringVal(p.SecretId)
}

func EncodeSecretsmanagerSecretRotation_Tags(p SecretsmanagerSecretRotationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSecretsmanagerSecretRotation_Id(p SecretsmanagerSecretRotationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSecretsmanagerSecretRotation_RotationLambdaArn(p SecretsmanagerSecretRotationParameters, vals map[string]cty.Value) {
	vals["rotation_lambda_arn"] = cty.StringVal(p.RotationLambdaArn)
}

func EncodeSecretsmanagerSecretRotation_RotationRules(p RotationRules, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSecretsmanagerSecretRotation_RotationRules_AutomaticallyAfterDays(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["rotation_rules"] = cty.ListVal(valsForCollection)
}

func EncodeSecretsmanagerSecretRotation_RotationRules_AutomaticallyAfterDays(p RotationRules, vals map[string]cty.Value) {
	vals["automatically_after_days"] = cty.NumberIntVal(p.AutomaticallyAfterDays)
}

func EncodeSecretsmanagerSecretRotation_RotationEnabled(p SecretsmanagerSecretRotationObservation, vals map[string]cty.Value) {
	vals["rotation_enabled"] = cty.BoolVal(p.RotationEnabled)
}