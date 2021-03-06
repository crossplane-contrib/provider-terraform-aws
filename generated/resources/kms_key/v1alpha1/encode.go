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
	r, ok := mr.(*KmsKey)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a KmsKey.")
	}
	return EncodeKmsKey(*r), nil
}

func EncodeKmsKey(r KmsKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKmsKey_Tags(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_EnableKeyRotation(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_IsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_KeyUsage(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_Policy(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_CustomerMasterKeySpec(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_DeletionWindowInDays(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_Description(r.Spec.ForProvider, ctyVal)
	EncodeKmsKey_Arn(r.Status.AtProvider, ctyVal)
	EncodeKmsKey_KeyId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeKmsKey_Tags(p KmsKeyParameters, vals map[string]cty.Value) {
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

func EncodeKmsKey_EnableKeyRotation(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["enable_key_rotation"] = cty.BoolVal(p.EnableKeyRotation)
}

func EncodeKmsKey_IsEnabled(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["is_enabled"] = cty.BoolVal(p.IsEnabled)
}

func EncodeKmsKey_KeyUsage(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["key_usage"] = cty.StringVal(p.KeyUsage)
}

func EncodeKmsKey_Policy(p KmsKeyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
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

func EncodeKmsKey_Arn(p KmsKeyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeKmsKey_KeyId(p KmsKeyObservation, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}