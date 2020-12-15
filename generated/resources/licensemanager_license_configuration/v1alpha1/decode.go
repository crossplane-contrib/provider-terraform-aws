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
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*LicensemanagerLicenseConfiguration)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLicensemanagerLicenseConfiguration(r, ctyValue)
}

func DecodeLicensemanagerLicenseConfiguration(prev *LicensemanagerLicenseConfiguration, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLicensemanagerLicenseConfiguration_Name(&new.Spec.ForProvider, valMap)
	DecodeLicensemanagerLicenseConfiguration_Tags(&new.Spec.ForProvider, valMap)
	DecodeLicensemanagerLicenseConfiguration_Description(&new.Spec.ForProvider, valMap)
	DecodeLicensemanagerLicenseConfiguration_LicenseCount(&new.Spec.ForProvider, valMap)
	DecodeLicensemanagerLicenseConfiguration_LicenseCountHardLimit(&new.Spec.ForProvider, valMap)
	DecodeLicensemanagerLicenseConfiguration_LicenseCountingType(&new.Spec.ForProvider, valMap)
	DecodeLicensemanagerLicenseConfiguration_LicenseRules(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeLicensemanagerLicenseConfiguration_Name(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeLicensemanagerLicenseConfiguration_Tags(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeLicensemanagerLicenseConfiguration_Description(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeLicensemanagerLicenseConfiguration_LicenseCount(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	p.LicenseCount = ctwhy.ValueAsInt64(vals["license_count"])
}

//primitiveTypeDecodeTemplate
func DecodeLicensemanagerLicenseConfiguration_LicenseCountHardLimit(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	p.LicenseCountHardLimit = ctwhy.ValueAsBool(vals["license_count_hard_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeLicensemanagerLicenseConfiguration_LicenseCountingType(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	p.LicenseCountingType = ctwhy.ValueAsString(vals["license_counting_type"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeLicensemanagerLicenseConfiguration_LicenseRules(p *LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["license_rules"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.LicenseRules = goVals
}