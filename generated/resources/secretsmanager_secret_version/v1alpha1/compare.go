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
	k := kube.(*SecretsmanagerSecretVersion)
	p := prov.(*SecretsmanagerSecretVersion)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSecretsmanagerSecretVersion_SecretBinary(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecretsmanagerSecretVersion_SecretId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecretsmanagerSecretVersion_SecretString(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecretsmanagerSecretVersion_VersionStages(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecretsmanagerSecretVersion_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecretsmanagerSecretVersion_VersionId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeSecretsmanagerSecretVersion_SecretBinary(k *SecretsmanagerSecretVersionParameters, p *SecretsmanagerSecretVersionParameters, md *plugin.MergeDescription) bool {
	if k.SecretBinary != p.SecretBinary {
		p.SecretBinary = k.SecretBinary
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSecretsmanagerSecretVersion_SecretId(k *SecretsmanagerSecretVersionParameters, p *SecretsmanagerSecretVersionParameters, md *plugin.MergeDescription) bool {
	if k.SecretId != p.SecretId {
		p.SecretId = k.SecretId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSecretsmanagerSecretVersion_SecretString(k *SecretsmanagerSecretVersionParameters, p *SecretsmanagerSecretVersionParameters, md *plugin.MergeDescription) bool {
	if k.SecretString != p.SecretString {
		p.SecretString = k.SecretString
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeSecretsmanagerSecretVersion_VersionStages(k *SecretsmanagerSecretVersionParameters, p *SecretsmanagerSecretVersionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.VersionStages, p.VersionStages) {
		p.VersionStages = k.VersionStages
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSecretsmanagerSecretVersion_Arn(k *SecretsmanagerSecretVersionObservation, p *SecretsmanagerSecretVersionObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSecretsmanagerSecretVersion_VersionId(k *SecretsmanagerSecretVersionObservation, p *SecretsmanagerSecretVersionObservation, md *plugin.MergeDescription) bool {
	if k.VersionId != p.VersionId {
		k.VersionId = p.VersionId
		md.StatusUpdated = true
		return true
	}
	return false
}