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
	r, ok := mr.(*DbProxyTarget)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDbProxyTarget(r, ctyValue)
}

func DecodeDbProxyTarget(prev *DbProxyTarget, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDbProxyTarget_DbClusterIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDbProxyTarget_DbInstanceIdentifier(&new.Spec.ForProvider, valMap)
	DecodeDbProxyTarget_DbProxyName(&new.Spec.ForProvider, valMap)
	DecodeDbProxyTarget_Id(&new.Spec.ForProvider, valMap)
	DecodeDbProxyTarget_TargetGroupName(&new.Spec.ForProvider, valMap)
	DecodeDbProxyTarget_TrackedClusterId(&new.Status.AtProvider, valMap)
	DecodeDbProxyTarget_Endpoint(&new.Status.AtProvider, valMap)
	DecodeDbProxyTarget_Port(&new.Status.AtProvider, valMap)
	DecodeDbProxyTarget_RdsResourceId(&new.Status.AtProvider, valMap)
	DecodeDbProxyTarget_TargetArn(&new.Status.AtProvider, valMap)
	DecodeDbProxyTarget_Type(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDbProxyTarget_DbClusterIdentifier(p *DbProxyTargetParameters, vals map[string]cty.Value) {
	p.DbClusterIdentifier = ctwhy.ValueAsString(vals["db_cluster_identifier"])
}

func DecodeDbProxyTarget_DbInstanceIdentifier(p *DbProxyTargetParameters, vals map[string]cty.Value) {
	p.DbInstanceIdentifier = ctwhy.ValueAsString(vals["db_instance_identifier"])
}

func DecodeDbProxyTarget_DbProxyName(p *DbProxyTargetParameters, vals map[string]cty.Value) {
	p.DbProxyName = ctwhy.ValueAsString(vals["db_proxy_name"])
}

func DecodeDbProxyTarget_Id(p *DbProxyTargetParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDbProxyTarget_TargetGroupName(p *DbProxyTargetParameters, vals map[string]cty.Value) {
	p.TargetGroupName = ctwhy.ValueAsString(vals["target_group_name"])
}

func DecodeDbProxyTarget_TrackedClusterId(p *DbProxyTargetObservation, vals map[string]cty.Value) {
	p.TrackedClusterId = ctwhy.ValueAsString(vals["tracked_cluster_id"])
}

func DecodeDbProxyTarget_Endpoint(p *DbProxyTargetObservation, vals map[string]cty.Value) {
	p.Endpoint = ctwhy.ValueAsString(vals["endpoint"])
}

func DecodeDbProxyTarget_Port(p *DbProxyTargetObservation, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

func DecodeDbProxyTarget_RdsResourceId(p *DbProxyTargetObservation, vals map[string]cty.Value) {
	p.RdsResourceId = ctwhy.ValueAsString(vals["rds_resource_id"])
}

func DecodeDbProxyTarget_TargetArn(p *DbProxyTargetObservation, vals map[string]cty.Value) {
	p.TargetArn = ctwhy.ValueAsString(vals["target_arn"])
}

func DecodeDbProxyTarget_Type(p *DbProxyTargetObservation, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}