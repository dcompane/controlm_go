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

type ValidationProperties struct {
	Max int32 `json:"max,omitempty"`
	Min int32 `json:"min,omitempty"`
	RegexPattern string `json:"regexPattern,omitempty"`
	RegexPatternJava string `json:"regexPatternJava,omitempty"`
	RegexPatternJavaScript string `json:"regexPatternJavaScript,omitempty"`
	RegexPatternNoMatchMessage string `json:"regexPatternNoMatchMessage,omitempty"`
	RegexPatternNoMatchMessageID string `json:"regexPatternNoMatchMessageID,omitempty"`
	Required bool `json:"required,omitempty"`
	RequiredIf *ConditionProperties `json:"requiredIf,omitempty"`
}