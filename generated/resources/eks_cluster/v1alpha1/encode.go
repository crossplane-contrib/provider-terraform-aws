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
	r, ok := mr.(*EksCluster)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EksCluster.")
	}
	return EncodeEksCluster(*r), nil
}

func EncodeEksCluster(r EksCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEksCluster_Name(r.Spec.ForProvider, ctyVal)
	EncodeEksCluster_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeEksCluster_Version(r.Spec.ForProvider, ctyVal)
	EncodeEksCluster_EnabledClusterLogTypes(r.Spec.ForProvider, ctyVal)
	EncodeEksCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeEksCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEksCluster_EncryptionConfig(r.Spec.ForProvider.EncryptionConfig, ctyVal)
	EncodeEksCluster_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeEksCluster_VpcConfig(r.Spec.ForProvider.VpcConfig, ctyVal)
	EncodeEksCluster_PlatformVersion(r.Status.AtProvider, ctyVal)
	EncodeEksCluster_CreatedAt(r.Status.AtProvider, ctyVal)
	EncodeEksCluster_Endpoint(r.Status.AtProvider, ctyVal)
	EncodeEksCluster_Identity(r.Status.AtProvider.Identity, ctyVal)
	EncodeEksCluster_Status(r.Status.AtProvider, ctyVal)
	EncodeEksCluster_Arn(r.Status.AtProvider, ctyVal)
	EncodeEksCluster_CertificateAuthority(r.Status.AtProvider.CertificateAuthority, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeEksCluster_Name(p EksClusterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEksCluster_RoleArn(p EksClusterParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeEksCluster_Version(p EksClusterParameters, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeEksCluster_EnabledClusterLogTypes(p EksClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EnabledClusterLogTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["enabled_cluster_log_types"] = cty.SetVal(colVals)
}

func EncodeEksCluster_Id(p EksClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEksCluster_Tags(p EksClusterParameters, vals map[string]cty.Value) {
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

func EncodeEksCluster_EncryptionConfig(p EncryptionConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEksCluster_EncryptionConfig_Resources(p, ctyVal)
	EncodeEksCluster_EncryptionConfig_Provider(p.Provider, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption_config"] = cty.ListVal(valsForCollection)
}

func EncodeEksCluster_EncryptionConfig_Resources(p EncryptionConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Resources {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["resources"] = cty.SetVal(colVals)
}

func EncodeEksCluster_EncryptionConfig_Provider(p Provider, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEksCluster_EncryptionConfig_Provider_KeyArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["provider"] = cty.ListVal(valsForCollection)
}

func EncodeEksCluster_EncryptionConfig_Provider_KeyArn(p Provider, vals map[string]cty.Value) {
	vals["key_arn"] = cty.StringVal(p.KeyArn)
}

func EncodeEksCluster_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeEksCluster_Timeouts_Update(p, ctyVal)
	EncodeEksCluster_Timeouts_Create(p, ctyVal)
	EncodeEksCluster_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeEksCluster_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeEksCluster_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeEksCluster_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeEksCluster_VpcConfig(p VpcConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEksCluster_VpcConfig_SubnetIds(p, ctyVal)
	EncodeEksCluster_VpcConfig_VpcId(p, ctyVal)
	EncodeEksCluster_VpcConfig_ClusterSecurityGroupId(p, ctyVal)
	EncodeEksCluster_VpcConfig_EndpointPrivateAccess(p, ctyVal)
	EncodeEksCluster_VpcConfig_EndpointPublicAccess(p, ctyVal)
	EncodeEksCluster_VpcConfig_PublicAccessCidrs(p, ctyVal)
	EncodeEksCluster_VpcConfig_SecurityGroupIds(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["vpc_config"] = cty.ListVal(valsForCollection)
}

func EncodeEksCluster_VpcConfig_SubnetIds(p VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeEksCluster_VpcConfig_VpcId(p VpcConfig, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeEksCluster_VpcConfig_ClusterSecurityGroupId(p VpcConfig, vals map[string]cty.Value) {
	vals["cluster_security_group_id"] = cty.StringVal(p.ClusterSecurityGroupId)
}

func EncodeEksCluster_VpcConfig_EndpointPrivateAccess(p VpcConfig, vals map[string]cty.Value) {
	vals["endpoint_private_access"] = cty.BoolVal(p.EndpointPrivateAccess)
}

func EncodeEksCluster_VpcConfig_EndpointPublicAccess(p VpcConfig, vals map[string]cty.Value) {
	vals["endpoint_public_access"] = cty.BoolVal(p.EndpointPublicAccess)
}

func EncodeEksCluster_VpcConfig_PublicAccessCidrs(p VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PublicAccessCidrs {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["public_access_cidrs"] = cty.SetVal(colVals)
}

func EncodeEksCluster_VpcConfig_SecurityGroupIds(p VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeEksCluster_PlatformVersion(p EksClusterObservation, vals map[string]cty.Value) {
	vals["platform_version"] = cty.StringVal(p.PlatformVersion)
}

func EncodeEksCluster_CreatedAt(p EksClusterObservation, vals map[string]cty.Value) {
	vals["created_at"] = cty.StringVal(p.CreatedAt)
}

func EncodeEksCluster_Endpoint(p EksClusterObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeEksCluster_Identity(p []Identity, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEksCluster_Identity_Oidc(v.Oidc, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["identity"] = cty.ListVal(valsForCollection)
}

func EncodeEksCluster_Identity_Oidc(p []Oidc, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEksCluster_Identity_Oidc_Issuer(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["oidc"] = cty.ListVal(valsForCollection)
}

func EncodeEksCluster_Identity_Oidc_Issuer(p Oidc, vals map[string]cty.Value) {
	vals["issuer"] = cty.StringVal(p.Issuer)
}

func EncodeEksCluster_Status(p EksClusterObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeEksCluster_Arn(p EksClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEksCluster_CertificateAuthority(p []CertificateAuthority, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEksCluster_CertificateAuthority_Data(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["certificate_authority"] = cty.ListVal(valsForCollection)
}

func EncodeEksCluster_CertificateAuthority_Data(p CertificateAuthority, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}