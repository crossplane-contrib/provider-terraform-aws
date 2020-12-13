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
	r, ok := mr.(*SecretsmanagerSecretRotation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SecretsmanagerSecretRotation.")
	}
	return EncodeSecretsmanagerSecretRotation(*r), nil
}

func EncodeSecretsmanagerSecretRotation(r SecretsmanagerSecretRotation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSecretsmanagerSecretRotation_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretRotation_Id(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretRotation_RotationLambdaArn(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretRotation_SecretId(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretRotation_RotationRules(r.Spec.ForProvider.RotationRules, ctyVal)
	EncodeSecretsmanagerSecretRotation_RotationEnabled(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeSecretsmanagerSecretRotation_Tags(p SecretsmanagerSecretRotationParameters, vals map[string]cty.Value) {
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

func EncodeSecretsmanagerSecretRotation_Id(p SecretsmanagerSecretRotationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSecretsmanagerSecretRotation_RotationLambdaArn(p SecretsmanagerSecretRotationParameters, vals map[string]cty.Value) {
	vals["rotation_lambda_arn"] = cty.StringVal(p.RotationLambdaArn)
}

func EncodeSecretsmanagerSecretRotation_SecretId(p SecretsmanagerSecretRotationParameters, vals map[string]cty.Value) {
	vals["secret_id"] = cty.StringVal(p.SecretId)
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