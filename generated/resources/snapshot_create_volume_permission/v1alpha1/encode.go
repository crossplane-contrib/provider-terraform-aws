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

package v1alpha1func EncodeSnapshotCreateVolumePermission(r SnapshotCreateVolumePermission) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSnapshotCreateVolumePermission_SnapshotId(r.Spec.ForProvider, ctyVal)
	EncodeSnapshotCreateVolumePermission_AccountId(r.Spec.ForProvider, ctyVal)
	EncodeSnapshotCreateVolumePermission_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeSnapshotCreateVolumePermission_SnapshotId(p *SnapshotCreateVolumePermissionParameters, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeSnapshotCreateVolumePermission_AccountId(p *SnapshotCreateVolumePermissionParameters, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeSnapshotCreateVolumePermission_Id(p *SnapshotCreateVolumePermissionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}