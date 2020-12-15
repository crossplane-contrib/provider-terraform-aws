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
	r, ok := mr.(*AutoscalingLifecycleHook)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeAutoscalingLifecycleHook(r, ctyValue)
}

func DecodeAutoscalingLifecycleHook(prev *AutoscalingLifecycleHook, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeAutoscalingLifecycleHook_LifecycleTransition(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingLifecycleHook_Name(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingLifecycleHook_NotificationMetadata(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingLifecycleHook_RoleArn(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingLifecycleHook_DefaultResult(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingLifecycleHook_HeartbeatTimeout(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingLifecycleHook_Id(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingLifecycleHook_NotificationTargetArn(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingLifecycleHook_AutoscalingGroupName(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingLifecycleHook_LifecycleTransition(p *AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	p.LifecycleTransition = ctwhy.ValueAsString(vals["lifecycle_transition"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingLifecycleHook_Name(p *AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingLifecycleHook_NotificationMetadata(p *AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	p.NotificationMetadata = ctwhy.ValueAsString(vals["notification_metadata"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingLifecycleHook_RoleArn(p *AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	p.RoleArn = ctwhy.ValueAsString(vals["role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingLifecycleHook_DefaultResult(p *AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	p.DefaultResult = ctwhy.ValueAsString(vals["default_result"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingLifecycleHook_HeartbeatTimeout(p *AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	p.HeartbeatTimeout = ctwhy.ValueAsInt64(vals["heartbeat_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingLifecycleHook_Id(p *AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingLifecycleHook_NotificationTargetArn(p *AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	p.NotificationTargetArn = ctwhy.ValueAsString(vals["notification_target_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingLifecycleHook_AutoscalingGroupName(p *AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	p.AutoscalingGroupName = ctwhy.ValueAsString(vals["autoscaling_group_name"])
}