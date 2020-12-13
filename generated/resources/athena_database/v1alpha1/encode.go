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
	r, ok := mr.(*AthenaDatabase)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a AthenaDatabase.")
	}
	return EncodeAthenaDatabase(*r), nil
}

func EncodeAthenaDatabase(r AthenaDatabase) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAthenaDatabase_Id(r.Spec.ForProvider, ctyVal)
	EncodeAthenaDatabase_Name(r.Spec.ForProvider, ctyVal)
	EncodeAthenaDatabase_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeAthenaDatabase_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeAthenaDatabase_EncryptionConfiguration(r.Spec.ForProvider.EncryptionConfiguration, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeAthenaDatabase_Id(p AthenaDatabaseParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAthenaDatabase_Name(p AthenaDatabaseParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAthenaDatabase_Bucket(p AthenaDatabaseParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeAthenaDatabase_ForceDestroy(p AthenaDatabaseParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeAthenaDatabase_EncryptionConfiguration(p EncryptionConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAthenaDatabase_EncryptionConfiguration_EncryptionOption(p, ctyVal)
	EncodeAthenaDatabase_EncryptionConfiguration_KmsKey(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAthenaDatabase_EncryptionConfiguration_EncryptionOption(p EncryptionConfiguration, vals map[string]cty.Value) {
	vals["encryption_option"] = cty.StringVal(p.EncryptionOption)
}

func EncodeAthenaDatabase_EncryptionConfiguration_KmsKey(p EncryptionConfiguration, vals map[string]cty.Value) {
	vals["kms_key"] = cty.StringVal(p.KmsKey)
}