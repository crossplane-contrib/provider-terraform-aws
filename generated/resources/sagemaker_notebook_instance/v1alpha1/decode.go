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
	r, ok := mr.(*SagemakerNotebookInstance)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSagemakerNotebookInstance(r, ctyValue)
}

func DecodeSagemakerNotebookInstance(prev *SagemakerNotebookInstance, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSagemakerNotebookInstance_Name(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_RootAccess(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_SecurityGroups(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_SubnetId(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_Id(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_LifecycleConfigName(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_InstanceType(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_KmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_RoleArn(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_Tags(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_DirectInternetAccess(&new.Spec.ForProvider, valMap)
	DecodeSagemakerNotebookInstance_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_Name(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_RootAccess(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	p.RootAccess = ctwhy.ValueAsString(vals["root_access"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_SecurityGroups(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["security_groups"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SecurityGroups = goVals
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_SubnetId(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	p.SubnetId = ctwhy.ValueAsString(vals["subnet_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_Id(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_LifecycleConfigName(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	p.LifecycleConfigName = ctwhy.ValueAsString(vals["lifecycle_config_name"])
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_InstanceType(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	p.InstanceType = ctwhy.ValueAsString(vals["instance_type"])
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_KmsKeyId(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_RoleArn(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	p.RoleArn = ctwhy.ValueAsString(vals["role_arn"])
}

//primitiveMapTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_Tags(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_DirectInternetAccess(p *SagemakerNotebookInstanceParameters, vals map[string]cty.Value) {
	p.DirectInternetAccess = ctwhy.ValueAsString(vals["direct_internet_access"])
}

//primitiveTypeDecodeTemplate
func DecodeSagemakerNotebookInstance_Arn(p *SagemakerNotebookInstanceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}