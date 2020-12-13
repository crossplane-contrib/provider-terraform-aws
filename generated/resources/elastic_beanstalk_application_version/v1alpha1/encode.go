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

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*ElasticBeanstalkApplicationVersion)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ElasticBeanstalkApplicationVersion.")
	}
	return EncodeElasticBeanstalkApplicationVersion(*r), nil
}

func EncodeElasticBeanstalkApplicationVersion(r ElasticBeanstalkApplicationVersion) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeElasticBeanstalkApplicationVersion_Tags(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Description(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Key(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Name(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Application(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_ForceDelete(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeElasticBeanstalkApplicationVersion_Tags(p ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeElasticBeanstalkApplicationVersion_Bucket(p ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeElasticBeanstalkApplicationVersion_Description(p ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeElasticBeanstalkApplicationVersion_Key(p ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeElasticBeanstalkApplicationVersion_Id(p ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticBeanstalkApplicationVersion_Name(p ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticBeanstalkApplicationVersion_Application(p ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["application"] = cty.StringVal(p.Application)
}

func EncodeElasticBeanstalkApplicationVersion_ForceDelete(p ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["force_delete"] = cty.BoolVal(p.ForceDelete)
}

func EncodeElasticBeanstalkApplicationVersion_Arn(p ElasticBeanstalkApplicationVersionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}