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

func EncodeEc2ClientVpnEndpoint(r Ec2ClientVpnEndpoint) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2ClientVpnEndpoint_ClientCidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_DnsServers(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_SplitTunnel(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_TransportProtocol(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_ServerCertificateArn(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_ConnectionLogOptions(r.Spec.ForProvider.ConnectionLogOptions, ctyVal)
	EncodeEc2ClientVpnEndpoint_AuthenticationOptions(r.Spec.ForProvider.AuthenticationOptions, ctyVal)
	EncodeEc2ClientVpnEndpoint_Arn(r.Status.AtProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_Status(r.Status.AtProvider, ctyVal)
	EncodeEc2ClientVpnEndpoint_DnsName(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2ClientVpnEndpoint_ClientCidrBlock(p Ec2ClientVpnEndpointParameters, vals map[string]cty.Value) {
	vals["client_cidr_block"] = cty.StringVal(p.ClientCidrBlock)
}

func EncodeEc2ClientVpnEndpoint_Description(p Ec2ClientVpnEndpointParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2ClientVpnEndpoint_DnsServers(p Ec2ClientVpnEndpointParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DnsServers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["dns_servers"] = cty.SetVal(colVals)
}

func EncodeEc2ClientVpnEndpoint_SplitTunnel(p Ec2ClientVpnEndpointParameters, vals map[string]cty.Value) {
	vals["split_tunnel"] = cty.BoolVal(p.SplitTunnel)
}

func EncodeEc2ClientVpnEndpoint_Tags(p Ec2ClientVpnEndpointParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2ClientVpnEndpoint_TransportProtocol(p Ec2ClientVpnEndpointParameters, vals map[string]cty.Value) {
	vals["transport_protocol"] = cty.StringVal(p.TransportProtocol)
}

func EncodeEc2ClientVpnEndpoint_Id(p Ec2ClientVpnEndpointParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2ClientVpnEndpoint_ServerCertificateArn(p Ec2ClientVpnEndpointParameters, vals map[string]cty.Value) {
	vals["server_certificate_arn"] = cty.StringVal(p.ServerCertificateArn)
}

func EncodeEc2ClientVpnEndpoint_ConnectionLogOptions(p ConnectionLogOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEc2ClientVpnEndpoint_ConnectionLogOptions_CloudwatchLogGroup(p, ctyVal)
	EncodeEc2ClientVpnEndpoint_ConnectionLogOptions_CloudwatchLogStream(p, ctyVal)
	EncodeEc2ClientVpnEndpoint_ConnectionLogOptions_Enabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["connection_log_options"] = cty.ListVal(valsForCollection)
}

func EncodeEc2ClientVpnEndpoint_ConnectionLogOptions_CloudwatchLogGroup(p ConnectionLogOptions, vals map[string]cty.Value) {
	vals["cloudwatch_log_group"] = cty.StringVal(p.CloudwatchLogGroup)
}

func EncodeEc2ClientVpnEndpoint_ConnectionLogOptions_CloudwatchLogStream(p ConnectionLogOptions, vals map[string]cty.Value) {
	vals["cloudwatch_log_stream"] = cty.StringVal(p.CloudwatchLogStream)
}

func EncodeEc2ClientVpnEndpoint_ConnectionLogOptions_Enabled(p ConnectionLogOptions, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeEc2ClientVpnEndpoint_AuthenticationOptions(p []AuthenticationOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEc2ClientVpnEndpoint_AuthenticationOptions_ActiveDirectoryId(v, ctyVal)
		EncodeEc2ClientVpnEndpoint_AuthenticationOptions_RootCertificateChainArn(v, ctyVal)
		EncodeEc2ClientVpnEndpoint_AuthenticationOptions_SamlProviderArn(v, ctyVal)
		EncodeEc2ClientVpnEndpoint_AuthenticationOptions_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["authentication_options"] = cty.ListVal(valsForCollection)
}

func EncodeEc2ClientVpnEndpoint_AuthenticationOptions_ActiveDirectoryId(p AuthenticationOptions, vals map[string]cty.Value) {
	vals["active_directory_id"] = cty.StringVal(p.ActiveDirectoryId)
}

func EncodeEc2ClientVpnEndpoint_AuthenticationOptions_RootCertificateChainArn(p AuthenticationOptions, vals map[string]cty.Value) {
	vals["root_certificate_chain_arn"] = cty.StringVal(p.RootCertificateChainArn)
}

func EncodeEc2ClientVpnEndpoint_AuthenticationOptions_SamlProviderArn(p AuthenticationOptions, vals map[string]cty.Value) {
	vals["saml_provider_arn"] = cty.StringVal(p.SamlProviderArn)
}

func EncodeEc2ClientVpnEndpoint_AuthenticationOptions_Type(p AuthenticationOptions, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEc2ClientVpnEndpoint_Arn(p Ec2ClientVpnEndpointObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEc2ClientVpnEndpoint_Status(p Ec2ClientVpnEndpointObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeEc2ClientVpnEndpoint_DnsName(p Ec2ClientVpnEndpointObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}