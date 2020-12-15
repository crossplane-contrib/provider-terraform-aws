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
	r, ok := mr.(*CloudfrontOriginAccessIdentity)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCloudfrontOriginAccessIdentity(r, ctyValue)
}

func DecodeCloudfrontOriginAccessIdentity(prev *CloudfrontOriginAccessIdentity, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCloudfrontOriginAccessIdentity_Comment(&new.Spec.ForProvider, valMap)
	DecodeCloudfrontOriginAccessIdentity_Id(&new.Spec.ForProvider, valMap)
	DecodeCloudfrontOriginAccessIdentity_CloudfrontAccessIdentityPath(&new.Status.AtProvider, valMap)
	DecodeCloudfrontOriginAccessIdentity_Etag(&new.Status.AtProvider, valMap)
	DecodeCloudfrontOriginAccessIdentity_IamArn(&new.Status.AtProvider, valMap)
	DecodeCloudfrontOriginAccessIdentity_S3CanonicalUserId(&new.Status.AtProvider, valMap)
	DecodeCloudfrontOriginAccessIdentity_CallerReference(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeCloudfrontOriginAccessIdentity_Comment(p *CloudfrontOriginAccessIdentityParameters, vals map[string]cty.Value) {
	p.Comment = ctwhy.ValueAsString(vals["comment"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudfrontOriginAccessIdentity_Id(p *CloudfrontOriginAccessIdentityParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudfrontOriginAccessIdentity_CloudfrontAccessIdentityPath(p *CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	p.CloudfrontAccessIdentityPath = ctwhy.ValueAsString(vals["cloudfront_access_identity_path"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudfrontOriginAccessIdentity_Etag(p *CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	p.Etag = ctwhy.ValueAsString(vals["etag"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudfrontOriginAccessIdentity_IamArn(p *CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	p.IamArn = ctwhy.ValueAsString(vals["iam_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudfrontOriginAccessIdentity_S3CanonicalUserId(p *CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	p.S3CanonicalUserId = ctwhy.ValueAsString(vals["s3_canonical_user_id"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudfrontOriginAccessIdentity_CallerReference(p *CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	p.CallerReference = ctwhy.ValueAsString(vals["caller_reference"])
}