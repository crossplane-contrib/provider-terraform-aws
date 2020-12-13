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
	r, ok := mr.(*Codepipeline)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Codepipeline.")
	}
	return EncodeCodepipeline(*r), nil
}

func EncodeCodepipeline(r Codepipeline) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodepipeline_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodepipeline_Name(r.Spec.ForProvider, ctyVal)
	EncodeCodepipeline_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCodepipeline_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCodepipeline_ArtifactStore(r.Spec.ForProvider.ArtifactStore, ctyVal)
	EncodeCodepipeline_Stage(r.Spec.ForProvider.Stage, ctyVal)
	EncodeCodepipeline_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCodepipeline_Id(p CodepipelineParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodepipeline_Name(p CodepipelineParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodepipeline_RoleArn(p CodepipelineParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCodepipeline_Tags(p CodepipelineParameters, vals map[string]cty.Value) {
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

func EncodeCodepipeline_ArtifactStore(p []ArtifactStore, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeCodepipeline_ArtifactStore_Location(v, ctyVal)
		EncodeCodepipeline_ArtifactStore_Region(v, ctyVal)
		EncodeCodepipeline_ArtifactStore_Type(v, ctyVal)
		EncodeCodepipeline_ArtifactStore_EncryptionKey(v.EncryptionKey, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["artifact_store"] = cty.SetVal(valsForCollection)
}

func EncodeCodepipeline_ArtifactStore_Location(p ArtifactStore, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeCodepipeline_ArtifactStore_Region(p ArtifactStore, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeCodepipeline_ArtifactStore_Type(p ArtifactStore, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodepipeline_ArtifactStore_EncryptionKey(p EncryptionKey, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodepipeline_ArtifactStore_EncryptionKey_Id(p, ctyVal)
	EncodeCodepipeline_ArtifactStore_EncryptionKey_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption_key"] = cty.ListVal(valsForCollection)
}

func EncodeCodepipeline_ArtifactStore_EncryptionKey_Id(p EncryptionKey, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodepipeline_ArtifactStore_EncryptionKey_Type(p EncryptionKey, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodepipeline_Stage(p []Stage, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeCodepipeline_Stage_Name(v, ctyVal)
		EncodeCodepipeline_Stage_Action(v.Action, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["stage"] = cty.ListVal(valsForCollection)
}

func EncodeCodepipeline_Stage_Name(p Stage, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodepipeline_Stage_Action(p []Action, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeCodepipeline_Stage_Action_Namespace(v, ctyVal)
		EncodeCodepipeline_Stage_Action_Owner(v, ctyVal)
		EncodeCodepipeline_Stage_Action_Region(v, ctyVal)
		EncodeCodepipeline_Stage_Action_RunOrder(v, ctyVal)
		EncodeCodepipeline_Stage_Action_OutputArtifacts(v, ctyVal)
		EncodeCodepipeline_Stage_Action_Provider(v, ctyVal)
		EncodeCodepipeline_Stage_Action_RoleArn(v, ctyVal)
		EncodeCodepipeline_Stage_Action_Version(v, ctyVal)
		EncodeCodepipeline_Stage_Action_Category(v, ctyVal)
		EncodeCodepipeline_Stage_Action_Configuration(v, ctyVal)
		EncodeCodepipeline_Stage_Action_InputArtifacts(v, ctyVal)
		EncodeCodepipeline_Stage_Action_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeCodepipeline_Stage_Action_Namespace(p Action, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeCodepipeline_Stage_Action_Owner(p Action, vals map[string]cty.Value) {
	vals["owner"] = cty.StringVal(p.Owner)
}

func EncodeCodepipeline_Stage_Action_Region(p Action, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeCodepipeline_Stage_Action_RunOrder(p Action, vals map[string]cty.Value) {
	vals["run_order"] = cty.NumberIntVal(p.RunOrder)
}

func EncodeCodepipeline_Stage_Action_OutputArtifacts(p Action, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.OutputArtifacts {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["output_artifacts"] = cty.ListVal(colVals)
}

func EncodeCodepipeline_Stage_Action_Provider(p Action, vals map[string]cty.Value) {
	vals["provider"] = cty.StringVal(p.Provider)
}

func EncodeCodepipeline_Stage_Action_RoleArn(p Action, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCodepipeline_Stage_Action_Version(p Action, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeCodepipeline_Stage_Action_Category(p Action, vals map[string]cty.Value) {
	vals["category"] = cty.StringVal(p.Category)
}

func EncodeCodepipeline_Stage_Action_Configuration(p Action, vals map[string]cty.Value) {
	if len(p.Configuration) == 0 {
		vals["configuration"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Configuration {
		mVals[key] = cty.StringVal(value)
	}
	vals["configuration"] = cty.MapVal(mVals)
}

func EncodeCodepipeline_Stage_Action_InputArtifacts(p Action, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.InputArtifacts {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["input_artifacts"] = cty.ListVal(colVals)
}

func EncodeCodepipeline_Stage_Action_Name(p Action, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodepipeline_Arn(p CodepipelineObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}