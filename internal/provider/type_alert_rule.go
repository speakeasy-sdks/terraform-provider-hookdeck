// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type AlertRule struct {
	Strategy types.String `tfsdk:"strategy"`
	Type     types.String `tfsdk:"type"`
}
