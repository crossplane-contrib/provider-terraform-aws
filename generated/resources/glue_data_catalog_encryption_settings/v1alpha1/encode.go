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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*GlueDataCatalogEncryptionSettings)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GlueDataCatalogEncryptionSettings.")
	}
	return EncodeGlueDataCatalogEncryptionSettings(*r), nil
}

func EncodeGlueDataCatalogEncryptionSettings(r GlueDataCatalogEncryptionSettings) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueDataCatalogEncryptionSettings_CatalogId(r.Spec.ForProvider, ctyVal)
	EncodeGlueDataCatalogEncryptionSettings_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings(r.Spec.ForProvider.DataCatalogEncryptionSettings, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeGlueDataCatalogEncryptionSettings_CatalogId(p GlueDataCatalogEncryptionSettingsParameters, vals map[string]cty.Value) {
	vals["catalog_id"] = cty.StringVal(p.CatalogId)
}

func EncodeGlueDataCatalogEncryptionSettings_Id(p GlueDataCatalogEncryptionSettingsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings(p DataCatalogEncryptionSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_ConnectionPasswordEncryption(p.ConnectionPasswordEncryption, ctyVal)
	EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_EncryptionAtRest(p.EncryptionAtRest, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["data_catalog_encryption_settings"] = cty.ListVal(valsForCollection)
}

func EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_ConnectionPasswordEncryption(p ConnectionPasswordEncryption, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_ConnectionPasswordEncryption_AwsKmsKeyId(p, ctyVal)
	EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_ConnectionPasswordEncryption_ReturnConnectionPasswordEncrypted(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["connection_password_encryption"] = cty.ListVal(valsForCollection)
}

func EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_ConnectionPasswordEncryption_AwsKmsKeyId(p ConnectionPasswordEncryption, vals map[string]cty.Value) {
	vals["aws_kms_key_id"] = cty.StringVal(p.AwsKmsKeyId)
}

func EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_ConnectionPasswordEncryption_ReturnConnectionPasswordEncrypted(p ConnectionPasswordEncryption, vals map[string]cty.Value) {
	vals["return_connection_password_encrypted"] = cty.BoolVal(p.ReturnConnectionPasswordEncrypted)
}

func EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_EncryptionAtRest(p EncryptionAtRest, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_EncryptionAtRest_CatalogEncryptionMode(p, ctyVal)
	EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_EncryptionAtRest_SseAwsKmsKeyId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption_at_rest"] = cty.ListVal(valsForCollection)
}

func EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_EncryptionAtRest_CatalogEncryptionMode(p EncryptionAtRest, vals map[string]cty.Value) {
	vals["catalog_encryption_mode"] = cty.StringVal(p.CatalogEncryptionMode)
}

func EncodeGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings_EncryptionAtRest_SseAwsKmsKeyId(p EncryptionAtRest, vals map[string]cty.Value) {
	vals["sse_aws_kms_key_id"] = cty.StringVal(p.SseAwsKmsKeyId)
}