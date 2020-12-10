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
	r, ok := mr.(*SecretsmanagerSecret)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SecretsmanagerSecret.")
	}
	return EncodeSecretsmanagerSecret(*r), nil
}

func EncodeSecretsmanagerSecret(r SecretsmanagerSecret) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSecretsmanagerSecret_Description(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecret_Id(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecret_Name(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecret_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecret_RotationLambdaArn(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecret_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecret_Policy(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecret_RecoveryWindowInDays(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecret_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecret_RotationRules(r.Spec.ForProvider.RotationRules, ctyVal)
	EncodeSecretsmanagerSecret_RotationEnabled(r.Status.AtProvider, ctyVal)
	EncodeSecretsmanagerSecret_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSecretsmanagerSecret_Description(p SecretsmanagerSecretParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSecretsmanagerSecret_Id(p SecretsmanagerSecretParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSecretsmanagerSecret_Name(p SecretsmanagerSecretParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSecretsmanagerSecret_NamePrefix(p SecretsmanagerSecretParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeSecretsmanagerSecret_RotationLambdaArn(p SecretsmanagerSecretParameters, vals map[string]cty.Value) {
	vals["rotation_lambda_arn"] = cty.StringVal(p.RotationLambdaArn)
}

func EncodeSecretsmanagerSecret_KmsKeyId(p SecretsmanagerSecretParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeSecretsmanagerSecret_Policy(p SecretsmanagerSecretParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeSecretsmanagerSecret_RecoveryWindowInDays(p SecretsmanagerSecretParameters, vals map[string]cty.Value) {
	vals["recovery_window_in_days"] = cty.NumberIntVal(p.RecoveryWindowInDays)
}

func EncodeSecretsmanagerSecret_Tags(p SecretsmanagerSecretParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSecretsmanagerSecret_RotationRules(p RotationRules, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSecretsmanagerSecret_RotationRules_AutomaticallyAfterDays(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["rotation_rules"] = cty.ListVal(valsForCollection)
}

func EncodeSecretsmanagerSecret_RotationRules_AutomaticallyAfterDays(p RotationRules, vals map[string]cty.Value) {
	vals["automatically_after_days"] = cty.NumberIntVal(p.AutomaticallyAfterDays)
}

func EncodeSecretsmanagerSecret_RotationEnabled(p SecretsmanagerSecretObservation, vals map[string]cty.Value) {
	vals["rotation_enabled"] = cty.BoolVal(p.RotationEnabled)
}

func EncodeSecretsmanagerSecret_Arn(p SecretsmanagerSecretObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}