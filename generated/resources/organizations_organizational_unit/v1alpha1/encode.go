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

package v1alpha1func EncodeOrganizationsOrganizationalUnit(r OrganizationsOrganizationalUnit) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeOrganizationsOrganizationalUnit_Id(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganizationalUnit_Name(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganizationalUnit_ParentId(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganizationalUnit_Accounts(r.Status.AtProvider.Accounts, ctyVal)
	EncodeOrganizationsOrganizationalUnit_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeOrganizationsOrganizationalUnit_Id(p *OrganizationsOrganizationalUnitParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsOrganizationalUnit_Name(p *OrganizationsOrganizationalUnitParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsOrganizationalUnit_ParentId(p *OrganizationsOrganizationalUnitParameters, vals map[string]cty.Value) {
	vals["parent_id"] = cty.StringVal(p.ParentId)
}

func EncodeOrganizationsOrganizationalUnit_Accounts(p *Accounts, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Accounts {
		ctyVal = make(map[string]cty.Value)
		EncodeOrganizationsOrganizationalUnit_Accounts_Id(v, ctyVal)
		EncodeOrganizationsOrganizationalUnit_Accounts_Name(v, ctyVal)
		EncodeOrganizationsOrganizationalUnit_Accounts_Arn(v, ctyVal)
		EncodeOrganizationsOrganizationalUnit_Accounts_Email(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["accounts"] = cty.ListVal(valsForCollection)
}

func EncodeOrganizationsOrganizationalUnit_Accounts_Id(p *Accounts, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsOrganizationalUnit_Accounts_Name(p *Accounts, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsOrganizationalUnit_Accounts_Arn(p *Accounts, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeOrganizationsOrganizationalUnit_Accounts_Email(p *Accounts, vals map[string]cty.Value) {
	vals["email"] = cty.StringVal(p.Email)
}

func EncodeOrganizationsOrganizationalUnit_Arn(p *OrganizationsOrganizationalUnitObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}