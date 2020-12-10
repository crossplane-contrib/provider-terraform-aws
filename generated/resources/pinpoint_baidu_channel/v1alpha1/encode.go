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
	r, ok := mr.(*PinpointBaiduChannel)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a PinpointBaiduChannel.")
	}
	return EncodePinpointBaiduChannel(*r), nil
}

func EncodePinpointBaiduChannel(r PinpointBaiduChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointBaiduChannel_SecretKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointBaiduChannel_ApiKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointBaiduChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointBaiduChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointBaiduChannel_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodePinpointBaiduChannel_SecretKey(p PinpointBaiduChannelParameters, vals map[string]cty.Value) {
	vals["secret_key"] = cty.StringVal(p.SecretKey)
}

func EncodePinpointBaiduChannel_ApiKey(p PinpointBaiduChannelParameters, vals map[string]cty.Value) {
	vals["api_key"] = cty.StringVal(p.ApiKey)
}

func EncodePinpointBaiduChannel_ApplicationId(p PinpointBaiduChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointBaiduChannel_Enabled(p PinpointBaiduChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointBaiduChannel_Id(p PinpointBaiduChannelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}