// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type Rule5 struct {
	RetryRule     *RetryRule     `tfsdk:"retry_rule"`
	AlertRule     *AlertRule     `tfsdk:"alert_rule"`
	FilterRule    *FilterRule5   `tfsdk:"filter_rule"`
	TransformRule *TransformRule `tfsdk:"transform_rule"`
	DelayRule     *DelayRule     `tfsdk:"delay_rule"`
}