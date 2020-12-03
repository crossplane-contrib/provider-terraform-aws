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

package v1alpha1func EncodeEfsFileSystem(r EfsFileSystem) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEfsFileSystem_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystem_PerformanceMode(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystem_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystem_ThroughputMode(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystem_CreationToken(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystem_Encrypted(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystem_Id(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystem_ProvisionedThroughputInMibps(r.Spec.ForProvider, ctyVal)
	EncodeEfsFileSystem_LifecyclePolicy(r.Spec.ForProvider.LifecyclePolicy, ctyVal)
	EncodeEfsFileSystem_Arn(r.Status.AtProvider, ctyVal)
	EncodeEfsFileSystem_DnsName(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEfsFileSystem_KmsKeyId(p *EfsFileSystemParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeEfsFileSystem_PerformanceMode(p *EfsFileSystemParameters, vals map[string]cty.Value) {
	vals["performance_mode"] = cty.StringVal(p.PerformanceMode)
}

func EncodeEfsFileSystem_Tags(p *EfsFileSystemParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEfsFileSystem_ThroughputMode(p *EfsFileSystemParameters, vals map[string]cty.Value) {
	vals["throughput_mode"] = cty.StringVal(p.ThroughputMode)
}

func EncodeEfsFileSystem_CreationToken(p *EfsFileSystemParameters, vals map[string]cty.Value) {
	vals["creation_token"] = cty.StringVal(p.CreationToken)
}

func EncodeEfsFileSystem_Encrypted(p *EfsFileSystemParameters, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeEfsFileSystem_Id(p *EfsFileSystemParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEfsFileSystem_ProvisionedThroughputInMibps(p *EfsFileSystemParameters, vals map[string]cty.Value) {
	vals["provisioned_throughput_in_mibps"] = cty.IntVal(p.ProvisionedThroughputInMibps)
}

func EncodeEfsFileSystem_LifecyclePolicy(p *LifecyclePolicy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LifecyclePolicy {
		ctyVal = make(map[string]cty.Value)
		EncodeEfsFileSystem_LifecyclePolicy_TransitionToIa(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lifecycle_policy"] = cty.ListVal(valsForCollection)
}

func EncodeEfsFileSystem_LifecyclePolicy_TransitionToIa(p *LifecyclePolicy, vals map[string]cty.Value) {
	vals["transition_to_ia"] = cty.StringVal(p.TransitionToIa)
}

func EncodeEfsFileSystem_Arn(p *EfsFileSystemObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEfsFileSystem_DnsName(p *EfsFileSystemObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}