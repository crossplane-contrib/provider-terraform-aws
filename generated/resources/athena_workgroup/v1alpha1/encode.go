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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*AthenaWorkgroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a AthenaWorkgroup.")
	}
	return EncodeAthenaWorkgroup(*r), nil
}

func EncodeAthenaWorkgroup(r AthenaWorkgroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAthenaWorkgroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeAthenaWorkgroup_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeAthenaWorkgroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeAthenaWorkgroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeAthenaWorkgroup_State(r.Spec.ForProvider, ctyVal)
	EncodeAthenaWorkgroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAthenaWorkgroup_Configuration(r.Spec.ForProvider.Configuration, ctyVal)
	EncodeAthenaWorkgroup_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeAthenaWorkgroup_Description(p AthenaWorkgroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeAthenaWorkgroup_ForceDestroy(p AthenaWorkgroupParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeAthenaWorkgroup_Id(p AthenaWorkgroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAthenaWorkgroup_Name(p AthenaWorkgroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAthenaWorkgroup_State(p AthenaWorkgroupParameters, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeAthenaWorkgroup_Tags(p AthenaWorkgroupParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAthenaWorkgroup_Configuration(p Configuration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAthenaWorkgroup_Configuration_PublishCloudwatchMetricsEnabled(p, ctyVal)
	EncodeAthenaWorkgroup_Configuration_BytesScannedCutoffPerQuery(p, ctyVal)
	EncodeAthenaWorkgroup_Configuration_EnforceWorkgroupConfiguration(p, ctyVal)
	EncodeAthenaWorkgroup_Configuration_ResultConfiguration(p.ResultConfiguration, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAthenaWorkgroup_Configuration_PublishCloudwatchMetricsEnabled(p Configuration, vals map[string]cty.Value) {
	vals["publish_cloudwatch_metrics_enabled"] = cty.BoolVal(p.PublishCloudwatchMetricsEnabled)
}

func EncodeAthenaWorkgroup_Configuration_BytesScannedCutoffPerQuery(p Configuration, vals map[string]cty.Value) {
	vals["bytes_scanned_cutoff_per_query"] = cty.NumberIntVal(p.BytesScannedCutoffPerQuery)
}

func EncodeAthenaWorkgroup_Configuration_EnforceWorkgroupConfiguration(p Configuration, vals map[string]cty.Value) {
	vals["enforce_workgroup_configuration"] = cty.BoolVal(p.EnforceWorkgroupConfiguration)
}

func EncodeAthenaWorkgroup_Configuration_ResultConfiguration(p ResultConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAthenaWorkgroup_Configuration_ResultConfiguration_OutputLocation(p, ctyVal)
	EncodeAthenaWorkgroup_Configuration_ResultConfiguration_EncryptionConfiguration(p.EncryptionConfiguration, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["result_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAthenaWorkgroup_Configuration_ResultConfiguration_OutputLocation(p ResultConfiguration, vals map[string]cty.Value) {
	vals["output_location"] = cty.StringVal(p.OutputLocation)
}

func EncodeAthenaWorkgroup_Configuration_ResultConfiguration_EncryptionConfiguration(p EncryptionConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAthenaWorkgroup_Configuration_ResultConfiguration_EncryptionConfiguration_EncryptionOption(p, ctyVal)
	EncodeAthenaWorkgroup_Configuration_ResultConfiguration_EncryptionConfiguration_KmsKeyArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAthenaWorkgroup_Configuration_ResultConfiguration_EncryptionConfiguration_EncryptionOption(p EncryptionConfiguration, vals map[string]cty.Value) {
	vals["encryption_option"] = cty.StringVal(p.EncryptionOption)
}

func EncodeAthenaWorkgroup_Configuration_ResultConfiguration_EncryptionConfiguration_KmsKeyArn(p EncryptionConfiguration, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeAthenaWorkgroup_Arn(p AthenaWorkgroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}