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

func EncodeEksFargateProfile(r EksFargateProfile) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEksFargateProfile_ClusterName(r.Spec.ForProvider, ctyVal)
	EncodeEksFargateProfile_FargateProfileName(r.Spec.ForProvider, ctyVal)
	EncodeEksFargateProfile_Id(r.Spec.ForProvider, ctyVal)
	EncodeEksFargateProfile_PodExecutionRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeEksFargateProfile_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeEksFargateProfile_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEksFargateProfile_Selector(r.Spec.ForProvider.Selector, ctyVal)
	EncodeEksFargateProfile_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeEksFargateProfile_Arn(r.Status.AtProvider, ctyVal)
	EncodeEksFargateProfile_Status(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEksFargateProfile_ClusterName(p EksFargateProfileParameters, vals map[string]cty.Value) {
	vals["cluster_name"] = cty.StringVal(p.ClusterName)
}

func EncodeEksFargateProfile_FargateProfileName(p EksFargateProfileParameters, vals map[string]cty.Value) {
	vals["fargate_profile_name"] = cty.StringVal(p.FargateProfileName)
}

func EncodeEksFargateProfile_Id(p EksFargateProfileParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEksFargateProfile_PodExecutionRoleArn(p EksFargateProfileParameters, vals map[string]cty.Value) {
	vals["pod_execution_role_arn"] = cty.StringVal(p.PodExecutionRoleArn)
}

func EncodeEksFargateProfile_SubnetIds(p EksFargateProfileParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeEksFargateProfile_Tags(p EksFargateProfileParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEksFargateProfile_Selector(p []Selector, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEksFargateProfile_Selector_Labels(v, ctyVal)
		EncodeEksFargateProfile_Selector_Namespace(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["selector"] = cty.SetVal(valsForCollection)
}

func EncodeEksFargateProfile_Selector_Labels(p Selector, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Labels {
		mVals[key] = cty.StringVal(value)
	}
	vals["labels"] = cty.MapVal(mVals)
}

func EncodeEksFargateProfile_Selector_Namespace(p Selector, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeEksFargateProfile_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeEksFargateProfile_Timeouts_Delete(p, ctyVal)
	EncodeEksFargateProfile_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeEksFargateProfile_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeEksFargateProfile_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeEksFargateProfile_Arn(p EksFargateProfileObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEksFargateProfile_Status(p EksFargateProfileObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}