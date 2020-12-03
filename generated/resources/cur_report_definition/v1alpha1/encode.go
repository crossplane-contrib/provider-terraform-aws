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

package v1alpha1func EncodeCurReportDefinition(r CurReportDefinition) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCurReportDefinition_Id(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_RefreshClosedReports(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_ReportVersioning(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_S3Prefix(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_TimeUnit(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_AdditionalSchemaElements(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_Compression(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_ReportName(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_S3Bucket(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_S3Region(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_AdditionalArtifacts(r.Spec.ForProvider, ctyVal)
	EncodeCurReportDefinition_Format(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeCurReportDefinition_Id(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCurReportDefinition_RefreshClosedReports(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["refresh_closed_reports"] = cty.BoolVal(p.RefreshClosedReports)
}

func EncodeCurReportDefinition_ReportVersioning(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["report_versioning"] = cty.StringVal(p.ReportVersioning)
}

func EncodeCurReportDefinition_S3Prefix(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["s3_prefix"] = cty.StringVal(p.S3Prefix)
}

func EncodeCurReportDefinition_TimeUnit(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["time_unit"] = cty.StringVal(p.TimeUnit)
}

func EncodeCurReportDefinition_AdditionalSchemaElements(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AdditionalSchemaElements {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["additional_schema_elements"] = cty.SetVal(colVals)
}

func EncodeCurReportDefinition_Compression(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["compression"] = cty.StringVal(p.Compression)
}

func EncodeCurReportDefinition_ReportName(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["report_name"] = cty.StringVal(p.ReportName)
}

func EncodeCurReportDefinition_S3Bucket(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["s3_bucket"] = cty.StringVal(p.S3Bucket)
}

func EncodeCurReportDefinition_S3Region(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["s3_region"] = cty.StringVal(p.S3Region)
}

func EncodeCurReportDefinition_AdditionalArtifacts(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AdditionalArtifacts {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["additional_artifacts"] = cty.SetVal(colVals)
}

func EncodeCurReportDefinition_Format(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	vals["format"] = cty.StringVal(p.Format)
}