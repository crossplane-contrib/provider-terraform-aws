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

func EncodeKmsAlias(r KmsAlias) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKmsAlias_Name(r.Spec.ForProvider, ctyVal)
	EncodeKmsAlias_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeKmsAlias_TargetKeyId(r.Spec.ForProvider, ctyVal)
	EncodeKmsAlias_Id(r.Spec.ForProvider, ctyVal)
	EncodeKmsAlias_TargetKeyArn(r.Status.AtProvider, ctyVal)
	EncodeKmsAlias_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeKmsAlias_Name(p KmsAliasParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeKmsAlias_NamePrefix(p KmsAliasParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeKmsAlias_TargetKeyId(p KmsAliasParameters, vals map[string]cty.Value) {
	vals["target_key_id"] = cty.StringVal(p.TargetKeyId)
}

func EncodeKmsAlias_Id(p KmsAliasParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKmsAlias_TargetKeyArn(p KmsAliasObservation, vals map[string]cty.Value) {
	vals["target_key_arn"] = cty.StringVal(p.TargetKeyArn)
}

func EncodeKmsAlias_Arn(p KmsAliasObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}