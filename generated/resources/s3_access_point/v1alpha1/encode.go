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
	r, ok := mr.(*S3AccessPoint)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a S3AccessPoint.")
	}
	return EncodeS3AccessPoint(*r), nil
}

func EncodeS3AccessPoint(r S3AccessPoint) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeS3AccessPoint_Policy(r.Spec.ForProvider, ctyVal)
	EncodeS3AccessPoint_AccountId(r.Spec.ForProvider, ctyVal)
	EncodeS3AccessPoint_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3AccessPoint_Name(r.Spec.ForProvider, ctyVal)
	EncodeS3AccessPoint_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3AccessPoint_PublicAccessBlockConfiguration(r.Spec.ForProvider.PublicAccessBlockConfiguration, ctyVal)
	EncodeS3AccessPoint_VpcConfiguration(r.Spec.ForProvider.VpcConfiguration, ctyVal)
	EncodeS3AccessPoint_Arn(r.Status.AtProvider, ctyVal)
	EncodeS3AccessPoint_HasPublicAccessPolicy(r.Status.AtProvider, ctyVal)
	EncodeS3AccessPoint_DomainName(r.Status.AtProvider, ctyVal)
	EncodeS3AccessPoint_NetworkOrigin(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeS3AccessPoint_Policy(p S3AccessPointParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeS3AccessPoint_AccountId(p S3AccessPointParameters, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeS3AccessPoint_Bucket(p S3AccessPointParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3AccessPoint_Name(p S3AccessPointParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeS3AccessPoint_Id(p S3AccessPointParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3AccessPoint_PublicAccessBlockConfiguration(p PublicAccessBlockConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3AccessPoint_PublicAccessBlockConfiguration_BlockPublicAcls(p, ctyVal)
	EncodeS3AccessPoint_PublicAccessBlockConfiguration_BlockPublicPolicy(p, ctyVal)
	EncodeS3AccessPoint_PublicAccessBlockConfiguration_IgnorePublicAcls(p, ctyVal)
	EncodeS3AccessPoint_PublicAccessBlockConfiguration_RestrictPublicBuckets(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["public_access_block_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeS3AccessPoint_PublicAccessBlockConfiguration_BlockPublicAcls(p PublicAccessBlockConfiguration, vals map[string]cty.Value) {
	vals["block_public_acls"] = cty.BoolVal(p.BlockPublicAcls)
}

func EncodeS3AccessPoint_PublicAccessBlockConfiguration_BlockPublicPolicy(p PublicAccessBlockConfiguration, vals map[string]cty.Value) {
	vals["block_public_policy"] = cty.BoolVal(p.BlockPublicPolicy)
}

func EncodeS3AccessPoint_PublicAccessBlockConfiguration_IgnorePublicAcls(p PublicAccessBlockConfiguration, vals map[string]cty.Value) {
	vals["ignore_public_acls"] = cty.BoolVal(p.IgnorePublicAcls)
}

func EncodeS3AccessPoint_PublicAccessBlockConfiguration_RestrictPublicBuckets(p PublicAccessBlockConfiguration, vals map[string]cty.Value) {
	vals["restrict_public_buckets"] = cty.BoolVal(p.RestrictPublicBuckets)
}

func EncodeS3AccessPoint_VpcConfiguration(p VpcConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3AccessPoint_VpcConfiguration_VpcId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["vpc_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeS3AccessPoint_VpcConfiguration_VpcId(p VpcConfiguration, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeS3AccessPoint_Arn(p S3AccessPointObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeS3AccessPoint_HasPublicAccessPolicy(p S3AccessPointObservation, vals map[string]cty.Value) {
	vals["has_public_access_policy"] = cty.BoolVal(p.HasPublicAccessPolicy)
}

func EncodeS3AccessPoint_DomainName(p S3AccessPointObservation, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeS3AccessPoint_NetworkOrigin(p S3AccessPointObservation, vals map[string]cty.Value) {
	vals["network_origin"] = cty.StringVal(p.NetworkOrigin)
}