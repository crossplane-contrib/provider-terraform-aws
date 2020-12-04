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

func EncodeCloudwatchDashboard(r CloudwatchDashboard) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchDashboard_DashboardBody(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchDashboard_DashboardName(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchDashboard_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchDashboard_DashboardArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchDashboard_DashboardBody(p CloudwatchDashboardParameters, vals map[string]cty.Value) {
	vals["dashboard_body"] = cty.StringVal(p.DashboardBody)
}

func EncodeCloudwatchDashboard_DashboardName(p CloudwatchDashboardParameters, vals map[string]cty.Value) {
	vals["dashboard_name"] = cty.StringVal(p.DashboardName)
}

func EncodeCloudwatchDashboard_Id(p CloudwatchDashboardParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchDashboard_DashboardArn(p CloudwatchDashboardObservation, vals map[string]cty.Value) {
	vals["dashboard_arn"] = cty.StringVal(p.DashboardArn)
}