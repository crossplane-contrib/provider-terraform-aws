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
	r, ok := mr.(*OrganizationsOrganizationalUnit)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a OrganizationsOrganizationalUnit.")
	}
	return EncodeOrganizationsOrganizationalUnit(*r), nil
}

func EncodeOrganizationsOrganizationalUnit(r OrganizationsOrganizationalUnit) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOrganizationsOrganizationalUnit_Id(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganizationalUnit_Name(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganizationalUnit_ParentId(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganizationalUnit_Accounts(r.Status.AtProvider.Accounts, ctyVal)
	EncodeOrganizationsOrganizationalUnit_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeOrganizationsOrganizationalUnit_Id(p OrganizationsOrganizationalUnitParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsOrganizationalUnit_Name(p OrganizationsOrganizationalUnitParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsOrganizationalUnit_ParentId(p OrganizationsOrganizationalUnitParameters, vals map[string]cty.Value) {
	vals["parent_id"] = cty.StringVal(p.ParentId)
}

func EncodeOrganizationsOrganizationalUnit_Accounts(p []Accounts, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeOrganizationsOrganizationalUnit_Accounts_Email(v, ctyVal)
		EncodeOrganizationsOrganizationalUnit_Accounts_Id(v, ctyVal)
		EncodeOrganizationsOrganizationalUnit_Accounts_Name(v, ctyVal)
		EncodeOrganizationsOrganizationalUnit_Accounts_Arn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["accounts"] = cty.ListVal(valsForCollection)
}

func EncodeOrganizationsOrganizationalUnit_Accounts_Email(p Accounts, vals map[string]cty.Value) {
	vals["email"] = cty.StringVal(p.Email)
}

func EncodeOrganizationsOrganizationalUnit_Accounts_Id(p Accounts, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsOrganizationalUnit_Accounts_Name(p Accounts, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsOrganizationalUnit_Accounts_Arn(p Accounts, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeOrganizationsOrganizationalUnit_Arn(p OrganizationsOrganizationalUnitObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}