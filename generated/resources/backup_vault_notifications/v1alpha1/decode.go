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
	r, ok := mr.(*BackupVaultNotifications)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeBackupVaultNotifications(r, ctyValue)
}

func DecodeBackupVaultNotifications(prev *BackupVaultNotifications, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeBackupVaultNotifications_BackupVaultEvents(&new.Spec.ForProvider, valMap)
	DecodeBackupVaultNotifications_BackupVaultName(&new.Spec.ForProvider, valMap)
	DecodeBackupVaultNotifications_Id(&new.Spec.ForProvider, valMap)
	DecodeBackupVaultNotifications_SnsTopicArn(&new.Spec.ForProvider, valMap)
	DecodeBackupVaultNotifications_BackupVaultArn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeBackupVaultNotifications_BackupVaultEvents(p *BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["backup_vault_events"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.BackupVaultEvents = goVals
}

//primitiveTypeDecodeTemplate
func DecodeBackupVaultNotifications_BackupVaultName(p *BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	p.BackupVaultName = ctwhy.ValueAsString(vals["backup_vault_name"])
}

//primitiveTypeDecodeTemplate
func DecodeBackupVaultNotifications_Id(p *BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeBackupVaultNotifications_SnsTopicArn(p *BackupVaultNotificationsParameters, vals map[string]cty.Value) {
	p.SnsTopicArn = ctwhy.ValueAsString(vals["sns_topic_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeBackupVaultNotifications_BackupVaultArn(p *BackupVaultNotificationsObservation, vals map[string]cty.Value) {
	p.BackupVaultArn = ctwhy.ValueAsString(vals["backup_vault_arn"])
}