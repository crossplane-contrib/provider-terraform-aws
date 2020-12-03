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

package v1alpha1func EncodeSsmMaintenanceWindowTarget(r SsmMaintenanceWindowTarget) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSsmMaintenanceWindowTarget_Description(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTarget_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTarget_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTarget_OwnerInformation(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTarget_ResourceType(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTarget_WindowId(r.Spec.ForProvider, ctyVal)
	EncodeSsmMaintenanceWindowTarget_Targets(r.Spec.ForProvider.Targets, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeSsmMaintenanceWindowTarget_Description(p *SsmMaintenanceWindowTargetParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSsmMaintenanceWindowTarget_Id(p *SsmMaintenanceWindowTargetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmMaintenanceWindowTarget_Name(p *SsmMaintenanceWindowTargetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmMaintenanceWindowTarget_OwnerInformation(p *SsmMaintenanceWindowTargetParameters, vals map[string]cty.Value) {
	vals["owner_information"] = cty.StringVal(p.OwnerInformation)
}

func EncodeSsmMaintenanceWindowTarget_ResourceType(p *SsmMaintenanceWindowTargetParameters, vals map[string]cty.Value) {
	vals["resource_type"] = cty.StringVal(p.ResourceType)
}

func EncodeSsmMaintenanceWindowTarget_WindowId(p *SsmMaintenanceWindowTargetParameters, vals map[string]cty.Value) {
	vals["window_id"] = cty.StringVal(p.WindowId)
}

func EncodeSsmMaintenanceWindowTarget_Targets(p *Targets, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Targets {
		ctyVal = make(map[string]cty.Value)
		EncodeSsmMaintenanceWindowTarget_Targets_Key(v, ctyVal)
		EncodeSsmMaintenanceWindowTarget_Targets_Values(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["targets"] = cty.ListVal(valsForCollection)
}

func EncodeSsmMaintenanceWindowTarget_Targets_Key(p *Targets, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeSsmMaintenanceWindowTarget_Targets_Values(p *Targets, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}