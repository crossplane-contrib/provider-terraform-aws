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

func EncodeRoute53DelegationSet(r Route53DelegationSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53DelegationSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53DelegationSet_ReferenceName(r.Spec.ForProvider, ctyVal)
	EncodeRoute53DelegationSet_NameServers(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53DelegationSet_Id(p Route53DelegationSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53DelegationSet_ReferenceName(p Route53DelegationSetParameters, vals map[string]cty.Value) {
	vals["reference_name"] = cty.StringVal(p.ReferenceName)
}

func EncodeRoute53DelegationSet_NameServers(p Route53DelegationSetObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NameServers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["name_servers"] = cty.ListVal(colVals)
}