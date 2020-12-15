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
	k := kube.(*ApiGatewayDeployment)
	p := prov.(*ApiGatewayDeployment)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApiGatewayDeployment_RestApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayDeployment_Triggers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayDeployment_Variables(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayDeployment_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayDeployment_StageDescription(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayDeployment_StageName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayDeployment_InvokeUrl(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayDeployment_CreatedDate(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayDeployment_ExecutionArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeApiGatewayDeployment_RestApiId(k *ApiGatewayDeploymentParameters, p *ApiGatewayDeploymentParameters, md *plugin.MergeDescription) bool {
	if k.RestApiId != p.RestApiId {
		p.RestApiId = k.RestApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayDeployment_Triggers(k *ApiGatewayDeploymentParameters, p *ApiGatewayDeploymentParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Triggers, p.Triggers) {
		p.Triggers = k.Triggers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayDeployment_Variables(k *ApiGatewayDeploymentParameters, p *ApiGatewayDeploymentParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Variables, p.Variables) {
		p.Variables = k.Variables
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayDeployment_Description(k *ApiGatewayDeploymentParameters, p *ApiGatewayDeploymentParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayDeployment_StageDescription(k *ApiGatewayDeploymentParameters, p *ApiGatewayDeploymentParameters, md *plugin.MergeDescription) bool {
	if k.StageDescription != p.StageDescription {
		p.StageDescription = k.StageDescription
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayDeployment_StageName(k *ApiGatewayDeploymentParameters, p *ApiGatewayDeploymentParameters, md *plugin.MergeDescription) bool {
	if k.StageName != p.StageName {
		p.StageName = k.StageName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApiGatewayDeployment_InvokeUrl(k *ApiGatewayDeploymentObservation, p *ApiGatewayDeploymentObservation, md *plugin.MergeDescription) bool {
	if k.InvokeUrl != p.InvokeUrl {
		k.InvokeUrl = p.InvokeUrl
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApiGatewayDeployment_CreatedDate(k *ApiGatewayDeploymentObservation, p *ApiGatewayDeploymentObservation, md *plugin.MergeDescription) bool {
	if k.CreatedDate != p.CreatedDate {
		k.CreatedDate = p.CreatedDate
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApiGatewayDeployment_ExecutionArn(k *ApiGatewayDeploymentObservation, p *ApiGatewayDeploymentObservation, md *plugin.MergeDescription) bool {
	if k.ExecutionArn != p.ExecutionArn {
		k.ExecutionArn = p.ExecutionArn
		md.StatusUpdated = true
		return true
	}
	return false
}