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

func EncodeIotThing(r IotThing) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIotThing_Id(r.Spec.ForProvider, ctyVal)
	EncodeIotThing_Name(r.Spec.ForProvider, ctyVal)
	EncodeIotThing_ThingTypeName(r.Spec.ForProvider, ctyVal)
	EncodeIotThing_Attributes(r.Spec.ForProvider, ctyVal)
	EncodeIotThing_Version(r.Status.AtProvider, ctyVal)
	EncodeIotThing_Arn(r.Status.AtProvider, ctyVal)
	EncodeIotThing_DefaultClientId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeIotThing_Id(p IotThingParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIotThing_Name(p IotThingParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIotThing_ThingTypeName(p IotThingParameters, vals map[string]cty.Value) {
	vals["thing_type_name"] = cty.StringVal(p.ThingTypeName)
}

func EncodeIotThing_Attributes(p IotThingParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Attributes {
		mVals[key] = cty.StringVal(value)
	}
	vals["attributes"] = cty.MapVal(mVals)
}

func EncodeIotThing_Version(p IotThingObservation, vals map[string]cty.Value) {
	vals["version"] = cty.NumberIntVal(p.Version)
}

func EncodeIotThing_Arn(p IotThingObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeIotThing_DefaultClientId(p IotThingObservation, vals map[string]cty.Value) {
	vals["default_client_id"] = cty.StringVal(p.DefaultClientId)
}