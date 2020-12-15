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
	k := kube.(*BackupVaultNotifications)
	p := prov.(*BackupVaultNotifications)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeBackupVaultNotifications_BackupVaultEvents(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeBackupVaultNotifications_BackupVaultName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeBackupVaultNotifications_SnsTopicArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeBackupVaultNotifications_BackupVaultArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeBackupVaultNotifications_BackupVaultEvents(k *BackupVaultNotificationsParameters, p *BackupVaultNotificationsParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.BackupVaultEvents, p.BackupVaultEvents) {
		p.BackupVaultEvents = k.BackupVaultEvents
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeBackupVaultNotifications_BackupVaultName(k *BackupVaultNotificationsParameters, p *BackupVaultNotificationsParameters, md *plugin.MergeDescription) bool {
	if k.BackupVaultName != p.BackupVaultName {
		p.BackupVaultName = k.BackupVaultName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeBackupVaultNotifications_SnsTopicArn(k *BackupVaultNotificationsParameters, p *BackupVaultNotificationsParameters, md *plugin.MergeDescription) bool {
	if k.SnsTopicArn != p.SnsTopicArn {
		p.SnsTopicArn = k.SnsTopicArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeBackupVaultNotifications_BackupVaultArn(k *BackupVaultNotificationsObservation, p *BackupVaultNotificationsObservation, md *plugin.MergeDescription) bool {
	if k.BackupVaultArn != p.BackupVaultArn {
		k.BackupVaultArn = p.BackupVaultArn
		md.StatusUpdated = true
		return true
	}
	return false
}