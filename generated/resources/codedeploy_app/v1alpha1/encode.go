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

func EncodeCodedeployApp(r CodedeployApp) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodedeployApp_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployApp_Name(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployApp_UniqueId(r.Spec.ForProvider, ctyVal)
	EncodeCodedeployApp_ComputePlatform(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeCodedeployApp_Id(p CodedeployAppParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodedeployApp_Name(p CodedeployAppParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodedeployApp_UniqueId(p CodedeployAppParameters, vals map[string]cty.Value) {
	vals["unique_id"] = cty.StringVal(p.UniqueId)
}

func EncodeCodedeployApp_ComputePlatform(p CodedeployAppParameters, vals map[string]cty.Value) {
	vals["compute_platform"] = cty.StringVal(p.ComputePlatform)
}