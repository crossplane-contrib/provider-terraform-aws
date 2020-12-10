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
	r, ok := mr.(*ConfigConfigurationRecorder)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ConfigConfigurationRecorder.")
	}
	return EncodeConfigConfigurationRecorder(*r), nil
}

func EncodeConfigConfigurationRecorder(r ConfigConfigurationRecorder) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeConfigConfigurationRecorder_Name(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigurationRecorder_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigurationRecorder_Id(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigurationRecorder_RecordingGroup(r.Spec.ForProvider.RecordingGroup, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeConfigConfigurationRecorder_Name(p ConfigConfigurationRecorderParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeConfigConfigurationRecorder_RoleArn(p ConfigConfigurationRecorderParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeConfigConfigurationRecorder_Id(p ConfigConfigurationRecorderParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeConfigConfigurationRecorder_RecordingGroup(p RecordingGroup, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeConfigConfigurationRecorder_RecordingGroup_AllSupported(p, ctyVal)
	EncodeConfigConfigurationRecorder_RecordingGroup_IncludeGlobalResourceTypes(p, ctyVal)
	EncodeConfigConfigurationRecorder_RecordingGroup_ResourceTypes(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["recording_group"] = cty.ListVal(valsForCollection)
}

func EncodeConfigConfigurationRecorder_RecordingGroup_AllSupported(p RecordingGroup, vals map[string]cty.Value) {
	vals["all_supported"] = cty.BoolVal(p.AllSupported)
}

func EncodeConfigConfigurationRecorder_RecordingGroup_IncludeGlobalResourceTypes(p RecordingGroup, vals map[string]cty.Value) {
	vals["include_global_resource_types"] = cty.BoolVal(p.IncludeGlobalResourceTypes)
}

func EncodeConfigConfigurationRecorder_RecordingGroup_ResourceTypes(p RecordingGroup, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ResourceTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["resource_types"] = cty.SetVal(colVals)
}