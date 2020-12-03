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

package v1alpha1func EncodeDmsReplicationSubnetGroup(r DmsReplicationSubnetGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDmsReplicationSubnetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationSubnetGroup_ReplicationSubnetGroupDescription(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationSubnetGroup_ReplicationSubnetGroupId(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationSubnetGroup_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationSubnetGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationSubnetGroup_ReplicationSubnetGroupArn(r.Status.AtProvider, ctyVal)
	EncodeDmsReplicationSubnetGroup_VpcId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeDmsReplicationSubnetGroup_Id(p *DmsReplicationSubnetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDmsReplicationSubnetGroup_ReplicationSubnetGroupDescription(p *DmsReplicationSubnetGroupParameters, vals map[string]cty.Value) {
	vals["replication_subnet_group_description"] = cty.StringVal(p.ReplicationSubnetGroupDescription)
}

func EncodeDmsReplicationSubnetGroup_ReplicationSubnetGroupId(p *DmsReplicationSubnetGroupParameters, vals map[string]cty.Value) {
	vals["replication_subnet_group_id"] = cty.StringVal(p.ReplicationSubnetGroupId)
}

func EncodeDmsReplicationSubnetGroup_SubnetIds(p *DmsReplicationSubnetGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeDmsReplicationSubnetGroup_Tags(p *DmsReplicationSubnetGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDmsReplicationSubnetGroup_ReplicationSubnetGroupArn(p *DmsReplicationSubnetGroupObservation, vals map[string]cty.Value) {
	vals["replication_subnet_group_arn"] = cty.StringVal(p.ReplicationSubnetGroupArn)
}

func EncodeDmsReplicationSubnetGroup_VpcId(p *DmsReplicationSubnetGroupObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}