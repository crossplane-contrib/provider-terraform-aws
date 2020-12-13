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
	r, ok := mr.(*DefaultSecurityGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DefaultSecurityGroup.")
	}
	return EncodeDefaultSecurityGroup(*r), nil
}

func EncodeDefaultSecurityGroup(r DefaultSecurityGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDefaultSecurityGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDefaultSecurityGroup_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeDefaultSecurityGroup_Egress(r.Spec.ForProvider.Egress, ctyVal)
	EncodeDefaultSecurityGroup_Ingress(r.Spec.ForProvider.Ingress, ctyVal)
	EncodeDefaultSecurityGroup_RevokeRulesOnDelete(r.Spec.ForProvider, ctyVal)
	EncodeDefaultSecurityGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDefaultSecurityGroup_Description(r.Status.AtProvider, ctyVal)
	EncodeDefaultSecurityGroup_Name(r.Status.AtProvider, ctyVal)
	EncodeDefaultSecurityGroup_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeDefaultSecurityGroup_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDefaultSecurityGroup_Id(p DefaultSecurityGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDefaultSecurityGroup_VpcId(p DefaultSecurityGroupParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeDefaultSecurityGroup_Egress(p []Egress, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeDefaultSecurityGroup_Egress_PrefixListIds(v, ctyVal)
		EncodeDefaultSecurityGroup_Egress_ToPort(v, ctyVal)
		EncodeDefaultSecurityGroup_Egress_Ipv6CidrBlocks(v, ctyVal)
		EncodeDefaultSecurityGroup_Egress_Self(v, ctyVal)
		EncodeDefaultSecurityGroup_Egress_CidrBlocks(v, ctyVal)
		EncodeDefaultSecurityGroup_Egress_Protocol(v, ctyVal)
		EncodeDefaultSecurityGroup_Egress_SecurityGroups(v, ctyVal)
		EncodeDefaultSecurityGroup_Egress_FromPort(v, ctyVal)
		EncodeDefaultSecurityGroup_Egress_Description(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["egress"] = cty.SetVal(valsForCollection)
}

func EncodeDefaultSecurityGroup_Egress_PrefixListIds(p Egress, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PrefixListIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["prefix_list_ids"] = cty.ListVal(colVals)
}

func EncodeDefaultSecurityGroup_Egress_ToPort(p Egress, vals map[string]cty.Value) {
	vals["to_port"] = cty.NumberIntVal(p.ToPort)
}

func EncodeDefaultSecurityGroup_Egress_Ipv6CidrBlocks(p Egress, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Ipv6CidrBlocks {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ipv6_cidr_blocks"] = cty.ListVal(colVals)
}

func EncodeDefaultSecurityGroup_Egress_Self(p Egress, vals map[string]cty.Value) {
	vals["self"] = cty.BoolVal(p.Self)
}

func EncodeDefaultSecurityGroup_Egress_CidrBlocks(p Egress, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CidrBlocks {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cidr_blocks"] = cty.ListVal(colVals)
}

func EncodeDefaultSecurityGroup_Egress_Protocol(p Egress, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeDefaultSecurityGroup_Egress_SecurityGroups(p Egress, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeDefaultSecurityGroup_Egress_FromPort(p Egress, vals map[string]cty.Value) {
	vals["from_port"] = cty.NumberIntVal(p.FromPort)
}

func EncodeDefaultSecurityGroup_Egress_Description(p Egress, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDefaultSecurityGroup_Ingress(p []Ingress, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeDefaultSecurityGroup_Ingress_PrefixListIds(v, ctyVal)
		EncodeDefaultSecurityGroup_Ingress_Description(v, ctyVal)
		EncodeDefaultSecurityGroup_Ingress_Ipv6CidrBlocks(v, ctyVal)
		EncodeDefaultSecurityGroup_Ingress_Protocol(v, ctyVal)
		EncodeDefaultSecurityGroup_Ingress_SecurityGroups(v, ctyVal)
		EncodeDefaultSecurityGroup_Ingress_CidrBlocks(v, ctyVal)
		EncodeDefaultSecurityGroup_Ingress_FromPort(v, ctyVal)
		EncodeDefaultSecurityGroup_Ingress_ToPort(v, ctyVal)
		EncodeDefaultSecurityGroup_Ingress_Self(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ingress"] = cty.SetVal(valsForCollection)
}

func EncodeDefaultSecurityGroup_Ingress_PrefixListIds(p Ingress, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PrefixListIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["prefix_list_ids"] = cty.ListVal(colVals)
}

func EncodeDefaultSecurityGroup_Ingress_Description(p Ingress, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDefaultSecurityGroup_Ingress_Ipv6CidrBlocks(p Ingress, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Ipv6CidrBlocks {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ipv6_cidr_blocks"] = cty.ListVal(colVals)
}

func EncodeDefaultSecurityGroup_Ingress_Protocol(p Ingress, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeDefaultSecurityGroup_Ingress_SecurityGroups(p Ingress, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeDefaultSecurityGroup_Ingress_CidrBlocks(p Ingress, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CidrBlocks {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cidr_blocks"] = cty.ListVal(colVals)
}

func EncodeDefaultSecurityGroup_Ingress_FromPort(p Ingress, vals map[string]cty.Value) {
	vals["from_port"] = cty.NumberIntVal(p.FromPort)
}

func EncodeDefaultSecurityGroup_Ingress_ToPort(p Ingress, vals map[string]cty.Value) {
	vals["to_port"] = cty.NumberIntVal(p.ToPort)
}

func EncodeDefaultSecurityGroup_Ingress_Self(p Ingress, vals map[string]cty.Value) {
	vals["self"] = cty.BoolVal(p.Self)
}

func EncodeDefaultSecurityGroup_RevokeRulesOnDelete(p DefaultSecurityGroupParameters, vals map[string]cty.Value) {
	vals["revoke_rules_on_delete"] = cty.BoolVal(p.RevokeRulesOnDelete)
}

func EncodeDefaultSecurityGroup_Tags(p DefaultSecurityGroupParameters, vals map[string]cty.Value) {
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

func EncodeDefaultSecurityGroup_Description(p DefaultSecurityGroupObservation, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDefaultSecurityGroup_Name(p DefaultSecurityGroupObservation, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDefaultSecurityGroup_OwnerId(p DefaultSecurityGroupObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeDefaultSecurityGroup_Arn(p DefaultSecurityGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}