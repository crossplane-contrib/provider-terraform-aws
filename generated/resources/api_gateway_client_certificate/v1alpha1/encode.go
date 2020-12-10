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
	r, ok := mr.(*ApiGatewayClientCertificate)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayClientCertificate.")
	}
	return EncodeApiGatewayClientCertificate(*r), nil
}

func EncodeApiGatewayClientCertificate(r ApiGatewayClientCertificate) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayClientCertificate_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayClientCertificate_Description(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayClientCertificate_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayClientCertificate_PemEncodedCertificate(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayClientCertificate_Arn(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayClientCertificate_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayClientCertificate_ExpirationDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayClientCertificate_Tags(p ApiGatewayClientCertificateParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeApiGatewayClientCertificate_Description(p ApiGatewayClientCertificateParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApiGatewayClientCertificate_Id(p ApiGatewayClientCertificateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayClientCertificate_PemEncodedCertificate(p ApiGatewayClientCertificateObservation, vals map[string]cty.Value) {
	vals["pem_encoded_certificate"] = cty.StringVal(p.PemEncodedCertificate)
}

func EncodeApiGatewayClientCertificate_Arn(p ApiGatewayClientCertificateObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeApiGatewayClientCertificate_CreatedDate(p ApiGatewayClientCertificateObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeApiGatewayClientCertificate_ExpirationDate(p ApiGatewayClientCertificateObservation, vals map[string]cty.Value) {
	vals["expiration_date"] = cty.StringVal(p.ExpirationDate)
}