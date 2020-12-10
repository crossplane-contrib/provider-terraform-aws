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
	r, ok := mr.(*IamSamlProvider)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IamSamlProvider.")
	}
	return EncodeIamSamlProvider(*r), nil
}

func EncodeIamSamlProvider(r IamSamlProvider) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamSamlProvider_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamSamlProvider_SamlMetadataDocument(r.Spec.ForProvider, ctyVal)
	EncodeIamSamlProvider_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamSamlProvider_ValidUntil(r.Status.AtProvider, ctyVal)
	EncodeIamSamlProvider_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeIamSamlProvider_Name(p IamSamlProviderParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamSamlProvider_SamlMetadataDocument(p IamSamlProviderParameters, vals map[string]cty.Value) {
	vals["saml_metadata_document"] = cty.StringVal(p.SamlMetadataDocument)
}

func EncodeIamSamlProvider_Id(p IamSamlProviderParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamSamlProvider_ValidUntil(p IamSamlProviderObservation, vals map[string]cty.Value) {
	vals["valid_until"] = cty.StringVal(p.ValidUntil)
}

func EncodeIamSamlProvider_Arn(p IamSamlProviderObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}