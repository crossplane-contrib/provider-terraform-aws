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

package v1alpha1func EncodeWafv2RuleGroup(r Wafv2RuleGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafv2RuleGroup_Scope(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RuleGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RuleGroup_Capacity(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RuleGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RuleGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RuleGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RuleGroup_Rule(r.Spec.ForProvider.Rule, ctyVal)
	EncodeWafv2RuleGroup_VisibilityConfig(r.Spec.ForProvider.VisibilityConfig, ctyVal)
	EncodeWafv2RuleGroup_Arn(r.Status.AtProvider, ctyVal)
	EncodeWafv2RuleGroup_LockToken(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeWafv2RuleGroup_Scope(p *Wafv2RuleGroupParameters, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeWafv2RuleGroup_Tags(p *Wafv2RuleGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWafv2RuleGroup_Capacity(p *Wafv2RuleGroupParameters, vals map[string]cty.Value) {
	vals["capacity"] = cty.IntVal(p.Capacity)
}

func EncodeWafv2RuleGroup_Description(p *Wafv2RuleGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeWafv2RuleGroup_Id(p *Wafv2RuleGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafv2RuleGroup_Name(p *Wafv2RuleGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule(p *Rule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Rule {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Name(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Action(v.Action, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement(v.Statement, ctyVal)
		EncodeWafv2RuleGroup_Rule_VisibilityConfig(v.VisibilityConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["rule"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Name(p *Rule, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Priority(p *Rule, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Action(p *Action, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Action {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Action_Allow(v.Allow, ctyVal)
		EncodeWafv2RuleGroup_Rule_Action_Block(v.Block, ctyVal)
		EncodeWafv2RuleGroup_Rule_Action_Count(v.Count, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Action_Allow(p *Allow, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Allow {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["allow"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Action_Block(p *Block, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Block {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["block"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Action_Count(p *Count, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Count {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["count"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement(v.NotStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement(v.OrStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement(v.AndStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement(p *NotStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NotStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["not_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement(v.NotStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement(v.OrStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement(v.AndStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement(p *NotStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NotStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["not_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement(p *OrStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OrStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["or_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement(p *AndStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AndStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["and_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement(p *OrStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OrStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["or_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement(v.AndStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement(v.NotStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement(v.OrStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement(p *AndStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AndStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["and_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement(p *NotStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NotStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["not_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement(p *OrStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OrStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["or_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement(p *AndStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AndStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["and_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement(v.OrStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement(v.AndStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement(v.NotStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement(p *OrStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OrStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["or_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_OrStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement(p *AndStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AndStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["and_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_AndStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement(p *NotStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NotStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement(v.Statement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["not_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement(p *Statement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Statement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement(v.SqliMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement(v.XssMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement(v.ByteMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement(v.GeoMatchStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement(v.IpSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement(v.RegexPatternSetReferenceStatement, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement(v.SizeConstraintStatement, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement(p *SqliMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqliMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sqli_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SqliMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement(p *GeoMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeoMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement_CountryCodes(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig(v.ForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geo_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement_CountryCodes(p *GeoMatchStatement, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CountryCodes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["country_codes"] = cty.ListVal(colVals)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig(p *ForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_FallbackBehavior(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_GeoMatchStatement_ForwardedIpConfig_HeaderName(p *ForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(v.IpSetForwardedIpConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_Arn(p *IpSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.IpSetForwardedIpConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_set_forwarded_ip_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_Position(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["position"] = cty.StringVal(p.Position)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_FallbackBehavior(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["fallback_behavior"] = cty.StringVal(p.FallbackBehavior)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_IpSetReferenceStatement_IpSetForwardedIpConfig_HeaderName(p *IpSetForwardedIpConfig, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexPatternSetReferenceStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_Arn(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_pattern_set_reference_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_Arn(p *RegexPatternSetReferenceStatement, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_RegexPatternSetReferenceStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement(p *SizeConstraintStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraintStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_ComparisonOperator(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_Size(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraint_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_ComparisonOperator(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_Size(p *SizeConstraintStatement, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_NotStatement_Statement_SizeConstraintStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement(p *XssMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.XssMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["xss_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_AndStatement_Statement_XssMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement(p *ByteMatchStatement, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchStatement {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_PositionalConstraint(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_SearchString(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch(v.FieldToMatch, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_TextTransformation(v.TextTransformation, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_statement"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_PositionalConstraint(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_SearchString(p *ByteMatchStatement, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(v.AllQueryArguments, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_Body(v.Body, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_Method(v.Method, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_QueryString(v.QueryString, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(v.SingleHeader, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(v.SingleQueryArgument, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_UriPath(v.UriPath, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_AllQueryArguments(p *AllQueryArguments, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AllQueryArguments {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["all_query_arguments"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_Body(p *Body, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Body {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["body"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_Method(p *Method, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Method {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["method"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_QueryString(p *QueryString, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QueryString {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["query_string"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_SingleHeader(p *SingleHeader, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleHeader {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_header"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_SingleHeader_Name(p *SingleHeader, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument(p *SingleQueryArgument, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SingleQueryArgument {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["single_query_argument"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_SingleQueryArgument_Name(p *SingleQueryArgument, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_FieldToMatch_UriPath(p *UriPath, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UriPath {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["uri_path"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_TextTransformation(p *TextTransformation, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TextTransformation {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_TextTransformation_Priority(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_TextTransformation_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["text_transformation"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_TextTransformation_Priority(p *TextTransformation, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafv2RuleGroup_Rule_Statement_ByteMatchStatement_TextTransformation_Type(p *TextTransformation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafv2RuleGroup_Rule_VisibilityConfig(p *VisibilityConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VisibilityConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_Rule_VisibilityConfig_CloudwatchMetricsEnabled(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_VisibilityConfig_MetricName(v, ctyVal)
		EncodeWafv2RuleGroup_Rule_VisibilityConfig_SampledRequestsEnabled(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["visibility_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_Rule_VisibilityConfig_CloudwatchMetricsEnabled(p *VisibilityConfig, vals map[string]cty.Value) {
	vals["cloudwatch_metrics_enabled"] = cty.BoolVal(p.CloudwatchMetricsEnabled)
}

func EncodeWafv2RuleGroup_Rule_VisibilityConfig_MetricName(p *VisibilityConfig, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeWafv2RuleGroup_Rule_VisibilityConfig_SampledRequestsEnabled(p *VisibilityConfig, vals map[string]cty.Value) {
	vals["sampled_requests_enabled"] = cty.BoolVal(p.SampledRequestsEnabled)
}

func EncodeWafv2RuleGroup_VisibilityConfig(p *VisibilityConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VisibilityConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RuleGroup_VisibilityConfig_CloudwatchMetricsEnabled(v, ctyVal)
		EncodeWafv2RuleGroup_VisibilityConfig_MetricName(v, ctyVal)
		EncodeWafv2RuleGroup_VisibilityConfig_SampledRequestsEnabled(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["visibility_config"] = cty.ListVal(valsForCollection)
}

func EncodeWafv2RuleGroup_VisibilityConfig_CloudwatchMetricsEnabled(p *VisibilityConfig, vals map[string]cty.Value) {
	vals["cloudwatch_metrics_enabled"] = cty.BoolVal(p.CloudwatchMetricsEnabled)
}

func EncodeWafv2RuleGroup_VisibilityConfig_MetricName(p *VisibilityConfig, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeWafv2RuleGroup_VisibilityConfig_SampledRequestsEnabled(p *VisibilityConfig, vals map[string]cty.Value) {
	vals["sampled_requests_enabled"] = cty.BoolVal(p.SampledRequestsEnabled)
}

func EncodeWafv2RuleGroup_Arn(p *Wafv2RuleGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeWafv2RuleGroup_LockToken(p *Wafv2RuleGroupObservation, vals map[string]cty.Value) {
	vals["lock_token"] = cty.StringVal(p.LockToken)
}