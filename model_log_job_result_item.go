/*
 * Control-M Services
 *
 * Provides access to BMC Control-M Services
 *
 * API version: 9.20.115
 * Contact: customer_support@bmc.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package controlm_go

type LogJobResultItem struct {
	DataArguments []LogDataArguments `json:"data_arguments,omitempty"`
	FormattedMessage string `json:"formatted_message,omitempty"`
	FullLine string `json:"full_line,omitempty"`
	LocalTimestampIso8601 string `json:"local_timestamp_iso8601,omitempty"`
	Message string `json:"message,omitempty"`
	MessageCode string `json:"message_code,omitempty"`
}