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

func EncodeLbCookieStickinessPolicy(r LbCookieStickinessPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLbCookieStickinessPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeLbCookieStickinessPolicy_CookieExpirationPeriod(r.Spec.ForProvider, ctyVal)
	EncodeLbCookieStickinessPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeLbCookieStickinessPolicy_LbPort(r.Spec.ForProvider, ctyVal)
	EncodeLbCookieStickinessPolicy_LoadBalancer(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeLbCookieStickinessPolicy_Name(p LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLbCookieStickinessPolicy_CookieExpirationPeriod(p LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	vals["cookie_expiration_period"] = cty.NumberIntVal(p.CookieExpirationPeriod)
}

func EncodeLbCookieStickinessPolicy_Id(p LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLbCookieStickinessPolicy_LbPort(p LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	vals["lb_port"] = cty.NumberIntVal(p.LbPort)
}

func EncodeLbCookieStickinessPolicy_LoadBalancer(p LbCookieStickinessPolicyParameters, vals map[string]cty.Value) {
	vals["load_balancer"] = cty.StringVal(p.LoadBalancer)
}