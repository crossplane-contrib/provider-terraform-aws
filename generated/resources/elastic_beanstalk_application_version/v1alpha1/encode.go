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

package v1alpha1func EncodeElasticBeanstalkApplicationVersion(r ElasticBeanstalkApplicationVersion) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeElasticBeanstalkApplicationVersion_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Key(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Application(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Description(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_ForceDelete(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Name(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Tags(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplicationVersion_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeElasticBeanstalkApplicationVersion_Bucket(p *ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeElasticBeanstalkApplicationVersion_Id(p *ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticBeanstalkApplicationVersion_Key(p *ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeElasticBeanstalkApplicationVersion_Application(p *ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["application"] = cty.StringVal(p.Application)
}

func EncodeElasticBeanstalkApplicationVersion_Description(p *ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeElasticBeanstalkApplicationVersion_ForceDelete(p *ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["force_delete"] = cty.BoolVal(p.ForceDelete)
}

func EncodeElasticBeanstalkApplicationVersion_Name(p *ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticBeanstalkApplicationVersion_Tags(p *ElasticBeanstalkApplicationVersionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeElasticBeanstalkApplicationVersion_Arn(p *ElasticBeanstalkApplicationVersionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}