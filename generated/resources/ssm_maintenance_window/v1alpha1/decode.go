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
	r, ok := mr.(*SsmMaintenanceWindow)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSsmMaintenanceWindow(r, ctyValue)
}

func DecodeSsmMaintenanceWindow(prev *SsmMaintenanceWindow, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSsmMaintenanceWindow_AllowUnassociatedTargets(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_Cutoff(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_Description(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_Duration(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_StartDate(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_Enabled(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_EndDate(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_Id(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_Name(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_Schedule(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_ScheduleTimezone(&new.Spec.ForProvider, valMap)
	DecodeSsmMaintenanceWindow_Tags(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeSsmMaintenanceWindow_AllowUnassociatedTargets(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.AllowUnassociatedTargets = ctwhy.ValueAsBool(vals["allow_unassociated_targets"])
}

func DecodeSsmMaintenanceWindow_Cutoff(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.Cutoff = ctwhy.ValueAsInt64(vals["cutoff"])
}

func DecodeSsmMaintenanceWindow_Description(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeSsmMaintenanceWindow_Duration(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.Duration = ctwhy.ValueAsInt64(vals["duration"])
}

func DecodeSsmMaintenanceWindow_StartDate(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.StartDate = ctwhy.ValueAsString(vals["start_date"])
}

func DecodeSsmMaintenanceWindow_Enabled(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

func DecodeSsmMaintenanceWindow_EndDate(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.EndDate = ctwhy.ValueAsString(vals["end_date"])
}

func DecodeSsmMaintenanceWindow_Id(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeSsmMaintenanceWindow_Name(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeSsmMaintenanceWindow_Schedule(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.Schedule = ctwhy.ValueAsString(vals["schedule"])
}

func DecodeSsmMaintenanceWindow_ScheduleTimezone(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	p.ScheduleTimezone = ctwhy.ValueAsString(vals["schedule_timezone"])
}

func DecodeSsmMaintenanceWindow_Tags(p *SsmMaintenanceWindowParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}