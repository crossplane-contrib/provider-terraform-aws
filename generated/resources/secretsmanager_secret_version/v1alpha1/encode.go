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
	r, ok := mr.(*SecretsmanagerSecretVersion)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SecretsmanagerSecretVersion.")
	}
	return EncodeSecretsmanagerSecretVersion(*r), nil
}

func EncodeSecretsmanagerSecretVersion(r SecretsmanagerSecretVersion) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSecretsmanagerSecretVersion_Id(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretVersion_SecretBinary(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretVersion_SecretId(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretVersion_SecretString(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretVersion_VersionStages(r.Spec.ForProvider, ctyVal)
	EncodeSecretsmanagerSecretVersion_Arn(r.Status.AtProvider, ctyVal)
	EncodeSecretsmanagerSecretVersion_VersionId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeSecretsmanagerSecretVersion_Id(p SecretsmanagerSecretVersionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSecretsmanagerSecretVersion_SecretBinary(p SecretsmanagerSecretVersionParameters, vals map[string]cty.Value) {
	vals["secret_binary"] = cty.StringVal(p.SecretBinary)
}

func EncodeSecretsmanagerSecretVersion_SecretId(p SecretsmanagerSecretVersionParameters, vals map[string]cty.Value) {
	vals["secret_id"] = cty.StringVal(p.SecretId)
}

func EncodeSecretsmanagerSecretVersion_SecretString(p SecretsmanagerSecretVersionParameters, vals map[string]cty.Value) {
	vals["secret_string"] = cty.StringVal(p.SecretString)
}

func EncodeSecretsmanagerSecretVersion_VersionStages(p SecretsmanagerSecretVersionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VersionStages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["version_stages"] = cty.SetVal(colVals)
}

func EncodeSecretsmanagerSecretVersion_Arn(p SecretsmanagerSecretVersionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeSecretsmanagerSecretVersion_VersionId(p SecretsmanagerSecretVersionObservation, vals map[string]cty.Value) {
	vals["version_id"] = cty.StringVal(p.VersionId)
}