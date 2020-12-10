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
	r, ok := mr.(*EcsService)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EcsService.")
	}
	return EncodeEcsService(*r), nil
}

func EncodeEcsService(r EcsService) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEcsService_EnableEcsManagedTags(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_IamRole(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_Name(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_PlatformVersion(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_Cluster(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_DesiredCount(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_Id(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_PropagateTags(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_DeploymentMaximumPercent(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_DeploymentMinimumHealthyPercent(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_ForceNewDeployment(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_HealthCheckGracePeriodSeconds(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_LaunchType(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_SchedulingStrategy(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_TaskDefinition(r.Spec.ForProvider, ctyVal)
	EncodeEcsService_OrderedPlacementStrategy(r.Spec.ForProvider.OrderedPlacementStrategy, ctyVal)
	EncodeEcsService_PlacementConstraints(r.Spec.ForProvider.PlacementConstraints, ctyVal)
	EncodeEcsService_ServiceRegistries(r.Spec.ForProvider.ServiceRegistries, ctyVal)
	EncodeEcsService_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeEcsService_CapacityProviderStrategy(r.Spec.ForProvider.CapacityProviderStrategy, ctyVal)
	EncodeEcsService_DeploymentController(r.Spec.ForProvider.DeploymentController, ctyVal)
	EncodeEcsService_LoadBalancer(r.Spec.ForProvider.LoadBalancer, ctyVal)
	EncodeEcsService_NetworkConfiguration(r.Spec.ForProvider.NetworkConfiguration, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeEcsService_EnableEcsManagedTags(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["enable_ecs_managed_tags"] = cty.BoolVal(p.EnableEcsManagedTags)
}

func EncodeEcsService_IamRole(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["iam_role"] = cty.StringVal(p.IamRole)
}

func EncodeEcsService_Name(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEcsService_PlatformVersion(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["platform_version"] = cty.StringVal(p.PlatformVersion)
}

func EncodeEcsService_Tags(p EcsServiceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEcsService_Cluster(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["cluster"] = cty.StringVal(p.Cluster)
}

func EncodeEcsService_DesiredCount(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["desired_count"] = cty.NumberIntVal(p.DesiredCount)
}

func EncodeEcsService_Id(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEcsService_PropagateTags(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["propagate_tags"] = cty.StringVal(p.PropagateTags)
}

func EncodeEcsService_DeploymentMaximumPercent(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["deployment_maximum_percent"] = cty.NumberIntVal(p.DeploymentMaximumPercent)
}

func EncodeEcsService_DeploymentMinimumHealthyPercent(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["deployment_minimum_healthy_percent"] = cty.NumberIntVal(p.DeploymentMinimumHealthyPercent)
}

func EncodeEcsService_ForceNewDeployment(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["force_new_deployment"] = cty.BoolVal(p.ForceNewDeployment)
}

func EncodeEcsService_HealthCheckGracePeriodSeconds(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["health_check_grace_period_seconds"] = cty.NumberIntVal(p.HealthCheckGracePeriodSeconds)
}

func EncodeEcsService_LaunchType(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["launch_type"] = cty.StringVal(p.LaunchType)
}

func EncodeEcsService_SchedulingStrategy(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["scheduling_strategy"] = cty.StringVal(p.SchedulingStrategy)
}

func EncodeEcsService_TaskDefinition(p EcsServiceParameters, vals map[string]cty.Value) {
	vals["task_definition"] = cty.StringVal(p.TaskDefinition)
}

func EncodeEcsService_OrderedPlacementStrategy(p []OrderedPlacementStrategy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEcsService_OrderedPlacementStrategy_Field(v, ctyVal)
		EncodeEcsService_OrderedPlacementStrategy_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ordered_placement_strategy"] = cty.ListVal(valsForCollection)
}

func EncodeEcsService_OrderedPlacementStrategy_Field(p OrderedPlacementStrategy, vals map[string]cty.Value) {
	vals["field"] = cty.StringVal(p.Field)
}

func EncodeEcsService_OrderedPlacementStrategy_Type(p OrderedPlacementStrategy, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEcsService_PlacementConstraints(p []PlacementConstraints, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEcsService_PlacementConstraints_Expression(v, ctyVal)
		EncodeEcsService_PlacementConstraints_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["placement_constraints"] = cty.SetVal(valsForCollection)
}

func EncodeEcsService_PlacementConstraints_Expression(p PlacementConstraints, vals map[string]cty.Value) {
	vals["expression"] = cty.StringVal(p.Expression)
}

func EncodeEcsService_PlacementConstraints_Type(p PlacementConstraints, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEcsService_ServiceRegistries(p ServiceRegistries, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsService_ServiceRegistries_ContainerName(p, ctyVal)
	EncodeEcsService_ServiceRegistries_ContainerPort(p, ctyVal)
	EncodeEcsService_ServiceRegistries_Port(p, ctyVal)
	EncodeEcsService_ServiceRegistries_RegistryArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["service_registries"] = cty.ListVal(valsForCollection)
}

func EncodeEcsService_ServiceRegistries_ContainerName(p ServiceRegistries, vals map[string]cty.Value) {
	vals["container_name"] = cty.StringVal(p.ContainerName)
}

func EncodeEcsService_ServiceRegistries_ContainerPort(p ServiceRegistries, vals map[string]cty.Value) {
	vals["container_port"] = cty.NumberIntVal(p.ContainerPort)
}

func EncodeEcsService_ServiceRegistries_Port(p ServiceRegistries, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeEcsService_ServiceRegistries_RegistryArn(p ServiceRegistries, vals map[string]cty.Value) {
	vals["registry_arn"] = cty.StringVal(p.RegistryArn)
}

func EncodeEcsService_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeEcsService_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeEcsService_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeEcsService_CapacityProviderStrategy(p CapacityProviderStrategy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsService_CapacityProviderStrategy_Base(p, ctyVal)
	EncodeEcsService_CapacityProviderStrategy_CapacityProvider(p, ctyVal)
	EncodeEcsService_CapacityProviderStrategy_Weight(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["capacity_provider_strategy"] = cty.SetVal(valsForCollection)
}

func EncodeEcsService_CapacityProviderStrategy_Base(p CapacityProviderStrategy, vals map[string]cty.Value) {
	vals["base"] = cty.NumberIntVal(p.Base)
}

func EncodeEcsService_CapacityProviderStrategy_CapacityProvider(p CapacityProviderStrategy, vals map[string]cty.Value) {
	vals["capacity_provider"] = cty.StringVal(p.CapacityProvider)
}

func EncodeEcsService_CapacityProviderStrategy_Weight(p CapacityProviderStrategy, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeEcsService_DeploymentController(p DeploymentController, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsService_DeploymentController_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["deployment_controller"] = cty.ListVal(valsForCollection)
}

func EncodeEcsService_DeploymentController_Type(p DeploymentController, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEcsService_LoadBalancer(p LoadBalancer, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsService_LoadBalancer_ElbName(p, ctyVal)
	EncodeEcsService_LoadBalancer_TargetGroupArn(p, ctyVal)
	EncodeEcsService_LoadBalancer_ContainerName(p, ctyVal)
	EncodeEcsService_LoadBalancer_ContainerPort(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["load_balancer"] = cty.SetVal(valsForCollection)
}

func EncodeEcsService_LoadBalancer_ElbName(p LoadBalancer, vals map[string]cty.Value) {
	vals["elb_name"] = cty.StringVal(p.ElbName)
}

func EncodeEcsService_LoadBalancer_TargetGroupArn(p LoadBalancer, vals map[string]cty.Value) {
	vals["target_group_arn"] = cty.StringVal(p.TargetGroupArn)
}

func EncodeEcsService_LoadBalancer_ContainerName(p LoadBalancer, vals map[string]cty.Value) {
	vals["container_name"] = cty.StringVal(p.ContainerName)
}

func EncodeEcsService_LoadBalancer_ContainerPort(p LoadBalancer, vals map[string]cty.Value) {
	vals["container_port"] = cty.NumberIntVal(p.ContainerPort)
}

func EncodeEcsService_NetworkConfiguration(p NetworkConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEcsService_NetworkConfiguration_AssignPublicIp(p, ctyVal)
	EncodeEcsService_NetworkConfiguration_SecurityGroups(p, ctyVal)
	EncodeEcsService_NetworkConfiguration_Subnets(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["network_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeEcsService_NetworkConfiguration_AssignPublicIp(p NetworkConfiguration, vals map[string]cty.Value) {
	vals["assign_public_ip"] = cty.BoolVal(p.AssignPublicIp)
}

func EncodeEcsService_NetworkConfiguration_SecurityGroups(p NetworkConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeEcsService_NetworkConfiguration_Subnets(p NetworkConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Subnets {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnets"] = cty.SetVal(colVals)
}