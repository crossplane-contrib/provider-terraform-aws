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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*DatasyncLocationFsxWindowsFileSystem)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDatasyncLocationFsxWindowsFileSystem(r, ctyValue)
}

func DecodeDatasyncLocationFsxWindowsFileSystem(prev *DatasyncLocationFsxWindowsFileSystem, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDatasyncLocationFsxWindowsFileSystem_User(&new.Spec.ForProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_Domain(&new.Spec.ForProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_FsxFilesystemArn(&new.Spec.ForProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_Id(&new.Spec.ForProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_SecurityGroupArns(&new.Spec.ForProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_Subdirectory(&new.Spec.ForProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_Password(&new.Spec.ForProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_Tags(&new.Spec.ForProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_Arn(&new.Status.AtProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_CreationTime(&new.Status.AtProvider, valMap)
	DecodeDatasyncLocationFsxWindowsFileSystem_Uri(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDatasyncLocationFsxWindowsFileSystem_User(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	p.User = ctwhy.ValueAsString(vals["user"])
}

func DecodeDatasyncLocationFsxWindowsFileSystem_Domain(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	p.Domain = ctwhy.ValueAsString(vals["domain"])
}

func DecodeDatasyncLocationFsxWindowsFileSystem_FsxFilesystemArn(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	p.FsxFilesystemArn = ctwhy.ValueAsString(vals["fsx_filesystem_arn"])
}

func DecodeDatasyncLocationFsxWindowsFileSystem_Id(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDatasyncLocationFsxWindowsFileSystem_SecurityGroupArns(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["security_group_arns"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SecurityGroupArns = goVals
}

func DecodeDatasyncLocationFsxWindowsFileSystem_Subdirectory(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	p.Subdirectory = ctwhy.ValueAsString(vals["subdirectory"])
}

func DecodeDatasyncLocationFsxWindowsFileSystem_Password(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	p.Password = ctwhy.ValueAsString(vals["password"])
}

func DecodeDatasyncLocationFsxWindowsFileSystem_Tags(p *DatasyncLocationFsxWindowsFileSystemParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeDatasyncLocationFsxWindowsFileSystem_Arn(p *DatasyncLocationFsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeDatasyncLocationFsxWindowsFileSystem_CreationTime(p *DatasyncLocationFsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	p.CreationTime = ctwhy.ValueAsString(vals["creation_time"])
}

func DecodeDatasyncLocationFsxWindowsFileSystem_Uri(p *DatasyncLocationFsxWindowsFileSystemObservation, vals map[string]cty.Value) {
	p.Uri = ctwhy.ValueAsString(vals["uri"])
}