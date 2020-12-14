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
	r, ok := mr.(*LambdaLayerVersion)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLambdaLayerVersion(r, ctyValue)
}

func DecodeLambdaLayerVersion(prev *LambdaLayerVersion, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLambdaLayerVersion_LicenseInfo(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_CompatibleRuntimes(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_Id(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_SourceCodeHash(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_Description(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_S3ObjectVersion(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_LayerName(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_S3Bucket(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_S3Key(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_Filename(&new.Spec.ForProvider, valMap)
	DecodeLambdaLayerVersion_SourceCodeSize(&new.Status.AtProvider, valMap)
	DecodeLambdaLayerVersion_LayerArn(&new.Status.AtProvider, valMap)
	DecodeLambdaLayerVersion_CreatedDate(&new.Status.AtProvider, valMap)
	DecodeLambdaLayerVersion_Version(&new.Status.AtProvider, valMap)
	DecodeLambdaLayerVersion_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeLambdaLayerVersion_LicenseInfo(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	p.LicenseInfo = ctwhy.ValueAsString(vals["license_info"])
}

func DecodeLambdaLayerVersion_CompatibleRuntimes(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["compatible_runtimes"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.CompatibleRuntimes = goVals
}

func DecodeLambdaLayerVersion_Id(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeLambdaLayerVersion_SourceCodeHash(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	p.SourceCodeHash = ctwhy.ValueAsString(vals["source_code_hash"])
}

func DecodeLambdaLayerVersion_Description(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeLambdaLayerVersion_S3ObjectVersion(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	p.S3ObjectVersion = ctwhy.ValueAsString(vals["s3_object_version"])
}

func DecodeLambdaLayerVersion_LayerName(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	p.LayerName = ctwhy.ValueAsString(vals["layer_name"])
}

func DecodeLambdaLayerVersion_S3Bucket(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	p.S3Bucket = ctwhy.ValueAsString(vals["s3_bucket"])
}

func DecodeLambdaLayerVersion_S3Key(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	p.S3Key = ctwhy.ValueAsString(vals["s3_key"])
}

func DecodeLambdaLayerVersion_Filename(p *LambdaLayerVersionParameters, vals map[string]cty.Value) {
	p.Filename = ctwhy.ValueAsString(vals["filename"])
}

func DecodeLambdaLayerVersion_SourceCodeSize(p *LambdaLayerVersionObservation, vals map[string]cty.Value) {
	p.SourceCodeSize = ctwhy.ValueAsInt64(vals["source_code_size"])
}

func DecodeLambdaLayerVersion_LayerArn(p *LambdaLayerVersionObservation, vals map[string]cty.Value) {
	p.LayerArn = ctwhy.ValueAsString(vals["layer_arn"])
}

func DecodeLambdaLayerVersion_CreatedDate(p *LambdaLayerVersionObservation, vals map[string]cty.Value) {
	p.CreatedDate = ctwhy.ValueAsString(vals["created_date"])
}

func DecodeLambdaLayerVersion_Version(p *LambdaLayerVersionObservation, vals map[string]cty.Value) {
	p.Version = ctwhy.ValueAsString(vals["version"])
}

func DecodeLambdaLayerVersion_Arn(p *LambdaLayerVersionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}