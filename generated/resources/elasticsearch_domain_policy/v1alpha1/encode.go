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

func EncodeElasticsearchDomainPolicy(r ElasticsearchDomainPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeElasticsearchDomainPolicy_AccessPolicies(r.Spec.ForProvider, ctyVal)
	EncodeElasticsearchDomainPolicy_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeElasticsearchDomainPolicy_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeElasticsearchDomainPolicy_AccessPolicies(p ElasticsearchDomainPolicyParameters, vals map[string]cty.Value) {
	vals["access_policies"] = cty.StringVal(p.AccessPolicies)
}

func EncodeElasticsearchDomainPolicy_DomainName(p ElasticsearchDomainPolicyParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeElasticsearchDomainPolicy_Id(p ElasticsearchDomainPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}