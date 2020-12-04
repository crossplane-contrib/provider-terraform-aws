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

func EncodeElb(r Elb) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeElb_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeElb_Subnets(r.Spec.ForProvider, ctyVal)
	EncodeElb_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeElb_SourceSecurityGroup(r.Spec.ForProvider, ctyVal)
	EncodeElb_Tags(r.Spec.ForProvider, ctyVal)
	EncodeElb_AvailabilityZones(r.Spec.ForProvider, ctyVal)
	EncodeElb_ConnectionDraining(r.Spec.ForProvider, ctyVal)
	EncodeElb_Name(r.Spec.ForProvider, ctyVal)
	EncodeElb_Internal(r.Spec.ForProvider, ctyVal)
	EncodeElb_ConnectionDrainingTimeout(r.Spec.ForProvider, ctyVal)
	EncodeElb_CrossZoneLoadBalancing(r.Spec.ForProvider, ctyVal)
	EncodeElb_Id(r.Spec.ForProvider, ctyVal)
	EncodeElb_IdleTimeout(r.Spec.ForProvider, ctyVal)
	EncodeElb_Instances(r.Spec.ForProvider, ctyVal)
	EncodeElb_AccessLogs(r.Spec.ForProvider.AccessLogs, ctyVal)
	EncodeElb_HealthCheck(r.Spec.ForProvider.HealthCheck, ctyVal)
	EncodeElb_Listener(r.Spec.ForProvider.Listener, ctyVal)
	EncodeElb_ZoneId(r.Status.AtProvider, ctyVal)
	EncodeElb_DnsName(r.Status.AtProvider, ctyVal)
	EncodeElb_Arn(r.Status.AtProvider, ctyVal)
	EncodeElb_SourceSecurityGroupId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeElb_NamePrefix(p ElbParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeElb_Subnets(p ElbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Subnets {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnets"] = cty.SetVal(colVals)
}

func EncodeElb_SecurityGroups(p ElbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeElb_SourceSecurityGroup(p ElbParameters, vals map[string]cty.Value) {
	vals["source_security_group"] = cty.StringVal(p.SourceSecurityGroup)
}

func EncodeElb_Tags(p ElbParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeElb_AvailabilityZones(p ElbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeElb_ConnectionDraining(p ElbParameters, vals map[string]cty.Value) {
	vals["connection_draining"] = cty.BoolVal(p.ConnectionDraining)
}

func EncodeElb_Name(p ElbParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElb_Internal(p ElbParameters, vals map[string]cty.Value) {
	vals["internal"] = cty.BoolVal(p.Internal)
}

func EncodeElb_ConnectionDrainingTimeout(p ElbParameters, vals map[string]cty.Value) {
	vals["connection_draining_timeout"] = cty.NumberIntVal(p.ConnectionDrainingTimeout)
}

func EncodeElb_CrossZoneLoadBalancing(p ElbParameters, vals map[string]cty.Value) {
	vals["cross_zone_load_balancing"] = cty.BoolVal(p.CrossZoneLoadBalancing)
}

func EncodeElb_Id(p ElbParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElb_IdleTimeout(p ElbParameters, vals map[string]cty.Value) {
	vals["idle_timeout"] = cty.NumberIntVal(p.IdleTimeout)
}

func EncodeElb_Instances(p ElbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Instances {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["instances"] = cty.SetVal(colVals)
}

func EncodeElb_AccessLogs(p AccessLogs, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeElb_AccessLogs_Bucket(p, ctyVal)
	EncodeElb_AccessLogs_BucketPrefix(p, ctyVal)
	EncodeElb_AccessLogs_Enabled(p, ctyVal)
	EncodeElb_AccessLogs_Interval(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["access_logs"] = cty.ListVal(valsForCollection)
}

func EncodeElb_AccessLogs_Bucket(p AccessLogs, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeElb_AccessLogs_BucketPrefix(p AccessLogs, vals map[string]cty.Value) {
	vals["bucket_prefix"] = cty.StringVal(p.BucketPrefix)
}

func EncodeElb_AccessLogs_Enabled(p AccessLogs, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeElb_AccessLogs_Interval(p AccessLogs, vals map[string]cty.Value) {
	vals["interval"] = cty.NumberIntVal(p.Interval)
}

func EncodeElb_HealthCheck(p HealthCheck, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeElb_HealthCheck_HealthyThreshold(p, ctyVal)
	EncodeElb_HealthCheck_Interval(p, ctyVal)
	EncodeElb_HealthCheck_Target(p, ctyVal)
	EncodeElb_HealthCheck_Timeout(p, ctyVal)
	EncodeElb_HealthCheck_UnhealthyThreshold(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["health_check"] = cty.ListVal(valsForCollection)
}

func EncodeElb_HealthCheck_HealthyThreshold(p HealthCheck, vals map[string]cty.Value) {
	vals["healthy_threshold"] = cty.NumberIntVal(p.HealthyThreshold)
}

func EncodeElb_HealthCheck_Interval(p HealthCheck, vals map[string]cty.Value) {
	vals["interval"] = cty.NumberIntVal(p.Interval)
}

func EncodeElb_HealthCheck_Target(p HealthCheck, vals map[string]cty.Value) {
	vals["target"] = cty.StringVal(p.Target)
}

func EncodeElb_HealthCheck_Timeout(p HealthCheck, vals map[string]cty.Value) {
	vals["timeout"] = cty.NumberIntVal(p.Timeout)
}

func EncodeElb_HealthCheck_UnhealthyThreshold(p HealthCheck, vals map[string]cty.Value) {
	vals["unhealthy_threshold"] = cty.NumberIntVal(p.UnhealthyThreshold)
}

func EncodeElb_Listener(p []Listener, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeElb_Listener_InstancePort(v, ctyVal)
		EncodeElb_Listener_InstanceProtocol(v, ctyVal)
		EncodeElb_Listener_LbPort(v, ctyVal)
		EncodeElb_Listener_LbProtocol(v, ctyVal)
		EncodeElb_Listener_SslCertificateId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["listener"] = cty.SetVal(valsForCollection)
}

func EncodeElb_Listener_InstancePort(p Listener, vals map[string]cty.Value) {
	vals["instance_port"] = cty.NumberIntVal(p.InstancePort)
}

func EncodeElb_Listener_InstanceProtocol(p Listener, vals map[string]cty.Value) {
	vals["instance_protocol"] = cty.StringVal(p.InstanceProtocol)
}

func EncodeElb_Listener_LbPort(p Listener, vals map[string]cty.Value) {
	vals["lb_port"] = cty.NumberIntVal(p.LbPort)
}

func EncodeElb_Listener_LbProtocol(p Listener, vals map[string]cty.Value) {
	vals["lb_protocol"] = cty.StringVal(p.LbProtocol)
}

func EncodeElb_Listener_SslCertificateId(p Listener, vals map[string]cty.Value) {
	vals["ssl_certificate_id"] = cty.StringVal(p.SslCertificateId)
}

func EncodeElb_ZoneId(p ElbObservation, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}

func EncodeElb_DnsName(p ElbObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeElb_Arn(p ElbObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeElb_SourceSecurityGroupId(p ElbObservation, vals map[string]cty.Value) {
	vals["source_security_group_id"] = cty.StringVal(p.SourceSecurityGroupId)
}