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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*Apigatewayv2Deployment)
	p := prov.(*Apigatewayv2Deployment)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApigatewayv2Deployment_ApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Deployment_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Deployment_Triggers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Deployment_AutoDeployed(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Deployment_ApiId(k *Apigatewayv2DeploymentParameters, p *Apigatewayv2DeploymentParameters, md *plugin.MergeDescription) bool {
	if k.ApiId != p.ApiId {
		p.ApiId = k.ApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Deployment_Description(k *Apigatewayv2DeploymentParameters, p *Apigatewayv2DeploymentParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApigatewayv2Deployment_Triggers(k *Apigatewayv2DeploymentParameters, p *Apigatewayv2DeploymentParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Triggers, p.Triggers) {
		p.Triggers = k.Triggers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApigatewayv2Deployment_AutoDeployed(k *Apigatewayv2DeploymentObservation, p *Apigatewayv2DeploymentObservation, md *plugin.MergeDescription) bool {
	if k.AutoDeployed != p.AutoDeployed {
		k.AutoDeployed = p.AutoDeployed
		md.StatusUpdated = true
		return true
	}
	return false
}