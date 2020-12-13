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
	r, ok := mr.(*WorklinkFleet)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WorklinkFleet.")
	}
	return EncodeWorklinkFleet(*r), nil
}

func EncodeWorklinkFleet(r WorklinkFleet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWorklinkFleet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_AuditStreamArn(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_DeviceCaCertificate(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_DisplayName(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_OptimizeForEndUserLocation(r.Spec.ForProvider, ctyVal)
	EncodeWorklinkFleet_IdentityProvider(r.Spec.ForProvider.IdentityProvider, ctyVal)
	EncodeWorklinkFleet_Network(r.Spec.ForProvider.Network, ctyVal)
	EncodeWorklinkFleet_CreatedTime(r.Status.AtProvider, ctyVal)
	EncodeWorklinkFleet_LastUpdatedTime(r.Status.AtProvider, ctyVal)
	EncodeWorklinkFleet_Arn(r.Status.AtProvider, ctyVal)
	EncodeWorklinkFleet_CompanyCode(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeWorklinkFleet_Name(p WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWorklinkFleet_AuditStreamArn(p WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["audit_stream_arn"] = cty.StringVal(p.AuditStreamArn)
}

func EncodeWorklinkFleet_DeviceCaCertificate(p WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["device_ca_certificate"] = cty.StringVal(p.DeviceCaCertificate)
}

func EncodeWorklinkFleet_DisplayName(p WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["display_name"] = cty.StringVal(p.DisplayName)
}

func EncodeWorklinkFleet_Id(p WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWorklinkFleet_OptimizeForEndUserLocation(p WorklinkFleetParameters, vals map[string]cty.Value) {
	vals["optimize_for_end_user_location"] = cty.BoolVal(p.OptimizeForEndUserLocation)
}

func EncodeWorklinkFleet_IdentityProvider(p IdentityProvider, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWorklinkFleet_IdentityProvider_SamlMetadata(p, ctyVal)
	EncodeWorklinkFleet_IdentityProvider_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["identity_provider"] = cty.ListVal(valsForCollection)
}

func EncodeWorklinkFleet_IdentityProvider_SamlMetadata(p IdentityProvider, vals map[string]cty.Value) {
	vals["saml_metadata"] = cty.StringVal(p.SamlMetadata)
}

func EncodeWorklinkFleet_IdentityProvider_Type(p IdentityProvider, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWorklinkFleet_Network(p Network, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWorklinkFleet_Network_VpcId(p, ctyVal)
	EncodeWorklinkFleet_Network_SecurityGroupIds(p, ctyVal)
	EncodeWorklinkFleet_Network_SubnetIds(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["network"] = cty.ListVal(valsForCollection)
}

func EncodeWorklinkFleet_Network_VpcId(p Network, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeWorklinkFleet_Network_SecurityGroupIds(p Network, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeWorklinkFleet_Network_SubnetIds(p Network, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeWorklinkFleet_CreatedTime(p WorklinkFleetObservation, vals map[string]cty.Value) {
	vals["created_time"] = cty.StringVal(p.CreatedTime)
}

func EncodeWorklinkFleet_LastUpdatedTime(p WorklinkFleetObservation, vals map[string]cty.Value) {
	vals["last_updated_time"] = cty.StringVal(p.LastUpdatedTime)
}

func EncodeWorklinkFleet_Arn(p WorklinkFleetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWorklinkFleet_CompanyCode(p WorklinkFleetObservation, vals map[string]cty.Value) {
	vals["company_code"] = cty.StringVal(p.CompanyCode)
}