// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type Rule2 struct {
	Any        *Any         `tfsdk:"any"`
	AlertRule  *AlertRule   `tfsdk:"alert_rule"`
	DelayRule  *DelayRule   `tfsdk:"delay_rule"`
	FilterRule *FilterRule2 `tfsdk:"filter_rule"`
	RetryRule  *RetryRule   `tfsdk:"retry_rule"`
}
