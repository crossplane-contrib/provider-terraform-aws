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
	r, ok := mr.(*CognitoUserPoolDomain)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CognitoUserPoolDomain.")
	}
	return EncodeCognitoUserPoolDomain(*r), nil
}

func EncodeCognitoUserPoolDomain(r CognitoUserPoolDomain) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCognitoUserPoolDomain_UserPoolId(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolDomain_CertificateArn(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolDomain_Domain(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolDomain_S3Bucket(r.Status.AtProvider, ctyVal)
	EncodeCognitoUserPoolDomain_Version(r.Status.AtProvider, ctyVal)
	EncodeCognitoUserPoolDomain_AwsAccountId(r.Status.AtProvider, ctyVal)
	EncodeCognitoUserPoolDomain_CloudfrontDistributionArn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeCognitoUserPoolDomain_UserPoolId(p CognitoUserPoolDomainParameters, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}

func EncodeCognitoUserPoolDomain_CertificateArn(p CognitoUserPoolDomainParameters, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}

func EncodeCognitoUserPoolDomain_Domain(p CognitoUserPoolDomainParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeCognitoUserPoolDomain_S3Bucket(p CognitoUserPoolDomainObservation, vals map[string]cty.Value) {
	vals["s3_bucket"] = cty.StringVal(p.S3Bucket)
}

func EncodeCognitoUserPoolDomain_Version(p CognitoUserPoolDomainObservation, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeCognitoUserPoolDomain_AwsAccountId(p CognitoUserPoolDomainObservation, vals map[string]cty.Value) {
	vals["aws_account_id"] = cty.StringVal(p.AwsAccountId)
}

func EncodeCognitoUserPoolDomain_CloudfrontDistributionArn(p CognitoUserPoolDomainObservation, vals map[string]cty.Value) {
	vals["cloudfront_distribution_arn"] = cty.StringVal(p.CloudfrontDistributionArn)
}