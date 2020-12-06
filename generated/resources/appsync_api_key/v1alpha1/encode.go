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

func EncodeAppsyncApiKey(r AppsyncApiKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncApiKey_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncApiKey_Description(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncApiKey_Expires(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncApiKey_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncApiKey_Key(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAppsyncApiKey_ApiId(p AppsyncApiKeyParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeAppsyncApiKey_Description(p AppsyncApiKeyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeAppsyncApiKey_Expires(p AppsyncApiKeyParameters, vals map[string]cty.Value) {
	vals["expires"] = cty.StringVal(p.Expires)
}

func EncodeAppsyncApiKey_Id(p AppsyncApiKeyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppsyncApiKey_Key(p AppsyncApiKeyObservation, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}