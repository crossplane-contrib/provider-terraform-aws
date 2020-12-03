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

package v1alpha1func EncodeRedshiftSnapshotSchedule(r RedshiftSnapshotSchedule) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeRedshiftSnapshotSchedule_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Identifier(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_IdentifierPrefix(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Definitions(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Description(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeRedshiftSnapshotSchedule_ForceDestroy(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeRedshiftSnapshotSchedule_Id(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftSnapshotSchedule_Identifier(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["identifier"] = cty.StringVal(p.Identifier)
}

func EncodeRedshiftSnapshotSchedule_IdentifierPrefix(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["identifier_prefix"] = cty.StringVal(p.IdentifierPrefix)
}

func EncodeRedshiftSnapshotSchedule_Tags(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRedshiftSnapshotSchedule_Definitions(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Definitions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["definitions"] = cty.SetVal(colVals)
}

func EncodeRedshiftSnapshotSchedule_Description(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeRedshiftSnapshotSchedule_Arn(p *RedshiftSnapshotScheduleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}