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

func EncodeAutoscalingAttachment(r AutoscalingAttachment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAutoscalingAttachment_AlbTargetGroupArn(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingAttachment_AutoscalingGroupName(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingAttachment_Elb(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingAttachment_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeAutoscalingAttachment_AlbTargetGroupArn(p AutoscalingAttachmentParameters, vals map[string]cty.Value) {
	vals["alb_target_group_arn"] = cty.StringVal(p.AlbTargetGroupArn)
}

func EncodeAutoscalingAttachment_AutoscalingGroupName(p AutoscalingAttachmentParameters, vals map[string]cty.Value) {
	vals["autoscaling_group_name"] = cty.StringVal(p.AutoscalingGroupName)
}

func EncodeAutoscalingAttachment_Elb(p AutoscalingAttachmentParameters, vals map[string]cty.Value) {
	vals["elb"] = cty.StringVal(p.Elb)
}

func EncodeAutoscalingAttachment_Id(p AutoscalingAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}