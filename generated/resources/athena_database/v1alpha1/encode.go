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

func EncodeAthenaDatabase(r AthenaDatabase) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAthenaDatabase_Id(r.Spec.ForProvider, ctyVal)
	EncodeAthenaDatabase_Name(r.Spec.ForProvider, ctyVal)
	EncodeAthenaDatabase_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeAthenaDatabase_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeAthenaDatabase_EncryptionConfiguration(r.Spec.ForProvider.EncryptionConfiguration, ctyVal)

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