// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Destination struct {
	ArchivedAt             types.String                 `tfsdk:"archived_at"`
	AuthMethod             *DestinationAuthMethodConfig `tfsdk:"auth_method"`
	CliPath                types.String                 `tfsdk:"cli_path"`
	CreatedAt              types.String                 `tfsdk:"created_at"`
	HTTPMethod             types.String                 `tfsdk:"http_method"`
	ID                     types.String                 `tfsdk:"id"`
	Name                   types.String                 `tfsdk:"name"`
	PathForwardingDisabled types.Bool                   `tfsdk:"path_forwarding_disabled"`
	RateLimit              types.Int64                  `tfsdk:"rate_limit"`
	RateLimitPeriod        types.String                 `tfsdk:"rate_limit_period"`
	TeamID                 types.String                 `tfsdk:"team_id"`
	UpdatedAt              types.String                 `tfsdk:"updated_at"`
	URL                    types.String                 `tfsdk:"url"`
}
