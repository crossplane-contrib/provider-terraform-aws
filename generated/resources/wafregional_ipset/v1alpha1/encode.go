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
	r, ok := mr.(*WafregionalIpset)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WafregionalIpset.")
	}
	return EncodeWafregionalIpset(*r), nil
}

func EncodeWafregionalIpset(r WafregionalIpset) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalIpset_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalIpset_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalIpset_IpSetDescriptor(r.Spec.ForProvider.IpSetDescriptor, ctyVal)
	EncodeWafregionalIpset_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafregionalIpset_Id(p WafregionalIpsetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalIpset_Name(p WafregionalIpsetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalIpset_IpSetDescriptor(p IpSetDescriptor, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalIpset_IpSetDescriptor_Type(p, ctyVal)
	EncodeWafregionalIpset_IpSetDescriptor_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ip_set_descriptor"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalIpset_IpSetDescriptor_Type(p IpSetDescriptor, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalIpset_IpSetDescriptor_Value(p IpSetDescriptor, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeWafregionalIpset_Arn(p WafregionalIpsetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}