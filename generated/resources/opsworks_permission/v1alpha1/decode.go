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
	r, ok := mr.(*OpsworksPermission)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeOpsworksPermission(r, ctyValue)
}

func DecodeOpsworksPermission(prev *OpsworksPermission, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeOpsworksPermission_AllowSsh(&new.Spec.ForProvider, valMap)
	DecodeOpsworksPermission_AllowSudo(&new.Spec.ForProvider, valMap)
	DecodeOpsworksPermission_Id(&new.Spec.ForProvider, valMap)
	DecodeOpsworksPermission_Level(&new.Spec.ForProvider, valMap)
	DecodeOpsworksPermission_StackId(&new.Spec.ForProvider, valMap)
	DecodeOpsworksPermission_UserArn(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeOpsworksPermission_AllowSsh(p *OpsworksPermissionParameters, vals map[string]cty.Value) {
	p.AllowSsh = ctwhy.ValueAsBool(vals["allow_ssh"])
}

func DecodeOpsworksPermission_AllowSudo(p *OpsworksPermissionParameters, vals map[string]cty.Value) {
	p.AllowSudo = ctwhy.ValueAsBool(vals["allow_sudo"])
}

func DecodeOpsworksPermission_Id(p *OpsworksPermissionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeOpsworksPermission_Level(p *OpsworksPermissionParameters, vals map[string]cty.Value) {
	p.Level = ctwhy.ValueAsString(vals["level"])
}

func DecodeOpsworksPermission_StackId(p *OpsworksPermissionParameters, vals map[string]cty.Value) {
	p.StackId = ctwhy.ValueAsString(vals["stack_id"])
}

func DecodeOpsworksPermission_UserArn(p *OpsworksPermissionParameters, vals map[string]cty.Value) {
	p.UserArn = ctwhy.ValueAsString(vals["user_arn"])
}