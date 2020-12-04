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
	"github.com/zclconf/go-cty/cty"
)

func EncodeCodecommitTrigger(r CodecommitTrigger) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodecommitTrigger_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitTrigger_RepositoryName(r.Spec.ForProvider, ctyVal)
	EncodeCodecommitTrigger_Trigger(r.Spec.ForProvider.Trigger, ctyVal)
	EncodeCodecommitTrigger_ConfigurationId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCodecommitTrigger_Id(p CodecommitTriggerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodecommitTrigger_RepositoryName(p CodecommitTriggerParameters, vals map[string]cty.Value) {
	vals["repository_name"] = cty.StringVal(p.RepositoryName)
}

func EncodeCodecommitTrigger_Trigger(p []Trigger, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeCodecommitTrigger_Trigger_Branches(v, ctyVal)
		EncodeCodecommitTrigger_Trigger_CustomData(v, ctyVal)
		EncodeCodecommitTrigger_Trigger_DestinationArn(v, ctyVal)
		EncodeCodecommitTrigger_Trigger_Events(v, ctyVal)
		EncodeCodecommitTrigger_Trigger_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["trigger"] = cty.SetVal(valsForCollection)
}

func EncodeCodecommitTrigger_Trigger_Branches(p Trigger, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Branches {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["branches"] = cty.ListVal(colVals)
}

func EncodeCodecommitTrigger_Trigger_CustomData(p Trigger, vals map[string]cty.Value) {
	vals["custom_data"] = cty.StringVal(p.CustomData)
}

func EncodeCodecommitTrigger_Trigger_DestinationArn(p Trigger, vals map[string]cty.Value) {
	vals["destination_arn"] = cty.StringVal(p.DestinationArn)
}

func EncodeCodecommitTrigger_Trigger_Events(p Trigger, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Events {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["events"] = cty.ListVal(colVals)
}

func EncodeCodecommitTrigger_Trigger_Name(p Trigger, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodecommitTrigger_ConfigurationId(p CodecommitTriggerObservation, vals map[string]cty.Value) {
	vals["configuration_id"] = cty.StringVal(p.ConfigurationId)
}