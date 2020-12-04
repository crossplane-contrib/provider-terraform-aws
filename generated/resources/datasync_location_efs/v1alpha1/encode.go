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

func EncodeDatasyncLocationEfs(r DatasyncLocationEfs) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncLocationEfs_EfsFileSystemArn(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationEfs_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationEfs_Subdirectory(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationEfs_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationEfs_Ec2Config(r.Spec.ForProvider.Ec2Config, ctyVal)
	EncodeDatasyncLocationEfs_Arn(r.Status.AtProvider, ctyVal)
	EncodeDatasyncLocationEfs_Uri(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDatasyncLocationEfs_EfsFileSystemArn(p DatasyncLocationEfsParameters, vals map[string]cty.Value) {
	vals["efs_file_system_arn"] = cty.StringVal(p.EfsFileSystemArn)
}

func EncodeDatasyncLocationEfs_Id(p DatasyncLocationEfsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncLocationEfs_Subdirectory(p DatasyncLocationEfsParameters, vals map[string]cty.Value) {
	vals["subdirectory"] = cty.StringVal(p.Subdirectory)
}

func EncodeDatasyncLocationEfs_Tags(p DatasyncLocationEfsParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDatasyncLocationEfs_Ec2Config(p Ec2Config, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncLocationEfs_Ec2Config_SecurityGroupArns(p, ctyVal)
	EncodeDatasyncLocationEfs_Ec2Config_SubnetArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ec2_config"] = cty.ListVal(valsForCollection)
}

func EncodeDatasyncLocationEfs_Ec2Config_SecurityGroupArns(p Ec2Config, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_arns"] = cty.SetVal(colVals)
}

func EncodeDatasyncLocationEfs_Ec2Config_SubnetArn(p Ec2Config, vals map[string]cty.Value) {
	vals["subnet_arn"] = cty.StringVal(p.SubnetArn)
}

func EncodeDatasyncLocationEfs_Arn(p DatasyncLocationEfsObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDatasyncLocationEfs_Uri(p DatasyncLocationEfsObservation, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}