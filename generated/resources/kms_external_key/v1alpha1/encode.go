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
	r, ok := mr.(*KmsExternalKey)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a KmsExternalKey.")
	}
	return EncodeKmsExternalKey(*r), nil
}

func EncodeKmsExternalKey(r KmsExternalKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKmsExternalKey_Policy(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_ValidTo(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_KeyMaterialBase64(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_DeletionWindowInDays(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_Description(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKmsExternalKey_KeyUsage(r.Status.AtProvider, ctyVal)
	EncodeKmsExternalKey_Arn(r.Status.AtProvider, ctyVal)
	EncodeKmsExternalKey_KeyState(r.Status.AtProvider, ctyVal)
	EncodeKmsExternalKey_ExpirationModel(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeKmsExternalKey_Policy(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeKmsExternalKey_ValidTo(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["valid_to"] = cty.StringVal(p.ValidTo)
}

func EncodeKmsExternalKey_Enabled(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeKmsExternalKey_KeyMaterialBase64(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["key_material_base64"] = cty.StringVal(p.KeyMaterialBase64)
}

func EncodeKmsExternalKey_DeletionWindowInDays(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["deletion_window_in_days"] = cty.NumberIntVal(p.DeletionWindowInDays)
}

func EncodeKmsExternalKey_Description(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeKmsExternalKey_Tags(p KmsExternalKeyParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeKmsExternalKey_KeyUsage(p KmsExternalKeyObservation, vals map[string]cty.Value) {
	vals["key_usage"] = cty.StringVal(p.KeyUsage)
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