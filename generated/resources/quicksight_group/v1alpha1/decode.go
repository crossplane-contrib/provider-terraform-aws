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
	r, ok := mr.(*QuicksightGroup)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeQuicksightGroup(r, ctyValue)
}

func DecodeQuicksightGroup(prev *QuicksightGroup, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeQuicksightGroup_Description(&new.Spec.ForProvider, valMap)
	DecodeQuicksightGroup_GroupName(&new.Spec.ForProvider, valMap)
	DecodeQuicksightGroup_Id(&new.Spec.ForProvider, valMap)
	DecodeQuicksightGroup_Namespace(&new.Spec.ForProvider, valMap)
	DecodeQuicksightGroup_AwsAccountId(&new.Spec.ForProvider, valMap)
	DecodeQuicksightGroup_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeQuicksightGroup_Description(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeQuicksightGroup_GroupName(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	p.GroupName = ctwhy.ValueAsString(vals["group_name"])
}

func DecodeQuicksightGroup_Id(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeQuicksightGroup_Namespace(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	p.Namespace = ctwhy.ValueAsString(vals["namespace"])
}

func DecodeQuicksightGroup_AwsAccountId(p *QuicksightGroupParameters, vals map[string]cty.Value) {
	p.AwsAccountId = ctwhy.ValueAsString(vals["aws_account_id"])
}

func DecodeQuicksightGroup_Arn(p *QuicksightGroupObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}