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
	r, ok := mr.(*TransferServer)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a TransferServer.")
	}
	return EncodeTransferServer(*r), nil
}

func EncodeTransferServer(r TransferServer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeTransferServer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeTransferServer_EndpointType(r.Spec.ForProvider, ctyVal)
	EncodeTransferServer_HostKey(r.Spec.ForProvider, ctyVal)
	EncodeTransferServer_IdentityProviderType(r.Spec.ForProvider, ctyVal)
	EncodeTransferServer_InvocationRole(r.Spec.ForProvider, ctyVal)
	EncodeTransferServer_LoggingRole(r.Spec.ForProvider, ctyVal)
	EncodeTransferServer_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeTransferServer_Id(r.Spec.ForProvider, ctyVal)
	EncodeTransferServer_Url(r.Spec.ForProvider, ctyVal)
	EncodeTransferServer_EndpointDetails(r.Spec.ForProvider.EndpointDetails, ctyVal)
	EncodeTransferServer_Arn(r.Status.AtProvider, ctyVal)
	EncodeTransferServer_HostKeyFingerprint(r.Status.AtProvider, ctyVal)
	EncodeTransferServer_Endpoint(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeTransferServer_Tags(p TransferServerParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeTransferServer_EndpointType(p TransferServerParameters, vals map[string]cty.Value) {
	vals["endpoint_type"] = cty.StringVal(p.EndpointType)
}

func EncodeTransferServer_HostKey(p TransferServerParameters, vals map[string]cty.Value) {
	vals["host_key"] = cty.StringVal(p.HostKey)
}

func EncodeTransferServer_IdentityProviderType(p TransferServerParameters, vals map[string]cty.Value) {
	vals["identity_provider_type"] = cty.StringVal(p.IdentityProviderType)
}

func EncodeTransferServer_InvocationRole(p TransferServerParameters, vals map[string]cty.Value) {
	vals["invocation_role"] = cty.StringVal(p.InvocationRole)
}

func EncodeTransferServer_LoggingRole(p TransferServerParameters, vals map[string]cty.Value) {
	vals["logging_role"] = cty.StringVal(p.LoggingRole)
}

func EncodeTransferServer_ForceDestroy(p TransferServerParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeTransferServer_Id(p TransferServerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeTransferServer_Url(p TransferServerParameters, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}

func EncodeTransferServer_EndpointDetails(p EndpointDetails, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeTransferServer_EndpointDetails_AddressAllocationIds(p, ctyVal)
	EncodeTransferServer_EndpointDetails_SubnetIds(p, ctyVal)
	EncodeTransferServer_EndpointDetails_VpcEndpointId(p, ctyVal)
	EncodeTransferServer_EndpointDetails_VpcId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["endpoint_details"] = cty.ListVal(valsForCollection)
}

func EncodeTransferServer_EndpointDetails_AddressAllocationIds(p EndpointDetails, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AddressAllocationIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["address_allocation_ids"] = cty.SetVal(colVals)
}

func EncodeTransferServer_EndpointDetails_SubnetIds(p EndpointDetails, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeTransferServer_EndpointDetails_VpcEndpointId(p EndpointDetails, vals map[string]cty.Value) {
	vals["vpc_endpoint_id"] = cty.StringVal(p.VpcEndpointId)
}

func EncodeTransferServer_EndpointDetails_VpcId(p EndpointDetails, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeTransferServer_Arn(p TransferServerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeTransferServer_HostKeyFingerprint(p TransferServerObservation, vals map[string]cty.Value) {
	vals["host_key_fingerprint"] = cty.StringVal(p.HostKeyFingerprint)
}

func EncodeTransferServer_Endpoint(p TransferServerObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}