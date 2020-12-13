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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*GlobalacceleratorListener)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GlobalacceleratorListener.")
	}
	return EncodeGlobalacceleratorListener(*r), nil
}

func EncodeGlobalacceleratorListener(r GlobalacceleratorListener) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlobalacceleratorListener_AcceleratorArn(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorListener_ClientAffinity(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorListener_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorListener_Protocol(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorListener_PortRange(r.Spec.ForProvider.PortRange, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeGlobalacceleratorListener_AcceleratorArn(p GlobalacceleratorListenerParameters, vals map[string]cty.Value) {
	vals["accelerator_arn"] = cty.StringVal(p.AcceleratorArn)
}

func EncodeGlobalacceleratorListener_ClientAffinity(p GlobalacceleratorListenerParameters, vals map[string]cty.Value) {
	vals["client_affinity"] = cty.StringVal(p.ClientAffinity)
}

func EncodeGlobalacceleratorListener_Id(p GlobalacceleratorListenerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlobalacceleratorListener_Protocol(p GlobalacceleratorListenerParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeGlobalacceleratorListener_PortRange(p []PortRange, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGlobalacceleratorListener_PortRange_FromPort(v, ctyVal)
		EncodeGlobalacceleratorListener_PortRange_ToPort(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["port_range"] = cty.SetVal(valsForCollection)
}

func EncodeGlobalacceleratorListener_PortRange_FromPort(p PortRange, vals map[string]cty.Value) {
	vals["from_port"] = cty.NumberIntVal(p.FromPort)
}

func EncodeGlobalacceleratorListener_PortRange_ToPort(p PortRange, vals map[string]cty.Value) {
	vals["to_port"] = cty.NumberIntVal(p.ToPort)
}