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
	r, ok := mr.(*S3BucketPublicAccessBlock)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a S3BucketPublicAccessBlock.")
	}
	return EncodeS3BucketPublicAccessBlock(*r), nil
}

func EncodeS3BucketPublicAccessBlock(r S3BucketPublicAccessBlock) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketPublicAccessBlock_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketPublicAccessBlock_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketPublicAccessBlock_IgnorePublicAcls(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketPublicAccessBlock_RestrictPublicBuckets(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketPublicAccessBlock_BlockPublicAcls(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketPublicAccessBlock_BlockPublicPolicy(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeS3BucketPublicAccessBlock_Bucket(p S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3BucketPublicAccessBlock_Id(p S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketPublicAccessBlock_IgnorePublicAcls(p S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["ignore_public_acls"] = cty.BoolVal(p.IgnorePublicAcls)
}

func EncodeS3BucketPublicAccessBlock_RestrictPublicBuckets(p S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["restrict_public_buckets"] = cty.BoolVal(p.RestrictPublicBuckets)
}

func EncodeS3BucketPublicAccessBlock_BlockPublicAcls(p S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["block_public_acls"] = cty.BoolVal(p.BlockPublicAcls)
}

func EncodeS3BucketPublicAccessBlock_BlockPublicPolicy(p S3BucketPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["block_public_policy"] = cty.BoolVal(p.BlockPublicPolicy)
}