// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceIntegration struct {
	Features []types.String `tfsdk:"features"`
	ID       types.String   `tfsdk:"id"`
	Label    types.String   `tfsdk:"label"`
	Provider types.String   `tfsdk:"provider"`
}
