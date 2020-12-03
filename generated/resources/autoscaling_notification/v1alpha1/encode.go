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

package v1alpha1func EncodeAutoscalingNotification(r AutoscalingNotification) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAutoscalingNotification_Id(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingNotification_Notifications(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingNotification_TopicArn(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingNotification_GroupNames(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeAutoscalingNotification_Id(p *AutoscalingNotificationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAutoscalingNotification_Notifications(p *AutoscalingNotificationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Notifications {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["notifications"] = cty.SetVal(colVals)
}

func EncodeAutoscalingNotification_TopicArn(p *AutoscalingNotificationParameters, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}

func EncodeAutoscalingNotification_GroupNames(p *AutoscalingNotificationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.GroupNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["group_names"] = cty.SetVal(colVals)
}