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

package v1alpha1func EncodeWorklinkFleet(r WorklinkFleet) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWorklinkFleet_DeviceCaCertificate(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_OptimizeForEndUserLocation(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_AuditStreamArn(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_DisplayName(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_IdentityProvider(r.Spec.ForProvider.IdentityProvider, ctyVal)
	EncodeWorklinkFleet_Network(r.Spec.ForProvider.Network, ctyVal)
	EncodeWorklinkFleet_Arn(r.Status.AtProvider, ctyVal)
	EncodeWorklinkFleet_CompanyCode(r.Status.AtProvider, ctyVal)
	EncodeWorklinkFleet_CreatedTime(r.Status.AtProvider, ctyVal)
	EncodeWorklinkFleet_LastUpdatedTime(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeWorklinkFleet_DeviceCaCertificate(p *WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["device_ca_certificate"] = cty.StringVal(p.DeviceCaCertificate)
}

func EncodeWorklinkFleet_Id(p *WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWorklinkFleet_OptimizeForEndUserLocation(p *WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["optimize_for_end_user_location"] = cty.BoolVal(p.OptimizeForEndUserLocation)
}

func EncodeWorklinkFleet_AuditStreamArn(p *WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["audit_stream_arn"] = cty.StringVal(p.AuditStreamArn)
}

func EncodeWorklinkFleet_DisplayName(p *WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["display_name"] = cty.StringVal(p.DisplayName)
}

func EncodeWorklinkFleet_Name(p *WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWorklinkFleet_IdentityProvider(p *IdentityProvider, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IdentityProvider {
		ctyVal = make(map[string]cty.Value)
		EncodeWorklinkFleet_IdentityProvider_SamlMetadata(v, ctyVal)
		EncodeWorklinkFleet_IdentityProvider_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["identity_provider"] = cty.ListVal(valsForCollection)
}

func EncodeWorklinkFleet_IdentityProvider_SamlMetadata(p *IdentityProvider, vals map[string]cty.Value) {
	vals["saml_metadata"] = cty.StringVal(p.SamlMetadata)
}

func EncodeWorklinkFleet_IdentityProvider_Type(p *IdentityProvider, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWorklinkFleet_Network(p *Network, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Network {
		ctyVal = make(map[string]cty.Value)
		EncodeWorklinkFleet_Network_SecurityGroupIds(v, ctyVal)
		EncodeWorklinkFleet_Network_SubnetIds(v, ctyVal)
		EncodeWorklinkFleet_Network_VpcId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["network"] = cty.ListVal(valsForCollection)
}

func EncodeWorklinkFleet_Network_SecurityGroupIds(p *Network, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeWorklinkFleet_Network_SubnetIds(p *Network, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeWorklinkFleet_Network_VpcId(p *Network, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeWorklinkFleet_Arn(p *WorklinkFleetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWorklinkFleet_CompanyCode(p *WorklinkFleetObservation, vals map[string]cty.Value) {
	vals["company_code"] = cty.StringVal(p.CompanyCode)
}

func EncodeWorklinkFleet_CreatedTime(p *WorklinkFleetObservation, vals map[string]cty.Value) {
	vals["created_time"] = cty.StringVal(p.CreatedTime)
}

func EncodeWorklinkFleet_LastUpdatedTime(p *WorklinkFleetObservation, vals map[string]cty.Value) {
	vals["last_updated_time"] = cty.StringVal(p.LastUpdatedTime)
}