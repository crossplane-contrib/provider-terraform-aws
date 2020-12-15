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
	k := kube.(*IamServerCertificate)
	p := prov.(*IamServerCertificate)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeIamServerCertificate_CertificateChain(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServerCertificate_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServerCertificate_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServerCertificate_Path(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServerCertificate_PrivateKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServerCertificate_Arn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServerCertificate_CertificateBody(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeIamServerCertificate_CertificateChain(k *IamServerCertificateParameters, p *IamServerCertificateParameters, md *plugin.MergeDescription) bool {
	if k.CertificateChain != p.CertificateChain {
		p.CertificateChain = k.CertificateChain
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamServerCertificate_Name(k *IamServerCertificateParameters, p *IamServerCertificateParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamServerCertificate_NamePrefix(k *IamServerCertificateParameters, p *IamServerCertificateParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamServerCertificate_Path(k *IamServerCertificateParameters, p *IamServerCertificateParameters, md *plugin.MergeDescription) bool {
	if k.Path != p.Path {
		p.Path = k.Path
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamServerCertificate_PrivateKey(k *IamServerCertificateParameters, p *IamServerCertificateParameters, md *plugin.MergeDescription) bool {
	if k.PrivateKey != p.PrivateKey {
		p.PrivateKey = k.PrivateKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamServerCertificate_Arn(k *IamServerCertificateParameters, p *IamServerCertificateParameters, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		p.Arn = k.Arn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamServerCertificate_CertificateBody(k *IamServerCertificateParameters, p *IamServerCertificateParameters, md *plugin.MergeDescription) bool {
	if k.CertificateBody != p.CertificateBody {
		p.CertificateBody = k.CertificateBody
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}