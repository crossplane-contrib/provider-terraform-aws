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

package v1alpha1func EncodeSagemakerNotebookInstanceLifecycleConfiguration(r SagemakerNotebookInstanceLifecycleConfiguration) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSagemakerNotebookInstanceLifecycleConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstanceLifecycleConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstanceLifecycleConfiguration_OnCreate(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstanceLifecycleConfiguration_OnStart(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstanceLifecycleConfiguration_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSagemakerNotebookInstanceLifecycleConfiguration_Id(p *SagemakerNotebookInstanceLifecycleConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSagemakerNotebookInstanceLifecycleConfiguration_Name(p *SagemakerNotebookInstanceLifecycleConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSagemakerNotebookInstanceLifecycleConfiguration_OnCreate(p *SagemakerNotebookInstanceLifecycleConfigurationParameters, vals map[string]cty.Value) {
	vals["on_create"] = cty.StringVal(p.OnCreate)
}

func EncodeSagemakerNotebookInstanceLifecycleConfiguration_OnStart(p *SagemakerNotebookInstanceLifecycleConfigurationParameters, vals map[string]cty.Value) {
	vals["on_start"] = cty.StringVal(p.OnStart)
}

func EncodeSagemakerNotebookInstanceLifecycleConfiguration_Arn(p *SagemakerNotebookInstanceLifecycleConfigurationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}