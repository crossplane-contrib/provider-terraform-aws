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

package v1alpha1func EncodeSsmMaintenanceWindow(r SsmMaintenanceWindow) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSsmMaintenanceWindow_ScheduleTimezone(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_Cutoff(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_Description(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_Schedule(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_AllowUnassociatedTargets(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_Duration(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_EndDate(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindow_StartDate(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeSsmMaintenanceWindow_ScheduleTimezone(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["schedule_timezone"] = cty.StringVal(p.ScheduleTimezone)
}

func EncodeSsmMaintenanceWindow_Tags(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSsmMaintenanceWindow_Cutoff(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["cutoff"] = cty.IntVal(p.Cutoff)
}

func EncodeSsmMaintenanceWindow_Description(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSsmMaintenanceWindow_Id(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmMaintenanceWindow_Name(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmMaintenanceWindow_Schedule(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["schedule"] = cty.StringVal(p.Schedule)
}

func EncodeSsmMaintenanceWindow_AllowUnassociatedTargets(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["allow_unassociated_targets"] = cty.BoolVal(p.AllowUnassociatedTargets)
}

func EncodeSsmMaintenanceWindow_Duration(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["duration"] = cty.IntVal(p.Duration)
}

func EncodeSsmMaintenanceWindow_Enabled(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeSsmMaintenanceWindow_EndDate(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["end_date"] = cty.StringVal(p.EndDate)
}

func EncodeSsmMaintenanceWindow_StartDate(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	vals["start_date"] = cty.StringVal(p.StartDate)
}