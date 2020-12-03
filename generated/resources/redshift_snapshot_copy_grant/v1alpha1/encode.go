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

package v1alpha1func EncodeRedshiftSnapshotCopyGrant(r RedshiftSnapshotCopyGrant) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeRedshiftSnapshotCopyGrant_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotCopyGrant_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotCopyGrant_SnapshotCopyGrantName(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotCopyGrant_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotCopyGrant_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeRedshiftSnapshotCopyGrant_Id(p *RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftSnapshotCopyGrant_KmsKeyId(p *RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeRedshiftSnapshotCopyGrant_SnapshotCopyGrantName(p *RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	vals["snapshot_copy_grant_name"] = cty.StringVal(p.SnapshotCopyGrantName)
}

func EncodeRedshiftSnapshotCopyGrant_Tags(p *RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRedshiftSnapshotCopyGrant_Arn(p *RedshiftSnapshotCopyGrantObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}