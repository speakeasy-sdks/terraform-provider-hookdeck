// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"hashicups/internal/sdk"
	"hashicups/internal/sdk/pkg/models/operations"
	"hashicups/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"hashicups/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceResource{}
var _ resource.ResourceWithImportState = &SourceResource{}

func NewSourceResource() resource.Resource {
	return &SourceResource{}
}

// SourceResource defines the resource implementation.
type SourceResource struct {
	client *sdk.SDK
}

// SourceResourceModel describes the resource data model.
type SourceResourceModel struct {
	AllowedHTTPMethods []types.String        `tfsdk:"allowed_http_methods"`
	ArchivedAt         types.String          `tfsdk:"archived_at"`
	Body               types.String          `tfsdk:"body"`
	ContentType        types.String          `tfsdk:"content_type"`
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

func (r *SourceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source"
}

func (r *SourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Source Resource",

		Attributes: map[string]schema.Attribute{
			"allowed_http_methods": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional:    true,
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
			"body": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required:    true,
				Description: `Body of the custom response`,
			},
			"content_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"json",
						"text",
						"xml",
					),
				},
				MarkdownDescription: `must be one of [json, text, xml]` + "\n" +
					`Content type of the custom response`,
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
						MarkdownDescription: `must be one of [json, text, xml]` + "\n" +
							`Content type of the custom response`,
					},
				},
				Description: `Custom response object`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
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
						MarkdownDescription: `must be one of [twitter, stripe, recharge, github, shopify, postmark, typeform, hmac, basic_auth, api_key, xero, svix, zoom, akeneo, adyen, gitlab, property-finder, woocommerce, oura, commercelayer, mailgun, pipedrive, sendgrid, workos, synctera, aws_sns]` + "\n" +
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required:    true,
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

func (r *SourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	var allowedHTTPMethods []shared.SourceAllowedHTTPMethod = nil
	for _, allowedHTTPMethodsItem := range data.AllowedHTTPMethods {
		allowedHTTPMethods = append(allowedHTTPMethods, shared.SourceAllowedHTTPMethod(allowedHTTPMethodsItem.ValueString()))
	}
	var customResponse *shared.SourceCustomResponse
	if data != nil {
		body := data.Body.ValueString()
		contentType := shared.SourceCustomResponseContentType(data.ContentType.ValueString())
		customResponse = &shared.SourceCustomResponse{
			Body:        body,
			ContentType: contentType,
		}
	}
	name := data.Name.ValueString()
	request := operations.CreateSourceRequestBody{
		AllowedHTTPMethods: allowedHTTPMethods,
		CustomResponse:     customResponse,
		Name:               name,
	}
	res, err := r.client.Source.CreateSource(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	data.RefreshFromCreateResponse(res.Source)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	res, err := r.client.Source.GetSource(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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

func (r *SourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.DeleteSourceRequest{
		ID: id,
	}
	res, err := r.client.Source.DeleteSource(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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

}

func (r *SourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}