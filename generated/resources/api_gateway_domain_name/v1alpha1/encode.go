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
	r, ok := mr.(*ApiGatewayDomainName)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayDomainName.")
	}
	return EncodeApiGatewayDomainName(*r), nil
}

func EncodeApiGatewayDomainName(r ApiGatewayDomainName) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayDomainName_CertificateArn(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_RegionalCertificateArn(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_RegionalCertificateName(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_CertificatePrivateKey(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_SecurityPolicy(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_CertificateChain(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_CertificateName(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_CertificateBody(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayDomainName_EndpointConfiguration(r.Spec.ForProvider.EndpointConfiguration, ctyVal)
	EncodeApiGatewayDomainName_RegionalZoneId(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayDomainName_Arn(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayDomainName_RegionalDomainName(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayDomainName_CloudfrontDomainName(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayDomainName_CloudfrontZoneId(r.Status.AtProvider, ctyVal)
	EncodeApiGatewayDomainName_CertificateUploadDate(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayDomainName_CertificateArn(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}

func EncodeApiGatewayDomainName_RegionalCertificateArn(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["regional_certificate_arn"] = cty.StringVal(p.RegionalCertificateArn)
}

func EncodeApiGatewayDomainName_Id(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayDomainName_RegionalCertificateName(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["regional_certificate_name"] = cty.StringVal(p.RegionalCertificateName)
}

func EncodeApiGatewayDomainName_Tags(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
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

func EncodeApiGatewayDomainName_CertificatePrivateKey(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["certificate_private_key"] = cty.StringVal(p.CertificatePrivateKey)
}

func EncodeApiGatewayDomainName_DomainName(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeApiGatewayDomainName_SecurityPolicy(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["security_policy"] = cty.StringVal(p.SecurityPolicy)
}

func EncodeApiGatewayDomainName_CertificateChain(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["certificate_chain"] = cty.StringVal(p.CertificateChain)
}

func EncodeApiGatewayDomainName_CertificateName(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["certificate_name"] = cty.StringVal(p.CertificateName)
}

func EncodeApiGatewayDomainName_CertificateBody(p ApiGatewayDomainNameParameters, vals map[string]cty.Value) {
	vals["certificate_body"] = cty.StringVal(p.CertificateBody)
}

func EncodeApiGatewayDomainName_EndpointConfiguration(p EndpointConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayDomainName_EndpointConfiguration_Types(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["endpoint_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayDomainName_EndpointConfiguration_Types(p EndpointConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Types {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["types"] = cty.ListVal(colVals)
}

func EncodeApiGatewayDomainName_RegionalZoneId(p ApiGatewayDomainNameObservation, vals map[string]cty.Value) {
	vals["regional_zone_id"] = cty.StringVal(p.RegionalZoneId)
}

func EncodeApiGatewayDomainName_Arn(p ApiGatewayDomainNameObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeApiGatewayDomainName_RegionalDomainName(p ApiGatewayDomainNameObservation, vals map[string]cty.Value) {
	vals["regional_domain_name"] = cty.StringVal(p.RegionalDomainName)
}

func EncodeApiGatewayDomainName_CloudfrontDomainName(p ApiGatewayDomainNameObservation, vals map[string]cty.Value) {
	vals["cloudfront_domain_name"] = cty.StringVal(p.CloudfrontDomainName)
}

func EncodeApiGatewayDomainName_CloudfrontZoneId(p ApiGatewayDomainNameObservation, vals map[string]cty.Value) {
	vals["cloudfront_zone_id"] = cty.StringVal(p.CloudfrontZoneId)
}

func EncodeApiGatewayDomainName_CertificateUploadDate(p ApiGatewayDomainNameObservation, vals map[string]cty.Value) {
	vals["certificate_upload_date"] = cty.StringVal(p.CertificateUploadDate)
}