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

package v1alpha1func EncodeSagemakerNotebookInstance(r SagemakerNotebookInstance) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSagemakerNotebookInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_Name(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_RootAccess(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_LifecycleConfigName(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_DirectInternetAccess(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerNotebookInstance_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSagemakerNotebookInstance_Id(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSagemakerNotebookInstance_InstanceType(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeSagemakerNotebookInstance_KmsKeyId(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeSagemakerNotebookInstance_Name(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSagemakerNotebookInstance_RoleArn(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeSagemakerNotebookInstance_RootAccess(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	vals["root_access"] = cty.StringVal(p.RootAccess)
}

func EncodeSagemakerNotebookInstance_SecurityGroups(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeSagemakerNotebookInstance_Tags(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSagemakerNotebookInstance_SubnetId(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeSagemakerNotebookInstance_LifecycleConfigName(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	vals["lifecycle_config_name"] = cty.StringVal(p.LifecycleConfigName)
}

func EncodeSagemakerNotebookInstance_DirectInternetAccess(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	vals["direct_internet_access"] = cty.StringVal(p.DirectInternetAccess)
}

func EncodeSagemakerNotebookInstance_Arn(p *SagemakerNotebookInstanceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}