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

func EncodeSwfDomain(r SwfDomain) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSwfDomain_WorkflowExecutionRetentionPeriodInDays(r.Spec.ForProvider, ctyVal)
	EncodeSwfDomain_Description(r.Spec.ForProvider, ctyVal)
	EncodeSwfDomain_Id(r.Spec.ForProvider, ctyVal)
	EncodeSwfDomain_Name(r.Spec.ForProvider, ctyVal)
	EncodeSwfDomain_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeSwfDomain_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSwfDomain_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSwfDomain_WorkflowExecutionRetentionPeriodInDays(p SwfDomainParameters, vals map[string]cty.Value) {
	vals["workflow_execution_retention_period_in_days"] = cty.StringVal(p.WorkflowExecutionRetentionPeriodInDays)
}

func EncodeSwfDomain_Description(p SwfDomainParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSwfDomain_Id(p SwfDomainParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSwfDomain_Name(p SwfDomainParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSwfDomain_NamePrefix(p SwfDomainParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeSwfDomain_Tags(p SwfDomainParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSwfDomain_Arn(p SwfDomainObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}