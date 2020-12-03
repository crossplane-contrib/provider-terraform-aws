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

package v1alpha1func EncodeNetworkAcl(r NetworkAcl) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeNetworkAcl_Ingress(r.Spec.ForProvider.Ingress, ctyVal)
	EncodeNetworkAcl_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAcl_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAcl_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAcl_Egress(r.Spec.ForProvider.Egress, ctyVal)
	EncodeNetworkAcl_Id(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAcl_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeNetworkAcl_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeNetworkAcl_Ingress(p *Ingress, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Ingress {
		ctyVal = make(map[string]cty.Value)
		EncodeNetworkAcl_Ingress_RuleNo(v, ctyVal)
		EncodeNetworkAcl_Ingress_Action(v, ctyVal)
		EncodeNetworkAcl_Ingress_FromPort(v, ctyVal)
		EncodeNetworkAcl_Ingress_IcmpCode(v, ctyVal)
		EncodeNetworkAcl_Ingress_IcmpType(v, ctyVal)
		EncodeNetworkAcl_Ingress_Protocol(v, ctyVal)
		EncodeNetworkAcl_Ingress_Ipv6CidrBlock(v, ctyVal)
		EncodeNetworkAcl_Ingress_ToPort(v, ctyVal)
		EncodeNetworkAcl_Ingress_CidrBlock(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ingress"] = cty.SetVal(valsForCollection)
}

func EncodeNetworkAcl_Ingress_RuleNo(p *Ingress, vals map[string]cty.Value) {
	vals["rule_no"] = cty.IntVal(p.RuleNo)
}

func EncodeNetworkAcl_Ingress_Action(p *Ingress, vals map[string]cty.Value) {
	vals["action"] = cty.StringVal(p.Action)
}

func EncodeNetworkAcl_Ingress_FromPort(p *Ingress, vals map[string]cty.Value) {
	vals["from_port"] = cty.IntVal(p.FromPort)
}

func EncodeNetworkAcl_Ingress_IcmpCode(p *Ingress, vals map[string]cty.Value) {
	vals["icmp_code"] = cty.IntVal(p.IcmpCode)
}

func EncodeNetworkAcl_Ingress_IcmpType(p *Ingress, vals map[string]cty.Value) {
	vals["icmp_type"] = cty.IntVal(p.IcmpType)
}

func EncodeNetworkAcl_Ingress_Protocol(p *Ingress, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeNetworkAcl_Ingress_Ipv6CidrBlock(p *Ingress, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}

func EncodeNetworkAcl_Ingress_ToPort(p *Ingress, vals map[string]cty.Value) {
	vals["to_port"] = cty.IntVal(p.ToPort)
}

func EncodeNetworkAcl_Ingress_CidrBlock(p *Ingress, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeNetworkAcl_SubnetIds(p *NetworkAclParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeNetworkAcl_Tags(p *NetworkAclParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeNetworkAcl_VpcId(p *NetworkAclParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeNetworkAcl_Egress(p *Egress, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Egress {
		ctyVal = make(map[string]cty.Value)
		EncodeNetworkAcl_Egress_ToPort(v, ctyVal)
		EncodeNetworkAcl_Egress_IcmpType(v, ctyVal)
		EncodeNetworkAcl_Egress_Ipv6CidrBlock(v, ctyVal)
		EncodeNetworkAcl_Egress_CidrBlock(v, ctyVal)
		EncodeNetworkAcl_Egress_FromPort(v, ctyVal)
		EncodeNetworkAcl_Egress_RuleNo(v, ctyVal)
		EncodeNetworkAcl_Egress_Action(v, ctyVal)
		EncodeNetworkAcl_Egress_IcmpCode(v, ctyVal)
		EncodeNetworkAcl_Egress_Protocol(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["egress"] = cty.SetVal(valsForCollection)
}

func EncodeNetworkAcl_Egress_ToPort(p *Egress, vals map[string]cty.Value) {
	vals["to_port"] = cty.IntVal(p.ToPort)
}

func EncodeNetworkAcl_Egress_IcmpType(p *Egress, vals map[string]cty.Value) {
	vals["icmp_type"] = cty.IntVal(p.IcmpType)
}

func EncodeNetworkAcl_Egress_Ipv6CidrBlock(p *Egress, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}

func EncodeNetworkAcl_Egress_CidrBlock(p *Egress, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeNetworkAcl_Egress_FromPort(p *Egress, vals map[string]cty.Value) {
	vals["from_port"] = cty.IntVal(p.FromPort)
}

func EncodeNetworkAcl_Egress_RuleNo(p *Egress, vals map[string]cty.Value) {
	vals["rule_no"] = cty.IntVal(p.RuleNo)
}

func EncodeNetworkAcl_Egress_Action(p *Egress, vals map[string]cty.Value) {
	vals["action"] = cty.StringVal(p.Action)
}

func EncodeNetworkAcl_Egress_IcmpCode(p *Egress, vals map[string]cty.Value) {
	vals["icmp_code"] = cty.IntVal(p.IcmpCode)
}

func EncodeNetworkAcl_Egress_Protocol(p *Egress, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeNetworkAcl_Id(p *NetworkAclParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNetworkAcl_OwnerId(p *NetworkAclObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeNetworkAcl_Arn(p *NetworkAclObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}