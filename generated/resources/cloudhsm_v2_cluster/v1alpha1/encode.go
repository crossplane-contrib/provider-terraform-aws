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

func EncodeCloudhsmV2Cluster(r CloudhsmV2Cluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudhsmV2Cluster_SourceBackupIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Cluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Cluster_HsmType(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Cluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Cluster_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeCloudhsmV2Cluster_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeCloudhsmV2Cluster_ClusterId(r.Status.AtProvider, ctyVal)
	EncodeCloudhsmV2Cluster_SecurityGroupId(r.Status.AtProvider, ctyVal)
	EncodeCloudhsmV2Cluster_VpcId(r.Status.AtProvider, ctyVal)
	EncodeCloudhsmV2Cluster_ClusterCertificates(r.Status.AtProvider.ClusterCertificates, ctyVal)
	EncodeCloudhsmV2Cluster_ClusterState(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudhsmV2Cluster_SourceBackupIdentifier(p CloudhsmV2ClusterParameters, vals map[string]cty.Value) {
	vals["source_backup_identifier"] = cty.StringVal(p.SourceBackupIdentifier)
}

func EncodeCloudhsmV2Cluster_Tags(p CloudhsmV2ClusterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudhsmV2Cluster_HsmType(p CloudhsmV2ClusterParameters, vals map[string]cty.Value) {
	vals["hsm_type"] = cty.StringVal(p.HsmType)
}

func EncodeCloudhsmV2Cluster_Id(p CloudhsmV2ClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudhsmV2Cluster_SubnetIds(p CloudhsmV2ClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeCloudhsmV2Cluster_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudhsmV2Cluster_Timeouts_Delete(p, ctyVal)
	EncodeCloudhsmV2Cluster_Timeouts_Update(p, ctyVal)
	EncodeCloudhsmV2Cluster_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeCloudhsmV2Cluster_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeCloudhsmV2Cluster_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeCloudhsmV2Cluster_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeCloudhsmV2Cluster_ClusterId(p CloudhsmV2ClusterObservation, vals map[string]cty.Value) {
	vals["cluster_id"] = cty.StringVal(p.ClusterId)
}

func EncodeCloudhsmV2Cluster_SecurityGroupId(p CloudhsmV2ClusterObservation, vals map[string]cty.Value) {
	vals["security_group_id"] = cty.StringVal(p.SecurityGroupId)
}

func EncodeCloudhsmV2Cluster_VpcId(p CloudhsmV2ClusterObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeCloudhsmV2Cluster_ClusterCertificates(p []ClusterCertificates, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeCloudhsmV2Cluster_ClusterCertificates_ManufacturerHardwareCertificate(v, ctyVal)
		EncodeCloudhsmV2Cluster_ClusterCertificates_AwsHardwareCertificate(v, ctyVal)
		EncodeCloudhsmV2Cluster_ClusterCertificates_ClusterCertificate(v, ctyVal)
		EncodeCloudhsmV2Cluster_ClusterCertificates_ClusterCsr(v, ctyVal)
		EncodeCloudhsmV2Cluster_ClusterCertificates_HsmCertificate(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cluster_certificates"] = cty.ListVal(valsForCollection)
}

func EncodeCloudhsmV2Cluster_ClusterCertificates_ManufacturerHardwareCertificate(p ClusterCertificates, vals map[string]cty.Value) {
	vals["manufacturer_hardware_certificate"] = cty.StringVal(p.ManufacturerHardwareCertificate)
}

func EncodeCloudhsmV2Cluster_ClusterCertificates_AwsHardwareCertificate(p ClusterCertificates, vals map[string]cty.Value) {
	vals["aws_hardware_certificate"] = cty.StringVal(p.AwsHardwareCertificate)
}

func EncodeCloudhsmV2Cluster_ClusterCertificates_ClusterCertificate(p ClusterCertificates, vals map[string]cty.Value) {
	vals["cluster_certificate"] = cty.StringVal(p.ClusterCertificate)
}

func EncodeCloudhsmV2Cluster_ClusterCertificates_ClusterCsr(p ClusterCertificates, vals map[string]cty.Value) {
	vals["cluster_csr"] = cty.StringVal(p.ClusterCsr)
}

func EncodeCloudhsmV2Cluster_ClusterCertificates_HsmCertificate(p ClusterCertificates, vals map[string]cty.Value) {
	vals["hsm_certificate"] = cty.StringVal(p.HsmCertificate)
}

func EncodeCloudhsmV2Cluster_ClusterState(p CloudhsmV2ClusterObservation, vals map[string]cty.Value) {
	vals["cluster_state"] = cty.StringVal(p.ClusterState)
}