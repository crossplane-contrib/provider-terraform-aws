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

package v1alpha1func EncodeBackupVaultNotifications(r BackupVaultNotifications) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeBackupVaultNotifications_BackupVaultEvents(r.Spec.ForProvider, ctyVal)
	EncodeBackupVaultNotifications_BackupVaultName(r.Spec.ForProvider, ctyVal)
	EncodeBackupVaultNotifications_Id(r.Spec.ForProvider, ctyVal)
	EncodeBackupVaultNotifications_SnsTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeBackupVaultNotifications_BackupVaultArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeBackupVaultNotifications_BackupVaultEvents(p *BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.BackupVaultEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["backup_vault_events"] = cty.SetVal(colVals)
}

func EncodeBackupVaultNotifications_BackupVaultName(p *BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	vals["backup_vault_name"] = cty.StringVal(p.BackupVaultName)
}

func EncodeBackupVaultNotifications_Id(p *BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeBackupVaultNotifications_SnsTopicArn(p *BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	vals["sns_topic_arn"] = cty.StringVal(p.SnsTopicArn)
}

func EncodeBackupVaultNotifications_BackupVaultArn(p *BackupVaultNotificationsObservation, vals map[string]cty.Value) {
	vals["backup_vault_arn"] = cty.StringVal(p.BackupVaultArn)
}