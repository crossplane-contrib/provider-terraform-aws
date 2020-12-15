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
	r, ok := mr.(*EfsMountTarget)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EfsMountTarget.")
	}
	return EncodeEfsMountTarget(*r), nil
}

func EncodeEfsMountTarget(r EfsMountTarget) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEfsMountTarget_IpAddress(r.Spec.ForProvider, ctyVal)
	EncodeEfsMountTarget_FileSystemId(r.Spec.ForProvider, ctyVal)
	EncodeEfsMountTarget_Id(r.Spec.ForProvider, ctyVal)
	EncodeEfsMountTarget_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeEfsMountTarget_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeEfsMountTarget_NetworkInterfaceId(r.Status.AtProvider, ctyVal)
	EncodeEfsMountTarget_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeEfsMountTarget_AvailabilityZoneId(r.Status.AtProvider, ctyVal)
	EncodeEfsMountTarget_AvailabilityZoneName(r.Status.AtProvider, ctyVal)
	EncodeEfsMountTarget_FileSystemArn(r.Status.AtProvider, ctyVal)
	EncodeEfsMountTarget_DnsName(r.Status.AtProvider, ctyVal)
	EncodeEfsMountTarget_MountTargetDnsName(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeEfsMountTarget_IpAddress(p EfsMountTargetParameters, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}

func EncodeEfsMountTarget_FileSystemId(p EfsMountTargetParameters, vals map[string]cty.Value) {
	vals["file_system_id"] = cty.StringVal(p.FileSystemId)
}

func EncodeEfsMountTarget_Id(p EfsMountTargetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEfsMountTarget_SubnetId(p EfsMountTargetParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeEfsMountTarget_SecurityGroups(p EfsMountTargetParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeEfsMountTarget_NetworkInterfaceId(p EfsMountTargetObservation, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeEfsMountTarget_OwnerId(p EfsMountTargetObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeEfsMountTarget_AvailabilityZoneId(p EfsMountTargetObservation, vals map[string]cty.Value) {
	vals["availability_zone_id"] = cty.StringVal(p.AvailabilityZoneId)
}

func EncodeEfsMountTarget_AvailabilityZoneName(p EfsMountTargetObservation, vals map[string]cty.Value) {
	vals["availability_zone_name"] = cty.StringVal(p.AvailabilityZoneName)
}

func EncodeEfsMountTarget_FileSystemArn(p EfsMountTargetObservation, vals map[string]cty.Value) {
	vals["file_system_arn"] = cty.StringVal(p.FileSystemArn)
}

func EncodeEfsMountTarget_DnsName(p EfsMountTargetObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeEfsMountTarget_MountTargetDnsName(p EfsMountTargetObservation, vals map[string]cty.Value) {
	vals["mount_target_dns_name"] = cty.StringVal(p.MountTargetDnsName)
}