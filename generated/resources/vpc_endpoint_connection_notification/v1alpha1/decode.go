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
	r, ok := mr.(*VpcEndpointConnectionNotification)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVpcEndpointConnectionNotification(r, ctyValue)
}

func DecodeVpcEndpointConnectionNotification(prev *VpcEndpointConnectionNotification, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVpcEndpointConnectionNotification_VpcEndpointServiceId(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointConnectionNotification_ConnectionEvents(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointConnectionNotification_ConnectionNotificationArn(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointConnectionNotification_Id(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointConnectionNotification_VpcEndpointId(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointConnectionNotification_NotificationType(&new.Status.AtProvider, valMap)
	DecodeVpcEndpointConnectionNotification_State(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointConnectionNotification_VpcEndpointServiceId(p *VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	p.VpcEndpointServiceId = ctwhy.ValueAsString(vals["vpc_endpoint_service_id"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVpcEndpointConnectionNotification_ConnectionEvents(p *VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["connection_events"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ConnectionEvents = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointConnectionNotification_ConnectionNotificationArn(p *VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	p.ConnectionNotificationArn = ctwhy.ValueAsString(vals["connection_notification_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointConnectionNotification_Id(p *VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointConnectionNotification_VpcEndpointId(p *VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	p.VpcEndpointId = ctwhy.ValueAsString(vals["vpc_endpoint_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointConnectionNotification_NotificationType(p *VpcEndpointConnectionNotificationObservation, vals map[string]cty.Value) {
	p.NotificationType = ctwhy.ValueAsString(vals["notification_type"])
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointConnectionNotification_State(p *VpcEndpointConnectionNotificationObservation, vals map[string]cty.Value) {
	p.State = ctwhy.ValueAsString(vals["state"])
}