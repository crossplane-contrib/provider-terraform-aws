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
	"github.com/zclconf/go-cty/cty"
)

func EncodeAutoscalingLifecycleHook(r AutoscalingLifecycleHook) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAutoscalingLifecycleHook_AutoscalingGroupName(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingLifecycleHook_LifecycleTransition(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingLifecycleHook_NotificationMetadata(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingLifecycleHook_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingLifecycleHook_DefaultResult(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingLifecycleHook_HeartbeatTimeout(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingLifecycleHook_Id(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingLifecycleHook_Name(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingLifecycleHook_NotificationTargetArn(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeAutoscalingLifecycleHook_AutoscalingGroupName(p AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	vals["autoscaling_group_name"] = cty.StringVal(p.AutoscalingGroupName)
}

func EncodeAutoscalingLifecycleHook_LifecycleTransition(p AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	vals["lifecycle_transition"] = cty.StringVal(p.LifecycleTransition)
}

func EncodeAutoscalingLifecycleHook_NotificationMetadata(p AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	vals["notification_metadata"] = cty.StringVal(p.NotificationMetadata)
}

func EncodeAutoscalingLifecycleHook_RoleArn(p AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeAutoscalingLifecycleHook_DefaultResult(p AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	vals["default_result"] = cty.StringVal(p.DefaultResult)
}

func EncodeAutoscalingLifecycleHook_HeartbeatTimeout(p AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	vals["heartbeat_timeout"] = cty.NumberIntVal(p.HeartbeatTimeout)
}

func EncodeAutoscalingLifecycleHook_Id(p AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAutoscalingLifecycleHook_Name(p AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAutoscalingLifecycleHook_NotificationTargetArn(p AutoscalingLifecycleHookParameters, vals map[string]cty.Value) {
	vals["notification_target_arn"] = cty.StringVal(p.NotificationTargetArn)
}