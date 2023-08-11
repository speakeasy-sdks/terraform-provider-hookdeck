// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"hashicups/internal/sdk"
	"hashicups/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &HashicupsProvider{}

type HashicupsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// HashicupsProviderModel describes the provider data model.
type HashicupsProviderModel struct {
	ServerURL  types.String `tfsdk:"server_url"`
	Password   types.String `tfsdk:"password"`
	Username   types.String `tfsdk:"username"`
	BearerAuth types.String `tfsdk:"bearer_auth"`
}

func (p *HashicupsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hashicups"
	resp.Version = p.version
}

func (p *HashicupsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://api.hookdeck.com/{version})",
				Optional:            true,
				Required:            false,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"username": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"bearer_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *HashicupsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data HashicupsProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://api.hookdeck.com/{version}"
	}

	var basicAuth *shared.SchemeBasicAuth
	password := data.Password.ValueString()
	username := data.Username.ValueString()
	basicAuth = &shared.SchemeBasicAuth{
		Password: password,
		Username: username,
	}
	bearerAuth := new(string)
	if !data.BearerAuth.IsUnknown() && !data.BearerAuth.IsNull() {
		*bearerAuth = data.BearerAuth.ValueString()
	} else {
		bearerAuth = nil
	}
	security := shared.Security{
		BasicAuth:  basicAuth,
		BearerAuth: bearerAuth,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *HashicupsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewConnectionResource,
		NewDestinationResource,
		NewIssueTriggerResource,
		NewSourceResource,
	}
}

func (p *HashicupsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewConnectionDataSource,
		NewDestinationDataSource,
		NewIssueTriggerDataSource,
		NewSourceDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &HashicupsProvider{
			version: version,
		}
	}
}
