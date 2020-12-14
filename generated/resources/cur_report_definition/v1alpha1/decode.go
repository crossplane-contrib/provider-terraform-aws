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
	r, ok := mr.(*CurReportDefinition)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCurReportDefinition(r, ctyValue)
}

func DecodeCurReportDefinition(prev *CurReportDefinition, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCurReportDefinition_Id(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_ReportVersioning(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_S3Bucket(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_S3Prefix(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_TimeUnit(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_AdditionalSchemaElements(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_Compression(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_RefreshClosedReports(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_ReportName(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_S3Region(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_AdditionalArtifacts(&new.Spec.ForProvider, valMap)
	DecodeCurReportDefinition_Format(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeCurReportDefinition_Id(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeCurReportDefinition_ReportVersioning(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.ReportVersioning = ctwhy.ValueAsString(vals["report_versioning"])
}

func DecodeCurReportDefinition_S3Bucket(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.S3Bucket = ctwhy.ValueAsString(vals["s3_bucket"])
}

func DecodeCurReportDefinition_S3Prefix(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.S3Prefix = ctwhy.ValueAsString(vals["s3_prefix"])
}

func DecodeCurReportDefinition_TimeUnit(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.TimeUnit = ctwhy.ValueAsString(vals["time_unit"])
}

func DecodeCurReportDefinition_AdditionalSchemaElements(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["additional_schema_elements"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AdditionalSchemaElements = goVals
}

func DecodeCurReportDefinition_Compression(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.Compression = ctwhy.ValueAsString(vals["compression"])
}

func DecodeCurReportDefinition_RefreshClosedReports(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.RefreshClosedReports = ctwhy.ValueAsBool(vals["refresh_closed_reports"])
}

func DecodeCurReportDefinition_ReportName(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.ReportName = ctwhy.ValueAsString(vals["report_name"])
}

func DecodeCurReportDefinition_S3Region(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.S3Region = ctwhy.ValueAsString(vals["s3_region"])
}

func DecodeCurReportDefinition_AdditionalArtifacts(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["additional_artifacts"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AdditionalArtifacts = goVals
}

func DecodeCurReportDefinition_Format(p *CurReportDefinitionParameters, vals map[string]cty.Value) {
	p.Format = ctwhy.ValueAsString(vals["format"])
}