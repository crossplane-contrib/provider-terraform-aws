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
	r, ok := mr.(*CognitoIdentityProvider)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCognitoIdentityProvider(r, ctyValue)
}

func DecodeCognitoIdentityProvider(prev *CognitoIdentityProvider, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCognitoIdentityProvider_IdpIdentifiers(&new.Spec.ForProvider, valMap)
	DecodeCognitoIdentityProvider_ProviderDetails(&new.Spec.ForProvider, valMap)
	DecodeCognitoIdentityProvider_ProviderName(&new.Spec.ForProvider, valMap)
	DecodeCognitoIdentityProvider_ProviderType(&new.Spec.ForProvider, valMap)
	DecodeCognitoIdentityProvider_UserPoolId(&new.Spec.ForProvider, valMap)
	DecodeCognitoIdentityProvider_AttributeMapping(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeCognitoIdentityProvider_IdpIdentifiers(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["idp_identifiers"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.IdpIdentifiers = goVals
}

//primitiveMapTypeDecodeTemplate
func DecodeCognitoIdentityProvider_ProviderDetails(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["provider_details"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.ProviderDetails = vMap
}

//primitiveTypeDecodeTemplate
func DecodeCognitoIdentityProvider_ProviderName(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	p.ProviderName = ctwhy.ValueAsString(vals["provider_name"])
}

//primitiveTypeDecodeTemplate
func DecodeCognitoIdentityProvider_ProviderType(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	p.ProviderType = ctwhy.ValueAsString(vals["provider_type"])
}

//primitiveTypeDecodeTemplate
func DecodeCognitoIdentityProvider_UserPoolId(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	p.UserPoolId = ctwhy.ValueAsString(vals["user_pool_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeCognitoIdentityProvider_AttributeMapping(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["attribute_mapping"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.AttributeMapping = vMap
}