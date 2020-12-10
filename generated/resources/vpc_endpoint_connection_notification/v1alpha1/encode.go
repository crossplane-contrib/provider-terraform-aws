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
	r, ok := mr.(*VpcEndpointConnectionNotification)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VpcEndpointConnectionNotification.")
	}
	return EncodeVpcEndpointConnectionNotification(*r), nil
}

func EncodeVpcEndpointConnectionNotification(r VpcEndpointConnectionNotification) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcEndpointConnectionNotification_VpcEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointConnectionNotification_VpcEndpointServiceId(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointConnectionNotification_ConnectionEvents(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointConnectionNotification_ConnectionNotificationArn(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointConnectionNotification_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointConnectionNotification_NotificationType(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointConnectionNotification_State(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeVpcEndpointConnectionNotification_VpcEndpointId(p VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	vals["vpc_endpoint_id"] = cty.StringVal(p.VpcEndpointId)
}

func EncodeVpcEndpointConnectionNotification_VpcEndpointServiceId(p VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	vals["vpc_endpoint_service_id"] = cty.StringVal(p.VpcEndpointServiceId)
}

func EncodeVpcEndpointConnectionNotification_ConnectionEvents(p VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ConnectionEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["connection_events"] = cty.SetVal(colVals)
}

func EncodeVpcEndpointConnectionNotification_ConnectionNotificationArn(p VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	vals["connection_notification_arn"] = cty.StringVal(p.ConnectionNotificationArn)
}

func EncodeVpcEndpointConnectionNotification_Id(p VpcEndpointConnectionNotificationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcEndpointConnectionNotification_NotificationType(p VpcEndpointConnectionNotificationObservation, vals map[string]cty.Value) {
	vals["notification_type"] = cty.StringVal(p.NotificationType)
}

func EncodeVpcEndpointConnectionNotification_State(p VpcEndpointConnectionNotificationObservation, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}