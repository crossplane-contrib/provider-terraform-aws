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

func EncodeMediaStoreContainer(r MediaStoreContainer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeMediaStoreContainer_Name(r.Spec.ForProvider, ctyVal)
	EncodeMediaStoreContainer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeMediaStoreContainer_Id(r.Spec.ForProvider, ctyVal)
	EncodeMediaStoreContainer_Arn(r.Status.AtProvider, ctyVal)
	EncodeMediaStoreContainer_Endpoint(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeMediaStoreContainer_Name(p MediaStoreContainerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeMediaStoreContainer_Tags(p MediaStoreContainerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeMediaStoreContainer_Id(p MediaStoreContainerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMediaStoreContainer_Arn(p MediaStoreContainerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeMediaStoreContainer_Endpoint(p MediaStoreContainerObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}