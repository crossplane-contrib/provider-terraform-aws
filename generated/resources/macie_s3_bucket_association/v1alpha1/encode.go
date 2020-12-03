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

package v1alpha1func EncodeMacieS3BucketAssociation(r MacieS3BucketAssociation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeMacieS3BucketAssociation_BucketName(r.Spec.ForProvider, ctyVal)
	EncodeMacieS3BucketAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeMacieS3BucketAssociation_MemberAccountId(r.Spec.ForProvider, ctyVal)
	EncodeMacieS3BucketAssociation_Prefix(r.Spec.ForProvider, ctyVal)
	EncodeMacieS3BucketAssociation_ClassificationType(r.Spec.ForProvider.ClassificationType, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeMacieS3BucketAssociation_BucketName(p *MacieS3BucketAssociationParameters, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeMacieS3BucketAssociation_Id(p *MacieS3BucketAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMacieS3BucketAssociation_MemberAccountId(p *MacieS3BucketAssociationParameters, vals map[string]cty.Value) {
	vals["member_account_id"] = cty.StringVal(p.MemberAccountId)
}

func EncodeMacieS3BucketAssociation_Prefix(p *MacieS3BucketAssociationParameters, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeMacieS3BucketAssociation_ClassificationType(p *ClassificationType, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ClassificationType {
		ctyVal = make(map[string]cty.Value)
		EncodeMacieS3BucketAssociation_ClassificationType_Continuous(v, ctyVal)
		EncodeMacieS3BucketAssociation_ClassificationType_OneTime(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["classification_type"] = cty.ListVal(valsForCollection)
}

func EncodeMacieS3BucketAssociation_ClassificationType_Continuous(p *ClassificationType, vals map[string]cty.Value) {
	vals["continuous"] = cty.StringVal(p.Continuous)
}

func EncodeMacieS3BucketAssociation_ClassificationType_OneTime(p *ClassificationType, vals map[string]cty.Value) {
	vals["one_time"] = cty.StringVal(p.OneTime)
}