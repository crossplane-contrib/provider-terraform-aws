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

package v1alpha1func EncodeEc2TrafficMirrorFilterRule(r Ec2TrafficMirrorFilterRule) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEc2TrafficMirrorFilterRule_SourceCidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_TrafficMirrorFilterId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_Protocol(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_RuleAction(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_RuleNumber(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_TrafficDirection(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_DestinationCidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_SourcePortRange(r.Spec.ForProvider.SourcePortRange, ctyVal)
	EncodeEc2TrafficMirrorFilterRule_DestinationPortRange(r.Spec.ForProvider.DestinationPortRange, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeEc2TrafficMirrorFilterRule_SourceCidrBlock(p *Ec2TrafficMirrorFilterRuleParameters, vals map[string]cty.Value) {
	vals["source_cidr_block"] = cty.StringVal(p.SourceCidrBlock)
}

func EncodeEc2TrafficMirrorFilterRule_TrafficMirrorFilterId(p *Ec2TrafficMirrorFilterRuleParameters, vals map[string]cty.Value) {
	vals["traffic_mirror_filter_id"] = cty.StringVal(p.TrafficMirrorFilterId)
}

func EncodeEc2TrafficMirrorFilterRule_Protocol(p *Ec2TrafficMirrorFilterRuleParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.IntVal(p.Protocol)
}

func EncodeEc2TrafficMirrorFilterRule_RuleAction(p *Ec2TrafficMirrorFilterRuleParameters, vals map[string]cty.Value) {
	vals["rule_action"] = cty.StringVal(p.RuleAction)
}

func EncodeEc2TrafficMirrorFilterRule_RuleNumber(p *Ec2TrafficMirrorFilterRuleParameters, vals map[string]cty.Value) {
	vals["rule_number"] = cty.IntVal(p.RuleNumber)
}

func EncodeEc2TrafficMirrorFilterRule_TrafficDirection(p *Ec2TrafficMirrorFilterRuleParameters, vals map[string]cty.Value) {
	vals["traffic_direction"] = cty.StringVal(p.TrafficDirection)
}

func EncodeEc2TrafficMirrorFilterRule_Description(p *Ec2TrafficMirrorFilterRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2TrafficMirrorFilterRule_DestinationCidrBlock(p *Ec2TrafficMirrorFilterRuleParameters, vals map[string]cty.Value) {
	vals["destination_cidr_block"] = cty.StringVal(p.DestinationCidrBlock)
}

func EncodeEc2TrafficMirrorFilterRule_Id(p *Ec2TrafficMirrorFilterRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TrafficMirrorFilterRule_SourcePortRange(p *SourcePortRange, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SourcePortRange {
		ctyVal = make(map[string]cty.Value)
		EncodeEc2TrafficMirrorFilterRule_SourcePortRange_FromPort(v, ctyVal)
		EncodeEc2TrafficMirrorFilterRule_SourcePortRange_ToPort(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["source_port_range"] = cty.ListVal(valsForCollection)
}

func EncodeEc2TrafficMirrorFilterRule_SourcePortRange_FromPort(p *SourcePortRange, vals map[string]cty.Value) {
	vals["from_port"] = cty.IntVal(p.FromPort)
}

func EncodeEc2TrafficMirrorFilterRule_SourcePortRange_ToPort(p *SourcePortRange, vals map[string]cty.Value) {
	vals["to_port"] = cty.IntVal(p.ToPort)
}

func EncodeEc2TrafficMirrorFilterRule_DestinationPortRange(p *DestinationPortRange, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DestinationPortRange {
		ctyVal = make(map[string]cty.Value)
		EncodeEc2TrafficMirrorFilterRule_DestinationPortRange_FromPort(v, ctyVal)
		EncodeEc2TrafficMirrorFilterRule_DestinationPortRange_ToPort(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["destination_port_range"] = cty.ListVal(valsForCollection)
}

func EncodeEc2TrafficMirrorFilterRule_DestinationPortRange_FromPort(p *DestinationPortRange, vals map[string]cty.Value) {
	vals["from_port"] = cty.IntVal(p.FromPort)
}

func EncodeEc2TrafficMirrorFilterRule_DestinationPortRange_ToPort(p *DestinationPortRange, vals map[string]cty.Value) {
	vals["to_port"] = cty.IntVal(p.ToPort)
}