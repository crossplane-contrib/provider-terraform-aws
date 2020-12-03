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

package v1alpha1func EncodeEfsAccessPoint(r EfsAccessPoint) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEfsAccessPoint_FileSystemId(r.Spec.ForProvider, ctyVal)
	EncodeEfsAccessPoint_Id(r.Spec.ForProvider, ctyVal)
	EncodeEfsAccessPoint_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEfsAccessPoint_PosixUser(r.Spec.ForProvider.PosixUser, ctyVal)
	EncodeEfsAccessPoint_RootDirectory(r.Spec.ForProvider.RootDirectory, ctyVal)
	EncodeEfsAccessPoint_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeEfsAccessPoint_Arn(r.Status.AtProvider, ctyVal)
	EncodeEfsAccessPoint_FileSystemArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEfsAccessPoint_FileSystemId(p *EfsAccessPointParameters, vals map[string]cty.Value) {
	vals["file_system_id"] = cty.StringVal(p.FileSystemId)
}

func EncodeEfsAccessPoint_Id(p *EfsAccessPointParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEfsAccessPoint_Tags(p *EfsAccessPointParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEfsAccessPoint_PosixUser(p *PosixUser, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.PosixUser {
		ctyVal = make(map[string]cty.Value)
		EncodeEfsAccessPoint_PosixUser_SecondaryGids(v, ctyVal)
		EncodeEfsAccessPoint_PosixUser_Uid(v, ctyVal)
		EncodeEfsAccessPoint_PosixUser_Gid(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["posix_user"] = cty.ListVal(valsForCollection)
}

func EncodeEfsAccessPoint_PosixUser_SecondaryGids(p *PosixUser, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecondaryGids {
		colVals = append(colVals, cty.IntVal(value))
	}
	vals["secondary_gids"] = cty.SetVal(colVals)
}

func EncodeEfsAccessPoint_PosixUser_Uid(p *PosixUser, vals map[string]cty.Value) {
	vals["uid"] = cty.IntVal(p.Uid)
}

func EncodeEfsAccessPoint_PosixUser_Gid(p *PosixUser, vals map[string]cty.Value) {
	vals["gid"] = cty.IntVal(p.Gid)
}

func EncodeEfsAccessPoint_RootDirectory(p *RootDirectory, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RootDirectory {
		ctyVal = make(map[string]cty.Value)
		EncodeEfsAccessPoint_RootDirectory_Path(v, ctyVal)
		EncodeEfsAccessPoint_RootDirectory_CreationInfo(v.CreationInfo, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["root_directory"] = cty.ListVal(valsForCollection)
}

func EncodeEfsAccessPoint_RootDirectory_Path(p *RootDirectory, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeEfsAccessPoint_RootDirectory_CreationInfo(p *CreationInfo, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CreationInfo {
		ctyVal = make(map[string]cty.Value)
		EncodeEfsAccessPoint_RootDirectory_CreationInfo_OwnerGid(v, ctyVal)
		EncodeEfsAccessPoint_RootDirectory_CreationInfo_OwnerUid(v, ctyVal)
		EncodeEfsAccessPoint_RootDirectory_CreationInfo_Permissions(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["creation_info"] = cty.ListVal(valsForCollection)
}

func EncodeEfsAccessPoint_RootDirectory_CreationInfo_OwnerGid(p *CreationInfo, vals map[string]cty.Value) {
	vals["owner_gid"] = cty.IntVal(p.OwnerGid)
}

func EncodeEfsAccessPoint_RootDirectory_CreationInfo_OwnerUid(p *CreationInfo, vals map[string]cty.Value) {
	vals["owner_uid"] = cty.IntVal(p.OwnerUid)
}

func EncodeEfsAccessPoint_RootDirectory_CreationInfo_Permissions(p *CreationInfo, vals map[string]cty.Value) {
	vals["permissions"] = cty.StringVal(p.Permissions)
}

func EncodeEfsAccessPoint_OwnerId(p *EfsAccessPointObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeEfsAccessPoint_Arn(p *EfsAccessPointObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEfsAccessPoint_FileSystemArn(p *EfsAccessPointObservation, vals map[string]cty.Value) {
	vals["file_system_arn"] = cty.StringVal(p.FileSystemArn)
}