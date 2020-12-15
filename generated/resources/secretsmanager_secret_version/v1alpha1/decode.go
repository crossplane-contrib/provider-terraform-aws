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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*SecretsmanagerSecretVersion)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSecretsmanagerSecretVersion(r, ctyValue)
}

func DecodeSecretsmanagerSecretVersion(prev *SecretsmanagerSecretVersion, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSecretsmanagerSecretVersion_SecretBinary(&new.Spec.ForProvider, valMap)
	DecodeSecretsmanagerSecretVersion_SecretId(&new.Spec.ForProvider, valMap)
	DecodeSecretsmanagerSecretVersion_SecretString(&new.Spec.ForProvider, valMap)
	DecodeSecretsmanagerSecretVersion_VersionStages(&new.Spec.ForProvider, valMap)
	DecodeSecretsmanagerSecretVersion_Arn(&new.Status.AtProvider, valMap)
	DecodeSecretsmanagerSecretVersion_VersionId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSecretsmanagerSecretVersion_SecretBinary(p *SecretsmanagerSecretVersionParameters, vals map[string]cty.Value) {
	p.SecretBinary = ctwhy.ValueAsString(vals["secret_binary"])
}

//primitiveTypeDecodeTemplate
func DecodeSecretsmanagerSecretVersion_SecretId(p *SecretsmanagerSecretVersionParameters, vals map[string]cty.Value) {
	p.SecretId = ctwhy.ValueAsString(vals["secret_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSecretsmanagerSecretVersion_SecretString(p *SecretsmanagerSecretVersionParameters, vals map[string]cty.Value) {
	p.SecretString = ctwhy.ValueAsString(vals["secret_string"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeSecretsmanagerSecretVersion_VersionStages(p *SecretsmanagerSecretVersionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["version_stages"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.VersionStages = goVals
}

//primitiveTypeDecodeTemplate
func DecodeSecretsmanagerSecretVersion_Arn(p *SecretsmanagerSecretVersionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSecretsmanagerSecretVersion_VersionId(p *SecretsmanagerSecretVersionObservation, vals map[string]cty.Value) {
	p.VersionId = ctwhy.ValueAsString(vals["version_id"])
}