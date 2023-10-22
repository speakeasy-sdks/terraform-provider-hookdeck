// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"hashicups/internal/sdk"
	"hashicups/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"hashicups/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceDataSource{}

func NewSourceDataSource() datasource.DataSource {
	return &SourceDataSource{}
}

// SourceDataSource is the data source implementation.
type SourceDataSource struct {
	client *sdk.SDK
}

// SourceDataSourceModel describes the data model.
type SourceDataSourceModel struct {
	AllowedHTTPMethods []types.String        `tfsdk:"allowed_http_methods"`
	ArchivedAt         types.String          `tfsdk:"archived_at"`
	CreatedAt          types.String          `tfsdk:"created_at"`
	CustomResponse     *SourceCustomResponse `tfsdk:"custom_response"`
	ID                 types.String          `tfsdk:"id"`
	Integration        *SourceIntegration    `tfsdk:"integration"`
	IntegrationID      types.String          `tfsdk:"integration_id"`
	Name               types.String          `tfsdk:"name"`
	TeamID             types.String          `tfsdk:"team_id"`
	UpdatedAt          types.String          `tfsdk:"updated_at"`
	URL                types.String          `tfsdk:"url"`
}

// Metadata returns the data source type name.
func (r *SourceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source"
}

// Schema defines the schema for the data source.
func (r *SourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Source DataSource",

		Attributes: map[string]schema.Attribute{
			"allowed_http_methods": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `List of allowed HTTP methods. Defaults to PUT, POST, PATCH, DELETE.`,
			},
			"archived_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `Date the source was archived`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `Date the source was created`,
			},
			"custom_response": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"body": schema.StringAttribute{
						Computed:    true,
						Description: `Body of the custom response`,
					},
					"content_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"json",
								"text",
								"xml",
							),
						},
						MarkdownDescription: `must be one of ["json", "text", "xml"]` + "\n" +
							`Content type of the custom response`,
					},
				},
				Description: `Custom response object`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the source`,
			},
			"integration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"features": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `List of enabled features`,
					},
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `ID of the integration`,
					},
					"label": schema.StringAttribute{
						Computed:    true,
						Description: `Label of the integration`,
					},
					"provider": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"twitter",
								"stripe",
								"recharge",
								"github",
								"shopify",
								"postmark",
								"typeform",
								"hmac",
								"basic_auth",
								"api_key",
								"xero",
								"svix",
								"zoom",
								"akeneo",
								"adyen",
								"gitlab",
								"property-finder",
								"woocommerce",
								"oura",
								"commercelayer",
								"mailgun",
								"pipedrive",
								"sendgrid",
								"workos",
								"synctera",
								"aws_sns",
							),
						},
						MarkdownDescription: `must be one of ["twitter", "stripe", "recharge", "github", "shopify", "postmark", "typeform", "hmac", "basic_auth", "api_key", "xero", "svix", "zoom", "akeneo", "adyen", "gitlab", "property-finder", "woocommerce", "oura", "commercelayer", "mailgun", "pipedrive", "sendgrid", "workos", "synctera", "aws_sns"]` + "\n" +
							`The provider name`,
					},
				},
				Description: `Integration object`,
			},
			"integration_id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the integration`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Name for the source`,
			},
			"team_id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the workspace`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `Date the source was last updated`,
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: `A unique URL that must be supplied to your webhook's provider`,
			},
		},
	}
}

func (r *SourceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()
	request := operations.GetSourceRequest{
		ID: id,
	}
	res, err := r.client.Source.Get(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Source == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.Source)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
