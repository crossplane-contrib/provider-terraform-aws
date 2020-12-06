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
	"github.com/zclconf/go-cty/cty"
)

func EncodeCodeartifactDomainPermissionsPolicy(r CodeartifactDomainPermissionsPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodeartifactDomainPermissionsPolicy_PolicyDocument(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactDomainPermissionsPolicy_PolicyRevision(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactDomainPermissionsPolicy_Domain(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactDomainPermissionsPolicy_DomainOwner(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactDomainPermissionsPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodeartifactDomainPermissionsPolicy_ResourceArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCodeartifactDomainPermissionsPolicy_PolicyDocument(p CodeartifactDomainPermissionsPolicyParameters, vals map[string]cty.Value) {
	vals["policy_document"] = cty.StringVal(p.PolicyDocument)
}

func EncodeCodeartifactDomainPermissionsPolicy_PolicyRevision(p CodeartifactDomainPermissionsPolicyParameters, vals map[string]cty.Value) {
	vals["policy_revision"] = cty.StringVal(p.PolicyRevision)
}

func EncodeCodeartifactDomainPermissionsPolicy_Domain(p CodeartifactDomainPermissionsPolicyParameters, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeCodeartifactDomainPermissionsPolicy_DomainOwner(p CodeartifactDomainPermissionsPolicyParameters, vals map[string]cty.Value) {
	vals["domain_owner"] = cty.StringVal(p.DomainOwner)
}

func EncodeCodeartifactDomainPermissionsPolicy_Id(p CodeartifactDomainPermissionsPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodeartifactDomainPermissionsPolicy_ResourceArn(p CodeartifactDomainPermissionsPolicyObservation, vals map[string]cty.Value) {
	vals["resource_arn"] = cty.StringVal(p.ResourceArn)
}