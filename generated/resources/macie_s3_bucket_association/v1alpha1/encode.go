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
	r, ok := mr.(*MacieS3BucketAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a MacieS3BucketAssociation.")
	}
	return EncodeMacieS3BucketAssociation(*r), nil
}

func EncodeMacieS3BucketAssociation(r MacieS3BucketAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeMacieS3BucketAssociation_MemberAccountId(r.Spec.ForProvider, ctyVal)
	EncodeMacieS3BucketAssociation_Prefix(r.Spec.ForProvider, ctyVal)
	EncodeMacieS3BucketAssociation_BucketName(r.Spec.ForProvider, ctyVal)
	EncodeMacieS3BucketAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeMacieS3BucketAssociation_ClassificationType(r.Spec.ForProvider.ClassificationType, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeMacieS3BucketAssociation_MemberAccountId(p MacieS3BucketAssociationParameters, vals map[string]cty.Value) {
	vals["member_account_id"] = cty.StringVal(p.MemberAccountId)
}

func EncodeMacieS3BucketAssociation_Prefix(p MacieS3BucketAssociationParameters, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeMacieS3BucketAssociation_BucketName(p MacieS3BucketAssociationParameters, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeMacieS3BucketAssociation_Id(p MacieS3BucketAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMacieS3BucketAssociation_ClassificationType(p ClassificationType, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMacieS3BucketAssociation_ClassificationType_Continuous(p, ctyVal)
	EncodeMacieS3BucketAssociation_ClassificationType_OneTime(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["classification_type"] = cty.ListVal(valsForCollection)
}

func EncodeMacieS3BucketAssociation_ClassificationType_Continuous(p ClassificationType, vals map[string]cty.Value) {
	vals["continuous"] = cty.StringVal(p.Continuous)
}

func EncodeMacieS3BucketAssociation_ClassificationType_OneTime(p ClassificationType, vals map[string]cty.Value) {
	vals["one_time"] = cty.StringVal(p.OneTime)
}