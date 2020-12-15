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
	r, ok := mr.(*BackupVaultNotifications)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a BackupVaultNotifications.")
	}
	return EncodeBackupVaultNotifications(*r), nil
}

func EncodeBackupVaultNotifications(r BackupVaultNotifications) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeBackupVaultNotifications_BackupVaultEvents(r.Spec.ForProvider, ctyVal)
	EncodeBackupVaultNotifications_BackupVaultName(r.Spec.ForProvider, ctyVal)
	EncodeBackupVaultNotifications_Id(r.Spec.ForProvider, ctyVal)
	EncodeBackupVaultNotifications_SnsTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeBackupVaultNotifications_BackupVaultArn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeBackupVaultNotifications_BackupVaultEvents(p BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.BackupVaultEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["backup_vault_events"] = cty.SetVal(colVals)
}

func EncodeBackupVaultNotifications_BackupVaultName(p BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	vals["backup_vault_name"] = cty.StringVal(p.BackupVaultName)
}

func EncodeBackupVaultNotifications_Id(p BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeBackupVaultNotifications_SnsTopicArn(p BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	vals["sns_topic_arn"] = cty.StringVal(p.SnsTopicArn)
}

func EncodeBackupVaultNotifications_BackupVaultArn(p BackupVaultNotificationsObservation, vals map[string]cty.Value) {
	vals["backup_vault_arn"] = cty.StringVal(p.BackupVaultArn)
}