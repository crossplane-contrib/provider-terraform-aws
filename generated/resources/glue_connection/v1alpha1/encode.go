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

package v1alpha1func EncodeGlueConnection(r GlueConnection) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeGlueConnection_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlueConnection_CatalogId(r.Spec.ForProvider, ctyVal)
	EncodeGlueConnection_ConnectionProperties(r.Spec.ForProvider, ctyVal)
	EncodeGlueConnection_ConnectionType(r.Spec.ForProvider, ctyVal)
	EncodeGlueConnection_Description(r.Spec.ForProvider, ctyVal)
	EncodeGlueConnection_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlueConnection_MatchCriteria(r.Spec.ForProvider, ctyVal)
	EncodeGlueConnection_PhysicalConnectionRequirements(r.Spec.ForProvider.PhysicalConnectionRequirements, ctyVal)
	EncodeGlueConnection_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeGlueConnection_Name(p *GlueConnectionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlueConnection_CatalogId(p *GlueConnectionParameters, vals map[string]cty.Value) {
	vals["catalog_id"] = cty.StringVal(p.CatalogId)
}

func EncodeGlueConnection_ConnectionProperties(p *GlueConnectionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ConnectionProperties {
		mVals[key] = cty.StringVal(value)
	}
	vals["connection_properties"] = cty.MapVal(mVals)
}

func EncodeGlueConnection_ConnectionType(p *GlueConnectionParameters, vals map[string]cty.Value) {
	vals["connection_type"] = cty.StringVal(p.ConnectionType)
}

func EncodeGlueConnection_Description(p *GlueConnectionParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGlueConnection_Id(p *GlueConnectionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlueConnection_MatchCriteria(p *GlueConnectionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.MatchCriteria {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["match_criteria"] = cty.ListVal(colVals)
}

func EncodeGlueConnection_PhysicalConnectionRequirements(p *PhysicalConnectionRequirements, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.PhysicalConnectionRequirements {
		ctyVal = make(map[string]cty.Value)
		EncodeGlueConnection_PhysicalConnectionRequirements_AvailabilityZone(v, ctyVal)
		EncodeGlueConnection_PhysicalConnectionRequirements_SecurityGroupIdList(v, ctyVal)
		EncodeGlueConnection_PhysicalConnectionRequirements_SubnetId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["physical_connection_requirements"] = cty.ListVal(valsForCollection)
}

func EncodeGlueConnection_PhysicalConnectionRequirements_AvailabilityZone(p *PhysicalConnectionRequirements, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeGlueConnection_PhysicalConnectionRequirements_SecurityGroupIdList(p *PhysicalConnectionRequirements, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIdList {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_id_list"] = cty.SetVal(colVals)
}

func EncodeGlueConnection_PhysicalConnectionRequirements_SubnetId(p *PhysicalConnectionRequirements, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeGlueConnection_Arn(p *GlueConnectionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}