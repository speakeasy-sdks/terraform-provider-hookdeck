// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type TransformFullTransformation struct {
	Code types.String            `tfsdk:"code"`
	Env  map[string]types.String `tfsdk:"env"`
	Name types.String            `tfsdk:"name"`
}
