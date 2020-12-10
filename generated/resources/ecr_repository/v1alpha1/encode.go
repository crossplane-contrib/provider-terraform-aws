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
	r, ok := mr.(*EcrRepository)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EcrRepository.")
	}
	return EncodeEcrRepository(*r), nil
}

func EncodeEcrRepository(r EcrRepository) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEcrRepository_Id(r.Spec.ForProvider, ctyVal)
	EncodeEcrRepository_ImageTagMutability(r.Spec.ForProvider, ctyVal)
	EncodeEcrRepository_Name(r.Spec.ForProvider, ctyVal)
	EncodeEcrRepository_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEcrRepository_EncryptionConfiguration(r.Spec.ForProvider.EncryptionConfiguration, ctyVal)
	EncodeEcrRepository_ImageScanningConfiguration(r.Spec.ForProvider.ImageScanningConfiguration, ctyVal)
	EncodeEcrRepository_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeEcrRepository_Arn(r.Status.AtProvider, ctyVal)
	EncodeEcrRepository_RegistryId(r.Status.AtProvider, ctyVal)
	EncodeEcrRepository_RepositoryUrl(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEcrRepository_Id(p EcrRepositoryParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEcrRepository_ImageTagMutability(p EcrRepositoryParameters, vals map[string]cty.Value) {
	vals["image_tag_mutability"] = cty.StringVal(p.ImageTagMutability)
}

func EncodeEcrRepository_Name(p EcrRepositoryParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEcrRepository_Tags(p EcrRepositoryParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEcrRepository_EncryptionConfiguration(p EncryptionConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcrRepository_EncryptionConfiguration_EncryptionType(p, ctyVal)
	EncodeEcrRepository_EncryptionConfiguration_KmsKey(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeEcrRepository_EncryptionConfiguration_EncryptionType(p EncryptionConfiguration, vals map[string]cty.Value) {
	vals["encryption_type"] = cty.StringVal(p.EncryptionType)
}

func EncodeEcrRepository_EncryptionConfiguration_KmsKey(p EncryptionConfiguration, vals map[string]cty.Value) {
	vals["kms_key"] = cty.StringVal(p.KmsKey)
}

func EncodeEcrRepository_ImageScanningConfiguration(p ImageScanningConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcrRepository_ImageScanningConfiguration_ScanOnPush(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["image_scanning_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeEcrRepository_ImageScanningConfiguration_ScanOnPush(p ImageScanningConfiguration, vals map[string]cty.Value) {
	vals["scan_on_push"] = cty.BoolVal(p.ScanOnPush)
}

func EncodeEcrRepository_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeEcrRepository_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeEcrRepository_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeEcrRepository_Arn(p EcrRepositoryObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEcrRepository_RegistryId(p EcrRepositoryObservation, vals map[string]cty.Value) {
	vals["registry_id"] = cty.StringVal(p.RegistryId)
}

func EncodeEcrRepository_RepositoryUrl(p EcrRepositoryObservation, vals map[string]cty.Value) {
	vals["repository_url"] = cty.StringVal(p.RepositoryUrl)
}