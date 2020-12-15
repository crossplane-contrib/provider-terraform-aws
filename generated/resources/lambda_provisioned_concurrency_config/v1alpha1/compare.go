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
	k := kube.(*LambdaProvisionedConcurrencyConfig)
	p := prov.(*LambdaProvisionedConcurrencyConfig)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeLambdaProvisionedConcurrencyConfig_ProvisionedConcurrentExecutions(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaProvisionedConcurrencyConfig_Qualifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaProvisionedConcurrencyConfig_FunctionName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaProvisionedConcurrencyConfig_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
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
func MergeLambdaProvisionedConcurrencyConfig_ProvisionedConcurrentExecutions(k *LambdaProvisionedConcurrencyConfigParameters, p *LambdaProvisionedConcurrencyConfigParameters, md *plugin.MergeDescription) bool {
	if k.ProvisionedConcurrentExecutions != p.ProvisionedConcurrentExecutions {
		p.ProvisionedConcurrentExecutions = k.ProvisionedConcurrentExecutions
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaProvisionedConcurrencyConfig_Qualifier(k *LambdaProvisionedConcurrencyConfigParameters, p *LambdaProvisionedConcurrencyConfigParameters, md *plugin.MergeDescription) bool {
	if k.Qualifier != p.Qualifier {
		p.Qualifier = k.Qualifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaProvisionedConcurrencyConfig_FunctionName(k *LambdaProvisionedConcurrencyConfigParameters, p *LambdaProvisionedConcurrencyConfigParameters, md *plugin.MergeDescription) bool {
	if k.FunctionName != p.FunctionName {
		p.FunctionName = k.FunctionName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeLambdaProvisionedConcurrencyConfig_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeLambdaProvisionedConcurrencyConfig_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLambdaProvisionedConcurrencyConfig_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeLambdaProvisionedConcurrencyConfig_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLambdaProvisionedConcurrencyConfig_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}