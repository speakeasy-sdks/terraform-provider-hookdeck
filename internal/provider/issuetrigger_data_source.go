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
var _ datasource.DataSource = &IssueTriggerDataSource{}
var _ datasource.DataSourceWithConfigure = &IssueTriggerDataSource{}

func NewIssueTriggerDataSource() datasource.DataSource {
	return &IssueTriggerDataSource{}
}

// IssueTriggerDataSource is the data source implementation.
type IssueTriggerDataSource struct {
	client *sdk.SDK
}

// IssueTriggerDataSourceModel describes the data model.
type IssueTriggerDataSourceModel struct {
	Channels   *IssueTriggerChannels `tfsdk:"channels"`
	Configs    IssueTriggerReference `tfsdk:"configs"`
	CreatedAt  types.String          `tfsdk:"created_at"`
	DeletedAt  types.String          `tfsdk:"deleted_at"`
	DisabledAt types.String          `tfsdk:"disabled_at"`
	ID         types.String          `tfsdk:"id"`
	Name       types.String          `tfsdk:"name"`
	TeamID     types.String          `tfsdk:"team_id"`
	Type       types.String          `tfsdk:"type"`
	UpdatedAt  types.String          `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *IssueTriggerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_issue_trigger"
}

// Schema defines the schema for the data source.
func (r *IssueTriggerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IssueTrigger DataSource",

		Attributes: map[string]schema.Attribute{
			"channels": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"email": schema.SingleNestedAttribute{
						Computed:    true,
						Attributes:  map[string]schema.Attribute{},
						Description: `Email channel for an issue trigger`,
					},
					"opsgenie": schema.SingleNestedAttribute{
						Computed:    true,
						Attributes:  map[string]schema.Attribute{},
						Description: `Integration channel for an issue trigger`,
					},
					"slack": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"channel_name": schema.StringAttribute{
								Computed:    true,
								Description: `Channel name`,
							},
						},
						Description: `Slack channel for an issue trigger`,
					},
				},
				Description: `Notification channels object for the specific channel type`,
			},
			"configs": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"issue_trigger_backpressure_configs": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"delay": schema.Int64Attribute{
								Computed:    true,
								Description: `The minimum delay (backpressure) to open the issue for min of 1 minute (60000) and max of 1 day (86400000)`,
							},
							"destinations": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"str": schema.StringAttribute{
										Computed: true,
									},
									"array_ofstr": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
								},
								Validators: []validator.Object{
									validators.ExactlyOneChild(),
								},
								Description: `A pattern to match on the destination name or array of destination IDs. Use ` + "`" + `*` + "`" + ` as wildcard.`,
							},
						},
						Description: `Configurations for a 'Backpressure' issue trigger`,
					},
					"issue_trigger_delivery_configs": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"connections": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"str": schema.StringAttribute{
										Computed: true,
									},
									"array_ofstr": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
								},
								Validators: []validator.Object{
									validators.ExactlyOneChild(),
								},
								Description: `A pattern to match on the connection name or array of connection IDs. Use ` + "`" + `*` + "`" + ` as wildcard.`,
							},
							"strategy": schema.StringAttribute{
								Computed: true,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"first_attempt",
										"final_attempt",
									),
								},
								MarkdownDescription: `must be one of ["first_attempt", "final_attempt"]` + "\n" +
									`The strategy uses to open the issue`,
							},
						},
						Description: `Configurations for a 'delivery' issue trigger`,
					},
					"issue_trigger_transformation_configs": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"log_level": schema.StringAttribute{
								Computed: true,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"debug",
										"info",
										"warn",
										"error",
										"fatal",
									),
								},
								MarkdownDescription: `must be one of ["debug", "info", "warn", "error", "fatal"]` + "\n" +
									`The minimum log level to open the issue on`,
							},
							"transformations": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"str": schema.StringAttribute{
										Computed: true,
									},
									"array_ofstr": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
								},
								Validators: []validator.Object{
									validators.ExactlyOneChild(),
								},
								Description: `A pattern to match on the transformation name or array of transformation IDs. Use ` + "`" + `*` + "`" + ` as wildcard.`,
							},
						},
						Description: `Configurations for a 'Transformation' issue trigger`,
					},
				},
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
				Description: `Configuration object for the specific issue type selected`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `ISO timestamp for when the issue trigger was created`,
			},
			"deleted_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `ISO timestamp for when the issue trigger was deleted`,
			},
			"disabled_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `ISO timestamp for when the issue trigger was disabled`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the issue trigger`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Optional unique name to use as reference when using the API`,
			},
			"team_id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the workspace`,
			},
			"type": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"delivery",
						"transformation",
						"backpressure",
					),
				},
				MarkdownDescription: `must be one of ["delivery", "transformation", "backpressure"]` + "\n" +
					`Issue type`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `ISO timestamp for when the issue trigger was last updated`,
			},
		},
	}
}

func (r *IssueTriggerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *IssueTriggerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *IssueTriggerDataSourceModel
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
	request := operations.GetIssueTriggerRequest{
		ID: id,
	}
	res, err := r.client.IssueTrigger.Get(ctx, request)
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
	if res.IssueTrigger == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.IssueTrigger)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
