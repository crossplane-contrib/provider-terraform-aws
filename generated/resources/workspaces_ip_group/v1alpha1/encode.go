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
	"fmt"
	
	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*WorkspacesIpGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WorkspacesIpGroup.")
	}
	return EncodeWorkspacesIpGroup(*r), nil
}

func EncodeWorkspacesIpGroup(r WorkspacesIpGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWorkspacesIpGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesIpGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesIpGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesIpGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWorkspacesIpGroup_Rules(r.Spec.ForProvider.Rules, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeWorkspacesIpGroup_Description(p WorkspacesIpGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeWorkspacesIpGroup_Id(p WorkspacesIpGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWorkspacesIpGroup_Name(p WorkspacesIpGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWorkspacesIpGroup_Tags(p WorkspacesIpGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWorkspacesIpGroup_Rules(p Rules, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWorkspacesIpGroup_Rules_Description(p, ctyVal)
	EncodeWorkspacesIpGroup_Rules_Source(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["rules"] = cty.SetVal(valsForCollection)
}

func EncodeWorkspacesIpGroup_Rules_Description(p Rules, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeWorkspacesIpGroup_Rules_Source(p Rules, vals map[string]cty.Value) {
	vals["source"] = cty.StringVal(p.Source)
}