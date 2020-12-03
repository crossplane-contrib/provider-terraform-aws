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

package v1alpha1func EncodeDatasyncLocationFsxWindowsFileSystem(r DatasyncLocationFsxWindowsFileSystem) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDatasyncLocationFsxWindowsFileSystem_FsxFilesystemArn(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_Password(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_SecurityGroupArns(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_Domain(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_Subdirectory(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_User(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_CreationTime(r.Status.AtProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_Arn(r.Status.AtProvider, ctyVal)
	EncodeDatasyncLocationFsxWindowsFileSystem_Uri(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_FsxFilesystemArn(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["fsx_filesystem_arn"] = cty.StringVal(p.FsxFilesystemArn)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_Id(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_Password(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_SecurityGroupArns(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_arns"] = cty.SetVal(colVals)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_Domain(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_Subdirectory(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["subdirectory"] = cty.StringVal(p.Subdirectory)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_Tags(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_User(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	vals["user"] = cty.StringVal(p.User)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_CreationTime(p *DatasyncLocationFsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	vals["creation_time"] = cty.StringVal(p.CreationTime)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_Arn(p *DatasyncLocationFsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDatasyncLocationFsxWindowsFileSystem_Uri(p *DatasyncLocationFsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}