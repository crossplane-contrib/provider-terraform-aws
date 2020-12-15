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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*BackupVault)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeBackupVault(r, ctyValue)
}

func DecodeBackupVault(prev *BackupVault, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeBackupVault_KmsKeyArn(&new.Spec.ForProvider, valMap)
	DecodeBackupVault_Name(&new.Spec.ForProvider, valMap)
	DecodeBackupVault_Tags(&new.Spec.ForProvider, valMap)
	DecodeBackupVault_Arn(&new.Status.AtProvider, valMap)
	DecodeBackupVault_RecoveryPoints(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeBackupVault_KmsKeyArn(p *BackupVaultParameters, vals map[string]cty.Value) {
	p.KmsKeyArn = ctwhy.ValueAsString(vals["kms_key_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeBackupVault_Name(p *BackupVaultParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeBackupVault_Tags(p *BackupVaultParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeBackupVault_Arn(p *BackupVaultObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeBackupVault_RecoveryPoints(p *BackupVaultObservation, vals map[string]cty.Value) {
	p.RecoveryPoints = ctwhy.ValueAsInt64(vals["recovery_points"])
}