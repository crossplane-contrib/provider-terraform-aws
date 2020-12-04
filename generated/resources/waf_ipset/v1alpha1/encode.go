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

func EncodeWafIpset(r WafIpset) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafIpset_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafIpset_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafIpset_IpSetDescriptors(r.Spec.ForProvider.IpSetDescriptors, ctyVal)
	EncodeWafIpset_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafIpset_Name(p WafIpsetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafIpset_Id(p WafIpsetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafIpset_IpSetDescriptors(p IpSetDescriptors, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafIpset_IpSetDescriptors_Type(p, ctyVal)
	EncodeWafIpset_IpSetDescriptors_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ip_set_descriptors"] = cty.SetVal(valsForCollection)
}

func EncodeWafIpset_IpSetDescriptors_Type(p IpSetDescriptors, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafIpset_IpSetDescriptors_Value(p IpSetDescriptors, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeWafIpset_Arn(p WafIpsetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}