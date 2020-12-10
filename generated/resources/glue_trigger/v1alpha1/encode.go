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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*GlueTrigger)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GlueTrigger.")
	}
	return EncodeGlueTrigger(*r), nil
}

func EncodeGlueTrigger(r GlueTrigger) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueTrigger_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueTrigger_Schedule(r.Spec.ForProvider, ctyVal)
	EncodeGlueTrigger_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeGlueTrigger_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlueTrigger_Type(r.Spec.ForProvider, ctyVal)
	EncodeGlueTrigger_WorkflowName(r.Spec.ForProvider, ctyVal)
	EncodeGlueTrigger_Description(r.Spec.ForProvider, ctyVal)
	EncodeGlueTrigger_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueTrigger_Actions(r.Spec.ForProvider.Actions, ctyVal)
	EncodeGlueTrigger_Predicate(r.Spec.ForProvider.Predicate, ctyVal)
	EncodeGlueTrigger_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeGlueTrigger_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlueTrigger_Id(p GlueTriggerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueTrigger_Schedule(p GlueTriggerParameters, vals map[string]cty.Value) {
	vals["schedule"] = cty.StringVal(p.Schedule)
}

func EncodeGlueTrigger_Enabled(p GlueTriggerParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeGlueTrigger_Tags(p GlueTriggerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGlueTrigger_Type(p GlueTriggerParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeGlueTrigger_WorkflowName(p GlueTriggerParameters, vals map[string]cty.Value) {
	vals["workflow_name"] = cty.StringVal(p.WorkflowName)
}

func EncodeGlueTrigger_Description(p GlueTriggerParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGlueTrigger_Name(p GlueTriggerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueTrigger_Actions(p []Actions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGlueTrigger_Actions_Arguments(v, ctyVal)
		EncodeGlueTrigger_Actions_CrawlerName(v, ctyVal)
		EncodeGlueTrigger_Actions_JobName(v, ctyVal)
		EncodeGlueTrigger_Actions_Timeout(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["actions"] = cty.ListVal(valsForCollection)
}

func EncodeGlueTrigger_Actions_Arguments(p Actions, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Arguments {
		mVals[key] = cty.StringVal(value)
	}
	vals["arguments"] = cty.MapVal(mVals)
}

func EncodeGlueTrigger_Actions_CrawlerName(p Actions, vals map[string]cty.Value) {
	vals["crawler_name"] = cty.StringVal(p.CrawlerName)
}

func EncodeGlueTrigger_Actions_JobName(p Actions, vals map[string]cty.Value) {
	vals["job_name"] = cty.StringVal(p.JobName)
}

func EncodeGlueTrigger_Actions_Timeout(p Actions, vals map[string]cty.Value) {
	vals["timeout"] = cty.NumberIntVal(p.Timeout)
}

func EncodeGlueTrigger_Predicate(p Predicate, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlueTrigger_Predicate_Logical(p, ctyVal)
	EncodeGlueTrigger_Predicate_Conditions(p.Conditions, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["predicate"] = cty.ListVal(valsForCollection)
}

func EncodeGlueTrigger_Predicate_Logical(p Predicate, vals map[string]cty.Value) {
	vals["logical"] = cty.StringVal(p.Logical)
}

func EncodeGlueTrigger_Predicate_Conditions(p []Conditions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGlueTrigger_Predicate_Conditions_CrawlerName(v, ctyVal)
		EncodeGlueTrigger_Predicate_Conditions_JobName(v, ctyVal)
		EncodeGlueTrigger_Predicate_Conditions_LogicalOperator(v, ctyVal)
		EncodeGlueTrigger_Predicate_Conditions_State(v, ctyVal)
		EncodeGlueTrigger_Predicate_Conditions_CrawlState(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["conditions"] = cty.ListVal(valsForCollection)
}

func EncodeGlueTrigger_Predicate_Conditions_CrawlerName(p Conditions, vals map[string]cty.Value) {
	vals["crawler_name"] = cty.StringVal(p.CrawlerName)
}

func EncodeGlueTrigger_Predicate_Conditions_JobName(p Conditions, vals map[string]cty.Value) {
	vals["job_name"] = cty.StringVal(p.JobName)
}

func EncodeGlueTrigger_Predicate_Conditions_LogicalOperator(p Conditions, vals map[string]cty.Value) {
	vals["logical_operator"] = cty.StringVal(p.LogicalOperator)
}

func EncodeGlueTrigger_Predicate_Conditions_State(p Conditions, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeGlueTrigger_Predicate_Conditions_CrawlState(p Conditions, vals map[string]cty.Value) {
	vals["crawl_state"] = cty.StringVal(p.CrawlState)
}

func EncodeGlueTrigger_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeGlueTrigger_Timeouts_Create(p, ctyVal)
	EncodeGlueTrigger_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeGlueTrigger_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeGlueTrigger_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeGlueTrigger_Arn(p GlueTriggerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}