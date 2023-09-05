// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"hashicups/internal/sdk/pkg/models/shared"
	"time"
)

func (r *DestinationDataSourceModel) RefreshFromGetResponse(resp *shared.Destination) {
	if resp.ArchivedAt != nil {
		r.ArchivedAt = types.StringValue(resp.ArchivedAt.Format(time.RFC3339))
	} else {
		r.ArchivedAt = types.StringNull()
	}
	if resp.AuthMethod == nil {
		r.AuthMethod = nil
	} else {
		r.AuthMethod = &DestinationAuthMethodConfig{}
		if resp.AuthMethod.HookdeckSignature != nil {
			r.AuthMethod.HookdeckSignature = &HookdeckSignature{}
			if r.AuthMethod.HookdeckSignature.Config == nil {
				r.AuthMethod.HookdeckSignature.Config = &DestinationAuthMethodSignatureConfig{}
			}
			if resp.AuthMethod.HookdeckSignature.Config == nil {
				r.AuthMethod.HookdeckSignature.Config = nil
			} else {
				r.AuthMethod.HookdeckSignature.Config = &DestinationAuthMethodSignatureConfig{}
			}
			r.AuthMethod.HookdeckSignature.Type = types.StringValue(string(resp.AuthMethod.HookdeckSignature.Type))
		}
		if resp.AuthMethod.BasicAuth != nil {
			r.AuthMethod.BasicAuth = &BasicAuth{}
			if r.AuthMethod.BasicAuth.Config == nil {
				r.AuthMethod.BasicAuth.Config = &DestinationAuthMethodBasicAuthConfig{}
			}
			if resp.AuthMethod.BasicAuth.Config == nil {
				r.AuthMethod.BasicAuth.Config = nil
			} else {
				r.AuthMethod.BasicAuth.Config = &DestinationAuthMethodBasicAuthConfig{}
				r.AuthMethod.BasicAuth.Config.Password = types.StringValue(resp.AuthMethod.BasicAuth.Config.Password)
				r.AuthMethod.BasicAuth.Config.Username = types.StringValue(resp.AuthMethod.BasicAuth.Config.Username)
			}
			r.AuthMethod.BasicAuth.Type = types.StringValue(string(resp.AuthMethod.BasicAuth.Type))
		}
		if resp.AuthMethod.APIKey != nil {
			r.AuthMethod.APIKey = &APIKey{}
			if r.AuthMethod.APIKey.Config == nil {
				r.AuthMethod.APIKey.Config = &DestinationAuthMethodAPIKeyConfig{}
			}
			if resp.AuthMethod.APIKey.Config == nil {
				r.AuthMethod.APIKey.Config = nil
			} else {
				r.AuthMethod.APIKey.Config = &DestinationAuthMethodAPIKeyConfig{}
				r.AuthMethod.APIKey.Config.APIKey = types.StringValue(resp.AuthMethod.APIKey.Config.APIKey)
				r.AuthMethod.APIKey.Config.Key = types.StringValue(resp.AuthMethod.APIKey.Config.Key)
				if resp.AuthMethod.APIKey.Config.To != nil {
					r.AuthMethod.APIKey.Config.To = types.StringValue(string(*resp.AuthMethod.APIKey.Config.To))
				} else {
					r.AuthMethod.APIKey.Config.To = types.StringNull()
				}
			}
			r.AuthMethod.APIKey.Type = types.StringValue(string(resp.AuthMethod.APIKey.Type))
		}
		if resp.AuthMethod.BearerToken != nil {
			r.AuthMethod.BearerToken = &BearerToken{}
			if r.AuthMethod.BearerToken.Config == nil {
				r.AuthMethod.BearerToken.Config = &DestinationAuthMethodBearerTokenConfig{}
			}
			if resp.AuthMethod.BearerToken.Config == nil {
				r.AuthMethod.BearerToken.Config = nil
			} else {
				r.AuthMethod.BearerToken.Config = &DestinationAuthMethodBearerTokenConfig{}
				r.AuthMethod.BearerToken.Config.Token = types.StringValue(resp.AuthMethod.BearerToken.Config.Token)
			}
			r.AuthMethod.BearerToken.Type = types.StringValue(string(resp.AuthMethod.BearerToken.Type))
		}
		if resp.AuthMethod.CustomSignature != nil {
			r.AuthMethod.CustomSignature = &CustomSignature{}
			r.AuthMethod.CustomSignature.Config.Key = types.StringValue(resp.AuthMethod.CustomSignature.Config.Key)
			if resp.AuthMethod.CustomSignature.Config.SigningSecret != nil {
				r.AuthMethod.CustomSignature.Config.SigningSecret = types.StringValue(*resp.AuthMethod.CustomSignature.Config.SigningSecret)
			} else {
				r.AuthMethod.CustomSignature.Config.SigningSecret = types.StringNull()
			}
			r.AuthMethod.CustomSignature.Type = types.StringValue(string(resp.AuthMethod.CustomSignature.Type))
		}
	}
	if resp.CliPath != nil {
		r.CliPath = types.StringValue(*resp.CliPath)
	} else {
		r.CliPath = types.StringNull()
	}
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	if resp.HTTPMethod != nil {
		r.HTTPMethod = types.StringValue(string(*resp.HTTPMethod))
	} else {
		r.HTTPMethod = types.StringNull()
	}
	r.ID = types.StringValue(resp.ID)
	r.Name = types.StringValue(resp.Name)
	if resp.PathForwardingDisabled != nil {
		r.PathForwardingDisabled = types.BoolValue(*resp.PathForwardingDisabled)
	} else {
		r.PathForwardingDisabled = types.BoolNull()
	}
	if resp.RateLimit != nil {
		r.RateLimit = types.Int64Value(*resp.RateLimit)
	} else {
		r.RateLimit = types.Int64Null()
	}
	if resp.RateLimitPeriod != nil {
		r.RateLimitPeriod = types.StringValue(string(*resp.RateLimitPeriod))
	} else {
		r.RateLimitPeriod = types.StringNull()
	}
	r.TeamID = types.StringValue(resp.TeamID)
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	if resp.URL != nil {
		r.URL = types.StringValue(*resp.URL)
	} else {
		r.URL = types.StringNull()
	}
}
