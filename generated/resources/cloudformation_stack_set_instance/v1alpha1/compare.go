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
	k := kube.(*CloudformationStackSetInstance)
	p := prov.(*CloudformationStackSetInstance)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCloudformationStackSetInstance_RetainStack(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSetInstance_StackSetName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSetInstance_AccountId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSetInstance_ParameterOverrides(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSetInstance_Region(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSetInstance_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSetInstance_StackId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCloudformationStackSetInstance_RetainStack(k *CloudformationStackSetInstanceParameters, p *CloudformationStackSetInstanceParameters, md *plugin.MergeDescription) bool {
	if k.RetainStack != p.RetainStack {
		p.RetainStack = k.RetainStack
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSetInstance_StackSetName(k *CloudformationStackSetInstanceParameters, p *CloudformationStackSetInstanceParameters, md *plugin.MergeDescription) bool {
	if k.StackSetName != p.StackSetName {
		p.StackSetName = k.StackSetName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSetInstance_AccountId(k *CloudformationStackSetInstanceParameters, p *CloudformationStackSetInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AccountId != p.AccountId {
		p.AccountId = k.AccountId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCloudformationStackSetInstance_ParameterOverrides(k *CloudformationStackSetInstanceParameters, p *CloudformationStackSetInstanceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.ParameterOverrides, p.ParameterOverrides) {
		p.ParameterOverrides = k.ParameterOverrides
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSetInstance_Region(k *CloudformationStackSetInstanceParameters, p *CloudformationStackSetInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Region != p.Region {
		p.Region = k.Region
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeCloudformationStackSetInstance_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeCloudformationStackSetInstance_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSetInstance_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSetInstance_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSetInstance_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSetInstance_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSetInstance_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudformationStackSetInstance_StackId(k *CloudformationStackSetInstanceObservation, p *CloudformationStackSetInstanceObservation, md *plugin.MergeDescription) bool {
	if k.StackId != p.StackId {
		k.StackId = p.StackId
		md.StatusUpdated = true
		return true
	}
	return false
}