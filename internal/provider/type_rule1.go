// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type Rule1 struct {
	Any        *Any1        `tfsdk:"any"`
	AlertRule  *AlertRule1  `tfsdk:"alert_rule"`
	DelayRule  *DelayRule1  `tfsdk:"delay_rule"`
	FilterRule *FilterRule1 `tfsdk:"filter_rule"`
	RetryRule  *RetryRule1  `tfsdk:"retry_rule"`
}