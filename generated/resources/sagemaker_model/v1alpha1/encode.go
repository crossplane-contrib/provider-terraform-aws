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

package v1alpha1func EncodeSagemakerModel(r SagemakerModel) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSagemakerModel_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerModel_EnableNetworkIsolation(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerModel_ExecutionRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerModel_Id(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerModel_Name(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerModel_Container(r.Spec.ForProvider.Container, ctyVal)
	EncodeSagemakerModel_PrimaryContainer(r.Spec.ForProvider.PrimaryContainer, ctyVal)
	EncodeSagemakerModel_VpcConfig(r.Spec.ForProvider.VpcConfig, ctyVal)
	EncodeSagemakerModel_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSagemakerModel_Tags(p *SagemakerModelParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSagemakerModel_EnableNetworkIsolation(p *SagemakerModelParameters, vals map[string]cty.Value) {
	vals["enable_network_isolation"] = cty.BoolVal(p.EnableNetworkIsolation)
}

func EncodeSagemakerModel_ExecutionRoleArn(p *SagemakerModelParameters, vals map[string]cty.Value) {
	vals["execution_role_arn"] = cty.StringVal(p.ExecutionRoleArn)
}

func EncodeSagemakerModel_Id(p *SagemakerModelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSagemakerModel_Name(p *SagemakerModelParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSagemakerModel_Container(p *Container, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Container {
		ctyVal = make(map[string]cty.Value)
		EncodeSagemakerModel_Container_ContainerHostname(v, ctyVal)
		EncodeSagemakerModel_Container_Environment(v, ctyVal)
		EncodeSagemakerModel_Container_Image(v, ctyVal)
		EncodeSagemakerModel_Container_ModelDataUrl(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["container"] = cty.ListVal(valsForCollection)
}

func EncodeSagemakerModel_Container_ContainerHostname(p *Container, vals map[string]cty.Value) {
	vals["container_hostname"] = cty.StringVal(p.ContainerHostname)
}

func EncodeSagemakerModel_Container_Environment(p *Container, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Environment {
		mVals[key] = cty.StringVal(value)
	}
	vals["environment"] = cty.MapVal(mVals)
}

func EncodeSagemakerModel_Container_Image(p *Container, vals map[string]cty.Value) {
	vals["image"] = cty.StringVal(p.Image)
}

func EncodeSagemakerModel_Container_ModelDataUrl(p *Container, vals map[string]cty.Value) {
	vals["model_data_url"] = cty.StringVal(p.ModelDataUrl)
}

func EncodeSagemakerModel_PrimaryContainer(p *PrimaryContainer, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.PrimaryContainer {
		ctyVal = make(map[string]cty.Value)
		EncodeSagemakerModel_PrimaryContainer_ContainerHostname(v, ctyVal)
		EncodeSagemakerModel_PrimaryContainer_Environment(v, ctyVal)
		EncodeSagemakerModel_PrimaryContainer_Image(v, ctyVal)
		EncodeSagemakerModel_PrimaryContainer_ModelDataUrl(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["primary_container"] = cty.ListVal(valsForCollection)
}

func EncodeSagemakerModel_PrimaryContainer_ContainerHostname(p *PrimaryContainer, vals map[string]cty.Value) {
	vals["container_hostname"] = cty.StringVal(p.ContainerHostname)
}

func EncodeSagemakerModel_PrimaryContainer_Environment(p *PrimaryContainer, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Environment {
		mVals[key] = cty.StringVal(value)
	}
	vals["environment"] = cty.MapVal(mVals)
}

func EncodeSagemakerModel_PrimaryContainer_Image(p *PrimaryContainer, vals map[string]cty.Value) {
	vals["image"] = cty.StringVal(p.Image)
}

func EncodeSagemakerModel_PrimaryContainer_ModelDataUrl(p *PrimaryContainer, vals map[string]cty.Value) {
	vals["model_data_url"] = cty.StringVal(p.ModelDataUrl)
}

func EncodeSagemakerModel_VpcConfig(p *VpcConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VpcConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeSagemakerModel_VpcConfig_SecurityGroupIds(v, ctyVal)
		EncodeSagemakerModel_VpcConfig_Subnets(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["vpc_config"] = cty.ListVal(valsForCollection)
}

func EncodeSagemakerModel_VpcConfig_SecurityGroupIds(p *VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeSagemakerModel_VpcConfig_Subnets(p *VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Subnets {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnets"] = cty.SetVal(colVals)
}

func EncodeSagemakerModel_Arn(p *SagemakerModelObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}