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

package v1alpha1func EncodeBackupPlan(r BackupPlan) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeBackupPlan_Tags(r.Spec.ForProvider, ctyVal)
	EncodeBackupPlan_Id(r.Spec.ForProvider, ctyVal)
	EncodeBackupPlan_Name(r.Spec.ForProvider, ctyVal)
	EncodeBackupPlan_Rule(r.Spec.ForProvider.Rule, ctyVal)
	EncodeBackupPlan_Version(r.Status.AtProvider, ctyVal)
	EncodeBackupPlan_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeBackupPlan_Tags(p *BackupPlanParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeBackupPlan_Id(p *BackupPlanParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeBackupPlan_Name(p *BackupPlanParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeBackupPlan_Rule(p *Rule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Rule {
		ctyVal = make(map[string]cty.Value)
		EncodeBackupPlan_Rule_CompletionWindow(v, ctyVal)
		EncodeBackupPlan_Rule_RecoveryPointTags(v, ctyVal)
		EncodeBackupPlan_Rule_RuleName(v, ctyVal)
		EncodeBackupPlan_Rule_Schedule(v, ctyVal)
		EncodeBackupPlan_Rule_StartWindow(v, ctyVal)
		EncodeBackupPlan_Rule_TargetVaultName(v, ctyVal)
		EncodeBackupPlan_Rule_CopyAction(v.CopyAction, ctyVal)
		EncodeBackupPlan_Rule_Lifecycle(v.Lifecycle, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["rule"] = cty.SetVal(valsForCollection)
}

func EncodeBackupPlan_Rule_CompletionWindow(p *Rule, vals map[string]cty.Value) {
	vals["completion_window"] = cty.IntVal(p.CompletionWindow)
}

func EncodeBackupPlan_Rule_RecoveryPointTags(p *Rule, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.RecoveryPointTags {
		mVals[key] = cty.StringVal(value)
	}
	vals["recovery_point_tags"] = cty.MapVal(mVals)
}

func EncodeBackupPlan_Rule_RuleName(p *Rule, vals map[string]cty.Value) {
	vals["rule_name"] = cty.StringVal(p.RuleName)
}

func EncodeBackupPlan_Rule_Schedule(p *Rule, vals map[string]cty.Value) {
	vals["schedule"] = cty.StringVal(p.Schedule)
}

func EncodeBackupPlan_Rule_StartWindow(p *Rule, vals map[string]cty.Value) {
	vals["start_window"] = cty.IntVal(p.StartWindow)
}

func EncodeBackupPlan_Rule_TargetVaultName(p *Rule, vals map[string]cty.Value) {
	vals["target_vault_name"] = cty.StringVal(p.TargetVaultName)
}

func EncodeBackupPlan_Rule_CopyAction(p *CopyAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CopyAction {
		ctyVal = make(map[string]cty.Value)
		EncodeBackupPlan_Rule_CopyAction_DestinationVaultArn(v, ctyVal)
		EncodeBackupPlan_Rule_CopyAction_Lifecycle(v.Lifecycle, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["copy_action"] = cty.SetVal(valsForCollection)
}

func EncodeBackupPlan_Rule_CopyAction_DestinationVaultArn(p *CopyAction, vals map[string]cty.Value) {
	vals["destination_vault_arn"] = cty.StringVal(p.DestinationVaultArn)
}

func EncodeBackupPlan_Rule_CopyAction_Lifecycle(p *Lifecycle, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Lifecycle {
		ctyVal = make(map[string]cty.Value)
		EncodeBackupPlan_Rule_CopyAction_Lifecycle_ColdStorageAfter(v, ctyVal)
		EncodeBackupPlan_Rule_CopyAction_Lifecycle_DeleteAfter(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lifecycle"] = cty.ListVal(valsForCollection)
}

func EncodeBackupPlan_Rule_CopyAction_Lifecycle_ColdStorageAfter(p *Lifecycle, vals map[string]cty.Value) {
	vals["cold_storage_after"] = cty.IntVal(p.ColdStorageAfter)
}

func EncodeBackupPlan_Rule_CopyAction_Lifecycle_DeleteAfter(p *Lifecycle, vals map[string]cty.Value) {
	vals["delete_after"] = cty.IntVal(p.DeleteAfter)
}

func EncodeBackupPlan_Rule_Lifecycle(p *Lifecycle, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Lifecycle {
		ctyVal = make(map[string]cty.Value)
		EncodeBackupPlan_Rule_Lifecycle_ColdStorageAfter(v, ctyVal)
		EncodeBackupPlan_Rule_Lifecycle_DeleteAfter(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lifecycle"] = cty.ListVal(valsForCollection)
}

func EncodeBackupPlan_Rule_Lifecycle_ColdStorageAfter(p *Lifecycle, vals map[string]cty.Value) {
	vals["cold_storage_after"] = cty.IntVal(p.ColdStorageAfter)
}

func EncodeBackupPlan_Rule_Lifecycle_DeleteAfter(p *Lifecycle, vals map[string]cty.Value) {
	vals["delete_after"] = cty.IntVal(p.DeleteAfter)
}

func EncodeBackupPlan_Version(p *BackupPlanObservation, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeBackupPlan_Arn(p *BackupPlanObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}