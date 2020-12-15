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
	r, ok := mr.(*Ec2CapacityReservation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2CapacityReservation(r, ctyValue)
}

func DecodeEc2CapacityReservation(prev *Ec2CapacityReservation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2CapacityReservation_InstanceType(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_Tags(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_InstanceCount(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_InstanceMatchCriteria(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_InstancePlatform(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_Tenancy(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_EbsOptimized(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_EndDate(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_EndDateType(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_EphemeralStorage(&new.Spec.ForProvider, valMap)
	DecodeEc2CapacityReservation_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_InstanceType(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.InstanceType = ctwhy.ValueAsString(vals["instance_type"])
}

//primitiveMapTypeDecodeTemplate
func DecodeEc2CapacityReservation_Tags(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_AvailabilityZone(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_InstanceCount(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.InstanceCount = ctwhy.ValueAsInt64(vals["instance_count"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_InstanceMatchCriteria(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.InstanceMatchCriteria = ctwhy.ValueAsString(vals["instance_match_criteria"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_InstancePlatform(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.InstancePlatform = ctwhy.ValueAsString(vals["instance_platform"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_Tenancy(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.Tenancy = ctwhy.ValueAsString(vals["tenancy"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_EbsOptimized(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.EbsOptimized = ctwhy.ValueAsBool(vals["ebs_optimized"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_EndDate(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.EndDate = ctwhy.ValueAsString(vals["end_date"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_EndDateType(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.EndDateType = ctwhy.ValueAsString(vals["end_date_type"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_EphemeralStorage(p *Ec2CapacityReservationParameters, vals map[string]cty.Value) {
	p.EphemeralStorage = ctwhy.ValueAsBool(vals["ephemeral_storage"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2CapacityReservation_Arn(p *Ec2CapacityReservationObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}