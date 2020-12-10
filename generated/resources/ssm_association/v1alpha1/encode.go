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
	r, ok := mr.(*SsmAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SsmAssociation.")
	}
	return EncodeSsmAssociation(*r), nil
}

func EncodeSsmAssociation(r SsmAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSsmAssociation_AssociationName(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_ComplianceSeverity(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_DocumentVersion(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_ScheduleExpression(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_AutomationTargetParameterName(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_InstanceId(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_MaxConcurrency(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_MaxErrors(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_Parameters(r.Spec.ForProvider, ctyVal)
	EncodeSsmAssociation_OutputLocation(r.Spec.ForProvider.OutputLocation, ctyVal)
	EncodeSsmAssociation_Targets(r.Spec.ForProvider.Targets, ctyVal)
	EncodeSsmAssociation_AssociationId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSsmAssociation_AssociationName(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["association_name"] = cty.StringVal(p.AssociationName)
}

func EncodeSsmAssociation_ComplianceSeverity(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["compliance_severity"] = cty.StringVal(p.ComplianceSeverity)
}

func EncodeSsmAssociation_DocumentVersion(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["document_version"] = cty.StringVal(p.DocumentVersion)
}

func EncodeSsmAssociation_Id(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmAssociation_ScheduleExpression(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["schedule_expression"] = cty.StringVal(p.ScheduleExpression)
}

func EncodeSsmAssociation_AutomationTargetParameterName(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["automation_target_parameter_name"] = cty.StringVal(p.AutomationTargetParameterName)
}

func EncodeSsmAssociation_InstanceId(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["instance_id"] = cty.StringVal(p.InstanceId)
}

func EncodeSsmAssociation_MaxConcurrency(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["max_concurrency"] = cty.StringVal(p.MaxConcurrency)
}

func EncodeSsmAssociation_MaxErrors(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["max_errors"] = cty.StringVal(p.MaxErrors)
}

func EncodeSsmAssociation_Name(p SsmAssociationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmAssociation_Parameters(p SsmAssociationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeSsmAssociation_OutputLocation(p OutputLocation, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSsmAssociation_OutputLocation_S3BucketName(p, ctyVal)
	EncodeSsmAssociation_OutputLocation_S3KeyPrefix(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["output_location"] = cty.ListVal(valsForCollection)
}

func EncodeSsmAssociation_OutputLocation_S3BucketName(p OutputLocation, vals map[string]cty.Value) {
	vals["s3_bucket_name"] = cty.StringVal(p.S3BucketName)
}

func EncodeSsmAssociation_OutputLocation_S3KeyPrefix(p OutputLocation, vals map[string]cty.Value) {
	vals["s3_key_prefix"] = cty.StringVal(p.S3KeyPrefix)
}

func EncodeSsmAssociation_Targets(p []Targets, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeSsmAssociation_Targets_Values(v, ctyVal)
		EncodeSsmAssociation_Targets_Key(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["targets"] = cty.ListVal(valsForCollection)
}

func EncodeSsmAssociation_Targets_Values(p Targets, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}

func EncodeSsmAssociation_Targets_Key(p Targets, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeSsmAssociation_AssociationId(p SsmAssociationObservation, vals map[string]cty.Value) {
	vals["association_id"] = cty.StringVal(p.AssociationId)
}