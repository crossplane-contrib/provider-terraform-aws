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
	"github.com/zclconf/go-cty/cty"
)

func EncodeCloudfrontOriginAccessIdentity(r CloudfrontOriginAccessIdentity) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudfrontOriginAccessIdentity_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_Comment(r.Spec.ForProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_IamArn(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_S3CanonicalUserId(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_CallerReference(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_CloudfrontAccessIdentityPath(r.Status.AtProvider, ctyVal)
	EncodeCloudfrontOriginAccessIdentity_Etag(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudfrontOriginAccessIdentity_Id(p CloudfrontOriginAccessIdentityParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudfrontOriginAccessIdentity_Comment(p CloudfrontOriginAccessIdentityParameters, vals map[string]cty.Value) {
	vals["comment"] = cty.StringVal(p.Comment)
}

func EncodeCloudfrontOriginAccessIdentity_IamArn(p CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	vals["iam_arn"] = cty.StringVal(p.IamArn)
}

func EncodeCloudfrontOriginAccessIdentity_S3CanonicalUserId(p CloudfrontOriginAccessIdentityObservation, vals map[string]cty.Value) {
	vals["s3_canonical_user_id"] = cty.StringVal(p.S3CanonicalUserId)
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