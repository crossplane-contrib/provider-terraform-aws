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
	r, ok := mr.(*CognitoUserPoolDomain)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCognitoUserPoolDomain(r, ctyValue)
}

func DecodeCognitoUserPoolDomain(prev *CognitoUserPoolDomain, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCognitoUserPoolDomain_CertificateArn(&new.Spec.ForProvider, valMap)
	DecodeCognitoUserPoolDomain_Domain(&new.Spec.ForProvider, valMap)
	DecodeCognitoUserPoolDomain_Id(&new.Spec.ForProvider, valMap)
	DecodeCognitoUserPoolDomain_UserPoolId(&new.Spec.ForProvider, valMap)
	DecodeCognitoUserPoolDomain_CloudfrontDistributionArn(&new.Status.AtProvider, valMap)
	DecodeCognitoUserPoolDomain_S3Bucket(&new.Status.AtProvider, valMap)
	DecodeCognitoUserPoolDomain_Version(&new.Status.AtProvider, valMap)
	DecodeCognitoUserPoolDomain_AwsAccountId(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeCognitoUserPoolDomain_CertificateArn(p *CognitoUserPoolDomainParameters, vals map[string]cty.Value) {
	p.CertificateArn = ctwhy.ValueAsString(vals["certificate_arn"])
}

func DecodeCognitoUserPoolDomain_Domain(p *CognitoUserPoolDomainParameters, vals map[string]cty.Value) {
	p.Domain = ctwhy.ValueAsString(vals["domain"])
}

func DecodeCognitoUserPoolDomain_Id(p *CognitoUserPoolDomainParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeCognitoUserPoolDomain_UserPoolId(p *CognitoUserPoolDomainParameters, vals map[string]cty.Value) {
	p.UserPoolId = ctwhy.ValueAsString(vals["user_pool_id"])
}

func DecodeCognitoUserPoolDomain_CloudfrontDistributionArn(p *CognitoUserPoolDomainObservation, vals map[string]cty.Value) {
	p.CloudfrontDistributionArn = ctwhy.ValueAsString(vals["cloudfront_distribution_arn"])
}

func DecodeCognitoUserPoolDomain_S3Bucket(p *CognitoUserPoolDomainObservation, vals map[string]cty.Value) {
	p.S3Bucket = ctwhy.ValueAsString(vals["s3_bucket"])
}

func DecodeCognitoUserPoolDomain_Version(p *CognitoUserPoolDomainObservation, vals map[string]cty.Value) {
	p.Version = ctwhy.ValueAsString(vals["version"])
}

func DecodeCognitoUserPoolDomain_AwsAccountId(p *CognitoUserPoolDomainObservation, vals map[string]cty.Value) {
	p.AwsAccountId = ctwhy.ValueAsString(vals["aws_account_id"])
}