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
	r, ok := mr.(*BackupVault)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a BackupVault.")
	}
	return EncodeBackupVault(*r), nil
}

func EncodeBackupVault(r BackupVault) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeBackupVault_Id(r.Spec.ForProvider, ctyVal)
	EncodeBackupVault_KmsKeyArn(r.Spec.ForProvider, ctyVal)
	EncodeBackupVault_Name(r.Spec.ForProvider, ctyVal)
	EncodeBackupVault_Tags(r.Spec.ForProvider, ctyVal)
	EncodeBackupVault_Arn(r.Status.AtProvider, ctyVal)
	EncodeBackupVault_RecoveryPoints(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeBackupVault_Id(p BackupVaultParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeBackupVault_KmsKeyArn(p BackupVaultParameters, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeBackupVault_Name(p BackupVaultParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeBackupVault_Tags(p BackupVaultParameters, vals map[string]cty.Value) {
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

func EncodeBackupVault_Arn(p BackupVaultObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeBackupVault_RecoveryPoints(p BackupVaultObservation, vals map[string]cty.Value) {
	vals["recovery_points"] = cty.NumberIntVal(p.RecoveryPoints)
}