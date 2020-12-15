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
	r, ok := mr.(*S3AccountPublicAccessBlock)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeS3AccountPublicAccessBlock(r, ctyValue)
}

func DecodeS3AccountPublicAccessBlock(prev *S3AccountPublicAccessBlock, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeS3AccountPublicAccessBlock_BlockPublicPolicy(&new.Spec.ForProvider, valMap)
	DecodeS3AccountPublicAccessBlock_Id(&new.Spec.ForProvider, valMap)
	DecodeS3AccountPublicAccessBlock_IgnorePublicAcls(&new.Spec.ForProvider, valMap)
	DecodeS3AccountPublicAccessBlock_RestrictPublicBuckets(&new.Spec.ForProvider, valMap)
	DecodeS3AccountPublicAccessBlock_AccountId(&new.Spec.ForProvider, valMap)
	DecodeS3AccountPublicAccessBlock_BlockPublicAcls(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeS3AccountPublicAccessBlock_BlockPublicPolicy(p *S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.BlockPublicPolicy = ctwhy.ValueAsBool(vals["block_public_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeS3AccountPublicAccessBlock_Id(p *S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeS3AccountPublicAccessBlock_IgnorePublicAcls(p *S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.IgnorePublicAcls = ctwhy.ValueAsBool(vals["ignore_public_acls"])
}

//primitiveTypeDecodeTemplate
func DecodeS3AccountPublicAccessBlock_RestrictPublicBuckets(p *S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.RestrictPublicBuckets = ctwhy.ValueAsBool(vals["restrict_public_buckets"])
}

//primitiveTypeDecodeTemplate
func DecodeS3AccountPublicAccessBlock_AccountId(p *S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.AccountId = ctwhy.ValueAsString(vals["account_id"])
}

//primitiveTypeDecodeTemplate
func DecodeS3AccountPublicAccessBlock_BlockPublicAcls(p *S3AccountPublicAccessBlockParameters, vals map[string]cty.Value) {
	p.BlockPublicAcls = ctwhy.ValueAsBool(vals["block_public_acls"])
}