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

func EncodeGameliftBuild(r GameliftBuild) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftBuild_Name(r.Spec.ForProvider, ctyVal)
	EncodeGameliftBuild_OperatingSystem(r.Spec.ForProvider, ctyVal)
	EncodeGameliftBuild_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGameliftBuild_Version(r.Spec.ForProvider, ctyVal)
	EncodeGameliftBuild_Id(r.Spec.ForProvider, ctyVal)
	EncodeGameliftBuild_StorageLocation(r.Spec.ForProvider.StorageLocation, ctyVal)
	EncodeGameliftBuild_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGameliftBuild_Name(p GameliftBuildParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGameliftBuild_OperatingSystem(p GameliftBuildParameters, vals map[string]cty.Value) {
	vals["operating_system"] = cty.StringVal(p.OperatingSystem)
}

func EncodeGameliftBuild_Tags(p GameliftBuildParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGameliftBuild_Version(p GameliftBuildParameters, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeGameliftBuild_Id(p GameliftBuildParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGameliftBuild_StorageLocation(p StorageLocation, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGameliftBuild_StorageLocation_RoleArn(p, ctyVal)
	EncodeGameliftBuild_StorageLocation_Bucket(p, ctyVal)
	EncodeGameliftBuild_StorageLocation_Key(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["storage_location"] = cty.ListVal(valsForCollection)
}

func EncodeGameliftBuild_StorageLocation_RoleArn(p StorageLocation, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeGameliftBuild_StorageLocation_Bucket(p StorageLocation, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeGameliftBuild_StorageLocation_Key(p StorageLocation, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeGameliftBuild_Arn(p GameliftBuildObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}