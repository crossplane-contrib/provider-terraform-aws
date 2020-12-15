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
	k := kube.(*LambdaPermission)
	p := prov.(*LambdaPermission)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeLambdaPermission_Action(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaPermission_Principal(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaPermission_StatementId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaPermission_Qualifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaPermission_SourceAccount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaPermission_SourceArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaPermission_StatementIdPrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaPermission_EventSourceToken(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaPermission_FunctionName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeLambdaPermission_Action(k *LambdaPermissionParameters, p *LambdaPermissionParameters, md *plugin.MergeDescription) bool {
	if k.Action != p.Action {
		p.Action = k.Action
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaPermission_Principal(k *LambdaPermissionParameters, p *LambdaPermissionParameters, md *plugin.MergeDescription) bool {
	if k.Principal != p.Principal {
		p.Principal = k.Principal
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaPermission_StatementId(k *LambdaPermissionParameters, p *LambdaPermissionParameters, md *plugin.MergeDescription) bool {
	if k.StatementId != p.StatementId {
		p.StatementId = k.StatementId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaPermission_Qualifier(k *LambdaPermissionParameters, p *LambdaPermissionParameters, md *plugin.MergeDescription) bool {
	if k.Qualifier != p.Qualifier {
		p.Qualifier = k.Qualifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaPermission_SourceAccount(k *LambdaPermissionParameters, p *LambdaPermissionParameters, md *plugin.MergeDescription) bool {
	if k.SourceAccount != p.SourceAccount {
		p.SourceAccount = k.SourceAccount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaPermission_SourceArn(k *LambdaPermissionParameters, p *LambdaPermissionParameters, md *plugin.MergeDescription) bool {
	if k.SourceArn != p.SourceArn {
		p.SourceArn = k.SourceArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaPermission_StatementIdPrefix(k *LambdaPermissionParameters, p *LambdaPermissionParameters, md *plugin.MergeDescription) bool {
	if k.StatementIdPrefix != p.StatementIdPrefix {
		p.StatementIdPrefix = k.StatementIdPrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaPermission_EventSourceToken(k *LambdaPermissionParameters, p *LambdaPermissionParameters, md *plugin.MergeDescription) bool {
	if k.EventSourceToken != p.EventSourceToken {
		p.EventSourceToken = k.EventSourceToken
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaPermission_FunctionName(k *LambdaPermissionParameters, p *LambdaPermissionParameters, md *plugin.MergeDescription) bool {
	if k.FunctionName != p.FunctionName {
		p.FunctionName = k.FunctionName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}