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
	r, ok := mr.(*CloudfrontOriginAccessIdentity)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CloudfrontOriginAccessIdentity.")
	}
	return EncodeCloudfrontOriginAccessIdentity(*r), nil
}

func EncodeCloudfrontOriginAccessIdentity(r CloudfrontOriginAccessIdentity) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudfrontOriginAccessIdentity_Comment(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_CallerReference(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_CloudfrontAccessIdentityPath(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_Etag(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_IamArn(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_S3CanonicalUserId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudfrontOriginAccessIdentity_Comment(p CloudfrontOriginAccessIdentityParameters, vals map[string]cty.Value) {
	vals["comment"] = cty.StringVal(p.Comment)
}

func EncodeCloudfrontOriginAccessIdentity_Id(p CloudfrontOriginAccessIdentityParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudfrontOriginAccessIdentity_CallerReference(p CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	vals["caller_reference"] = cty.StringVal(p.CallerReference)
}

func EncodeCloudfrontOriginAccessIdentity_CloudfrontAccessIdentityPath(p CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	vals["cloudfront_access_identity_path"] = cty.StringVal(p.CloudfrontAccessIdentityPath)
}

func EncodeCloudfrontOriginAccessIdentity_Etag(p CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	vals["etag"] = cty.StringVal(p.Etag)
}

func EncodeCloudfrontOriginAccessIdentity_IamArn(p CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	vals["iam_arn"] = cty.StringVal(p.IamArn)
}

func EncodeCloudfrontOriginAccessIdentity_S3CanonicalUserId(p CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	vals["s3_canonical_user_id"] = cty.StringVal(p.S3CanonicalUserId)
}