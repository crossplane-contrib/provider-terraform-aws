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

func EncodeKmsKey(r KmsKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKmsKey_IsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_KeyUsage(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_CustomerMasterKeySpec(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_DeletionWindowInDays(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_Description(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_EnableKeyRotation(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_Id(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_Policy(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_Arn(r.Status.AtProvider, ctyVal)
	EncodeKmsKey_KeyId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeKmsKey_IsEnabled(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["is_enabled"] = cty.BoolVal(p.IsEnabled)
}

func EncodeKmsKey_KeyUsage(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["key_usage"] = cty.StringVal(p.KeyUsage)
}

func EncodeKmsKey_Tags(p KmsKeyParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeKmsKey_CustomerMasterKeySpec(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["customer_master_key_spec"] = cty.StringVal(p.CustomerMasterKeySpec)
}

func EncodeKmsKey_DeletionWindowInDays(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["deletion_window_in_days"] = cty.NumberIntVal(p.DeletionWindowInDays)
}

func EncodeKmsKey_Description(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeKmsKey_EnableKeyRotation(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["enable_key_rotation"] = cty.BoolVal(p.EnableKeyRotation)
}

func EncodeKmsKey_Id(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKmsKey_Policy(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeKmsKey_Arn(p KmsKeyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeKmsKey_KeyId(p KmsKeyObservation, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}