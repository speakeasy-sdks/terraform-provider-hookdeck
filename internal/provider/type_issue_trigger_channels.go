// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type IssueTriggerChannels struct {
	Email    *IssueTriggerEmailChannel       `tfsdk:"email"`
	Opsgenie *IssueTriggerIntegrationChannel `tfsdk:"opsgenie"`
	Slack    *IssueTriggerSlackChannel       `tfsdk:"slack"`
}
