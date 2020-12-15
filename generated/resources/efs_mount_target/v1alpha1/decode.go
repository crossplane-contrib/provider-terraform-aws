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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*EfsMountTarget)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEfsMountTarget(r, ctyValue)
}

func DecodeEfsMountTarget(prev *EfsMountTarget, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEfsMountTarget_IpAddress(&new.Spec.ForProvider, valMap)
	DecodeEfsMountTarget_FileSystemId(&new.Spec.ForProvider, valMap)
	DecodeEfsMountTarget_Id(&new.Spec.ForProvider, valMap)
	DecodeEfsMountTarget_SubnetId(&new.Spec.ForProvider, valMap)
	DecodeEfsMountTarget_SecurityGroups(&new.Spec.ForProvider, valMap)
	DecodeEfsMountTarget_NetworkInterfaceId(&new.Status.AtProvider, valMap)
	DecodeEfsMountTarget_OwnerId(&new.Status.AtProvider, valMap)
	DecodeEfsMountTarget_AvailabilityZoneId(&new.Status.AtProvider, valMap)
	DecodeEfsMountTarget_AvailabilityZoneName(&new.Status.AtProvider, valMap)
	DecodeEfsMountTarget_FileSystemArn(&new.Status.AtProvider, valMap)
	DecodeEfsMountTarget_DnsName(&new.Status.AtProvider, valMap)
	DecodeEfsMountTarget_MountTargetDnsName(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_IpAddress(p *EfsMountTargetParameters, vals map[string]cty.Value) {
	p.IpAddress = ctwhy.ValueAsString(vals["ip_address"])
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_FileSystemId(p *EfsMountTargetParameters, vals map[string]cty.Value) {
	p.FileSystemId = ctwhy.ValueAsString(vals["file_system_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_Id(p *EfsMountTargetParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_SubnetId(p *EfsMountTargetParameters, vals map[string]cty.Value) {
	p.SubnetId = ctwhy.ValueAsString(vals["subnet_id"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeEfsMountTarget_SecurityGroups(p *EfsMountTargetParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["security_groups"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SecurityGroups = goVals
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_NetworkInterfaceId(p *EfsMountTargetObservation, vals map[string]cty.Value) {
	p.NetworkInterfaceId = ctwhy.ValueAsString(vals["network_interface_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_OwnerId(p *EfsMountTargetObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_AvailabilityZoneId(p *EfsMountTargetObservation, vals map[string]cty.Value) {
	p.AvailabilityZoneId = ctwhy.ValueAsString(vals["availability_zone_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_AvailabilityZoneName(p *EfsMountTargetObservation, vals map[string]cty.Value) {
	p.AvailabilityZoneName = ctwhy.ValueAsString(vals["availability_zone_name"])
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_FileSystemArn(p *EfsMountTargetObservation, vals map[string]cty.Value) {
	p.FileSystemArn = ctwhy.ValueAsString(vals["file_system_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_DnsName(p *EfsMountTargetObservation, vals map[string]cty.Value) {
	p.DnsName = ctwhy.ValueAsString(vals["dns_name"])
}

//primitiveTypeDecodeTemplate
func DecodeEfsMountTarget_MountTargetDnsName(p *EfsMountTargetObservation, vals map[string]cty.Value) {
	p.MountTargetDnsName = ctwhy.ValueAsString(vals["mount_target_dns_name"])
}