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
	r, ok := mr.(*CurReportDefinition)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CurReportDefinition.")
	}
	return EncodeCurReportDefinition(*r), nil
}

func EncodeCurReportDefinition(r CurReportDefinition) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCurReportDefinition_Format(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_Id(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_RefreshClosedReports(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_ReportName(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_S3Region(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_TimeUnit(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_AdditionalArtifacts(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_Compression(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_S3Bucket(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_S3Prefix(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_AdditionalSchemaElements(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_ReportVersioning(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCurReportDefinition_Format(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["format"] = cty.StringVal(p.Format)
}

func EncodeCurReportDefinition_Id(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCurReportDefinition_RefreshClosedReports(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["refresh_closed_reports"] = cty.BoolVal(p.RefreshClosedReports)
}

func EncodeCurReportDefinition_ReportName(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["report_name"] = cty.StringVal(p.ReportName)
}

func EncodeCurReportDefinition_S3Region(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["s3_region"] = cty.StringVal(p.S3Region)
}

func EncodeCurReportDefinition_TimeUnit(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["time_unit"] = cty.StringVal(p.TimeUnit)
}

func EncodeCurReportDefinition_AdditionalArtifacts(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AdditionalArtifacts {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["additional_artifacts"] = cty.SetVal(colVals)
}

func EncodeCurReportDefinition_Compression(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["compression"] = cty.StringVal(p.Compression)
}

func EncodeCurReportDefinition_S3Bucket(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["s3_bucket"] = cty.StringVal(p.S3Bucket)
}

func EncodeCurReportDefinition_S3Prefix(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["s3_prefix"] = cty.StringVal(p.S3Prefix)
}

func EncodeCurReportDefinition_AdditionalSchemaElements(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AdditionalSchemaElements {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["additional_schema_elements"] = cty.SetVal(colVals)
}

func EncodeCurReportDefinition_ReportVersioning(p CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["report_versioning"] = cty.StringVal(p.ReportVersioning)
}