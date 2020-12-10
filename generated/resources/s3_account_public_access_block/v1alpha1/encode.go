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
	r, ok := mr.(*S3AccountPublicAccessBlock)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a S3AccountPublicAccessBlock.")
	}
	return EncodeS3AccountPublicAccessBlock(*r), nil
}

func EncodeS3AccountPublicAccessBlock(r S3AccountPublicAccessBlock) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeS3AccountPublicAccessBlock_BlockPublicPolicy(r.Spec.ForProvider, ctyVal)
	EncodeS3AccountPublicAccessBlock_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3AccountPublicAccessBlock_IgnorePublicAcls(r.Spec.ForProvider, ctyVal)
	EncodeS3AccountPublicAccessBlock_RestrictPublicBuckets(r.Spec.ForProvider, ctyVal)
	EncodeS3AccountPublicAccessBlock_AccountId(r.Spec.ForProvider, ctyVal)
	EncodeS3AccountPublicAccessBlock_BlockPublicAcls(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeS3AccountPublicAccessBlock_BlockPublicPolicy(p S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["block_public_policy"] = cty.BoolVal(p.BlockPublicPolicy)
}

func EncodeS3AccountPublicAccessBlock_Id(p S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3AccountPublicAccessBlock_IgnorePublicAcls(p S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["ignore_public_acls"] = cty.BoolVal(p.IgnorePublicAcls)
}

func EncodeS3AccountPublicAccessBlock_RestrictPublicBuckets(p S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["restrict_public_buckets"] = cty.BoolVal(p.RestrictPublicBuckets)
}

func EncodeS3AccountPublicAccessBlock_AccountId(p S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeS3AccountPublicAccessBlock_BlockPublicAcls(p S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	vals["block_public_acls"] = cty.BoolVal(p.BlockPublicAcls)
}