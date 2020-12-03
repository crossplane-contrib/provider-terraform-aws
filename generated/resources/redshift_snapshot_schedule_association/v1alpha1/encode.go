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

package v1alpha1func EncodeRedshiftSnapshotScheduleAssociation(r RedshiftSnapshotScheduleAssociation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeRedshiftSnapshotScheduleAssociation_ClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotScheduleAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotScheduleAssociation_ScheduleIdentifier(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeRedshiftSnapshotScheduleAssociation_ClusterIdentifier(p *RedshiftSnapshotScheduleAssociationParameters, vals map[string]cty.Value) {
	vals["cluster_identifier"] = cty.StringVal(p.ClusterIdentifier)
}

func EncodeRedshiftSnapshotScheduleAssociation_Id(p *RedshiftSnapshotScheduleAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftSnapshotScheduleAssociation_ScheduleIdentifier(p *RedshiftSnapshotScheduleAssociationParameters, vals map[string]cty.Value) {
	vals["schedule_identifier"] = cty.StringVal(p.ScheduleIdentifier)
}