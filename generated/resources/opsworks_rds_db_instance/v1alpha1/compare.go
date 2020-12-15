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
	k := kube.(*OpsworksRdsDbInstance)
	p := prov.(*OpsworksRdsDbInstance)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeOpsworksRdsDbInstance_RdsDbInstanceArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksRdsDbInstance_StackId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksRdsDbInstance_DbPassword(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksRdsDbInstance_DbUser(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeOpsworksRdsDbInstance_RdsDbInstanceArn(k *OpsworksRdsDbInstanceParameters, p *OpsworksRdsDbInstanceParameters, md *plugin.MergeDescription) bool {
	if k.RdsDbInstanceArn != p.RdsDbInstanceArn {
		p.RdsDbInstanceArn = k.RdsDbInstanceArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksRdsDbInstance_StackId(k *OpsworksRdsDbInstanceParameters, p *OpsworksRdsDbInstanceParameters, md *plugin.MergeDescription) bool {
	if k.StackId != p.StackId {
		p.StackId = k.StackId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksRdsDbInstance_DbPassword(k *OpsworksRdsDbInstanceParameters, p *OpsworksRdsDbInstanceParameters, md *plugin.MergeDescription) bool {
	if k.DbPassword != p.DbPassword {
		p.DbPassword = k.DbPassword
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksRdsDbInstance_DbUser(k *OpsworksRdsDbInstanceParameters, p *OpsworksRdsDbInstanceParameters, md *plugin.MergeDescription) bool {
	if k.DbUser != p.DbUser {
		p.DbUser = k.DbUser
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}