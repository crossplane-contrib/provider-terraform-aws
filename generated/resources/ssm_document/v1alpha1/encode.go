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
	r, ok := mr.(*SsmDocument)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SsmDocument.")
	}
	return EncodeSsmDocument(*r), nil
}

func EncodeSsmDocument(r SsmDocument) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSsmDocument_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmDocument_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmDocument_Content(r.Spec.ForProvider, ctyVal)
	EncodeSsmDocument_TargetType(r.Spec.ForProvider, ctyVal)
	EncodeSsmDocument_DocumentFormat(r.Spec.ForProvider, ctyVal)
	EncodeSsmDocument_Permissions(r.Spec.ForProvider, ctyVal)
	EncodeSsmDocument_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSsmDocument_DocumentType(r.Spec.ForProvider, ctyVal)
	EncodeSsmDocument_AttachmentsSource(r.Spec.ForProvider.AttachmentsSource, ctyVal)
	EncodeSsmDocument_LatestVersion(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_DefaultVersion(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_DocumentVersion(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_Hash(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_Parameter(r.Status.AtProvider.Parameter, ctyVal)
	EncodeSsmDocument_Arn(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_Description(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_SchemaVersion(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_Owner(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_Status(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_HashType(r.Status.AtProvider, ctyVal)
	EncodeSsmDocument_PlatformTypes(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSsmDocument_Id(p SsmDocumentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmDocument_Name(p SsmDocumentParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmDocument_Content(p SsmDocumentParameters, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeSsmDocument_TargetType(p SsmDocumentParameters, vals map[string]cty.Value) {
	vals["target_type"] = cty.StringVal(p.TargetType)
}

func EncodeSsmDocument_DocumentFormat(p SsmDocumentParameters, vals map[string]cty.Value) {
	vals["document_format"] = cty.StringVal(p.DocumentFormat)
}

func EncodeSsmDocument_Permissions(p SsmDocumentParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Permissions {
		mVals[key] = cty.StringVal(value)
	}
	vals["permissions"] = cty.MapVal(mVals)
}

func EncodeSsmDocument_Tags(p SsmDocumentParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSsmDocument_DocumentType(p SsmDocumentParameters, vals map[string]cty.Value) {
	vals["document_type"] = cty.StringVal(p.DocumentType)
}

func EncodeSsmDocument_AttachmentsSource(p AttachmentsSource, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSsmDocument_AttachmentsSource_Key(p, ctyVal)
	EncodeSsmDocument_AttachmentsSource_Name(p, ctyVal)
	EncodeSsmDocument_AttachmentsSource_Values(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["attachments_source"] = cty.ListVal(valsForCollection)
}

func EncodeSsmDocument_AttachmentsSource_Key(p AttachmentsSource, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeSsmDocument_AttachmentsSource_Name(p AttachmentsSource, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmDocument_AttachmentsSource_Values(p AttachmentsSource, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}

func EncodeSsmDocument_LatestVersion(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["latest_version"] = cty.StringVal(p.LatestVersion)
}

func EncodeSsmDocument_DefaultVersion(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["default_version"] = cty.StringVal(p.DefaultVersion)
}

func EncodeSsmDocument_DocumentVersion(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["document_version"] = cty.StringVal(p.DocumentVersion)
}

func EncodeSsmDocument_Hash(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["hash"] = cty.StringVal(p.Hash)
}

func EncodeSsmDocument_Parameter(p []Parameter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeSsmDocument_Parameter_DefaultValue(v, ctyVal)
		EncodeSsmDocument_Parameter_Description(v, ctyVal)
		EncodeSsmDocument_Parameter_Name(v, ctyVal)
		EncodeSsmDocument_Parameter_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameter"] = cty.ListVal(valsForCollection)
}

func EncodeSsmDocument_Parameter_DefaultValue(p Parameter, vals map[string]cty.Value) {
	vals["default_value"] = cty.StringVal(p.DefaultValue)
}

func EncodeSsmDocument_Parameter_Description(p Parameter, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSsmDocument_Parameter_Name(p Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmDocument_Parameter_Type(p Parameter, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeSsmDocument_Arn(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeSsmDocument_CreatedDate(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeSsmDocument_Description(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSsmDocument_SchemaVersion(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["schema_version"] = cty.StringVal(p.SchemaVersion)
}

func EncodeSsmDocument_Owner(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["owner"] = cty.StringVal(p.Owner)
}

func EncodeSsmDocument_Status(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeSsmDocument_HashType(p SsmDocumentObservation, vals map[string]cty.Value) {
	vals["hash_type"] = cty.StringVal(p.HashType)
}

func EncodeSsmDocument_PlatformTypes(p SsmDocumentObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PlatformTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["platform_types"] = cty.ListVal(colVals)
}