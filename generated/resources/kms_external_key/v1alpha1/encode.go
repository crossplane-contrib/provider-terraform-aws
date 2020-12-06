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

func EncodeKmsExternalKey(r KmsExternalKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKmsExternalKey_DeletionWindowInDays(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_Description(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_Policy(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_Id(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_KeyMaterialBase64(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_ValidTo(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_Arn(r.Status.AtProvider, ctyVal)
	EncodeKmsExternalKey_KeyState(r.Status.AtProvider, ctyVal)
	EncodeKmsExternalKey_ExpirationModel(r.Status.AtProvider, ctyVal)
	EncodeKmsExternalKey_KeyUsage(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeKmsExternalKey_DeletionWindowInDays(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["deletion_window_in_days"] = cty.NumberIntVal(p.DeletionWindowInDays)
}

func EncodeKmsExternalKey_Description(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeKmsExternalKey_Enabled(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKmsExternalKey_Policy(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeKmsExternalKey_Id(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKmsExternalKey_KeyMaterialBase64(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["key_material_base64"] = cty.StringVal(p.KeyMaterialBase64)
}

func EncodeKmsExternalKey_Tags(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeKmsExternalKey_ValidTo(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["valid_to"] = cty.StringVal(p.ValidTo)
}

func EncodeKmsExternalKey_Arn(p KmsExternalKeyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeKmsExternalKey_KeyState(p KmsExternalKeyObservation, vals map[string]cty.Value) {
	vals["key_state"] = cty.StringVal(p.KeyState)
}

func EncodeKmsExternalKey_ExpirationModel(p KmsExternalKeyObservation, vals map[string]cty.Value) {
	vals["expiration_model"] = cty.StringVal(p.ExpirationModel)
}

func EncodeKmsExternalKey_KeyUsage(p KmsExternalKeyObservation, vals map[string]cty.Value) {
	vals["key_usage"] = cty.StringVal(p.KeyUsage)
}