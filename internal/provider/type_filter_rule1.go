// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type FilterRule1 struct {
	Body    *ConnectionFilterProperty5 `tfsdk:"body"`
	Headers *ConnectionFilterProperty6 `tfsdk:"headers"`
	Path    *ConnectionFilterProperty7 `tfsdk:"path"`
	Query   *ConnectionFilterProperty8 `tfsdk:"query"`
	Type    types.String               `tfsdk:"type"`
}