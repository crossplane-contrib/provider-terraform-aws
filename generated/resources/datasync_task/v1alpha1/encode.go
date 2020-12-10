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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*DatasyncTask)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DatasyncTask.")
	}
	return EncodeDatasyncTask(*r), nil
}

func EncodeDatasyncTask(r DatasyncTask) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncTask_SourceLocationArn(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncTask_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncTask_CloudwatchLogGroupArn(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncTask_DestinationLocationArn(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncTask_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncTask_Name(r.Spec.ForProvider, ctyVal)
	EncodeDatasyncTask_Options(r.Spec.ForProvider.Options, ctyVal)
	EncodeDatasyncTask_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDatasyncTask_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDatasyncTask_SourceLocationArn(p DatasyncTaskParameters, vals map[string]cty.Value) {
	vals["source_location_arn"] = cty.StringVal(p.SourceLocationArn)
}

func EncodeDatasyncTask_Tags(p DatasyncTaskParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDatasyncTask_CloudwatchLogGroupArn(p DatasyncTaskParameters, vals map[string]cty.Value) {
	vals["cloudwatch_log_group_arn"] = cty.StringVal(p.CloudwatchLogGroupArn)
}

func EncodeDatasyncTask_DestinationLocationArn(p DatasyncTaskParameters, vals map[string]cty.Value) {
	vals["destination_location_arn"] = cty.StringVal(p.DestinationLocationArn)
}

func EncodeDatasyncTask_Id(p DatasyncTaskParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatasyncTask_Name(p DatasyncTaskParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDatasyncTask_Options(p Options, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncTask_Options_PreserveDevices(p, ctyVal)
	EncodeDatasyncTask_Options_Uid(p, ctyVal)
	EncodeDatasyncTask_Options_VerifyMode(p, ctyVal)
	EncodeDatasyncTask_Options_Atime(p, ctyVal)
	EncodeDatasyncTask_Options_BytesPerSecond(p, ctyVal)
	EncodeDatasyncTask_Options_Gid(p, ctyVal)
	EncodeDatasyncTask_Options_Mtime(p, ctyVal)
	EncodeDatasyncTask_Options_PosixPermissions(p, ctyVal)
	EncodeDatasyncTask_Options_PreserveDeletedFiles(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["options"] = cty.ListVal(valsForCollection)
}

func EncodeDatasyncTask_Options_PreserveDevices(p Options, vals map[string]cty.Value) {
	vals["preserve_devices"] = cty.StringVal(p.PreserveDevices)
}

func EncodeDatasyncTask_Options_Uid(p Options, vals map[string]cty.Value) {
	vals["uid"] = cty.StringVal(p.Uid)
}

func EncodeDatasyncTask_Options_VerifyMode(p Options, vals map[string]cty.Value) {
	vals["verify_mode"] = cty.StringVal(p.VerifyMode)
}

func EncodeDatasyncTask_Options_Atime(p Options, vals map[string]cty.Value) {
	vals["atime"] = cty.StringVal(p.Atime)
}

func EncodeDatasyncTask_Options_BytesPerSecond(p Options, vals map[string]cty.Value) {
	vals["bytes_per_second"] = cty.NumberIntVal(p.BytesPerSecond)
}

func EncodeDatasyncTask_Options_Gid(p Options, vals map[string]cty.Value) {
	vals["gid"] = cty.StringVal(p.Gid)
}

func EncodeDatasyncTask_Options_Mtime(p Options, vals map[string]cty.Value) {
	vals["mtime"] = cty.StringVal(p.Mtime)
}

func EncodeDatasyncTask_Options_PosixPermissions(p Options, vals map[string]cty.Value) {
	vals["posix_permissions"] = cty.StringVal(p.PosixPermissions)
}

func EncodeDatasyncTask_Options_PreserveDeletedFiles(p Options, vals map[string]cty.Value) {
	vals["preserve_deleted_files"] = cty.StringVal(p.PreserveDeletedFiles)
}

func EncodeDatasyncTask_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDatasyncTask_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDatasyncTask_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDatasyncTask_Arn(p DatasyncTaskObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}