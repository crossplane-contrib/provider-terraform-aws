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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*OrganizationsPolicy)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a OrganizationsPolicy.")
	}
	return EncodeOrganizationsPolicy(*r), nil
}

func EncodeOrganizationsPolicy(r OrganizationsPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOrganizationsPolicy_Content(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Description(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Type(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicy_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeOrganizationsPolicy_Content(p OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeOrganizationsPolicy_Description(p OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeOrganizationsPolicy_Id(p OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsPolicy_Name(p OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsPolicy_Tags(p OrganizationsPolicyParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOrganizationsPolicy_Type(p OrganizationsPolicyParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOrganizationsPolicy_Arn(p OrganizationsPolicyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}