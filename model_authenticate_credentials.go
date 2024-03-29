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

type AuthenticateCredentials struct {
	Username string `json:"username,omitempty"`
	Msg string `json:"msg,omitempty"`
	Sessiontoken string `json:"sessiontoken,omitempty"`
	Groups []string `json:"groups,omitempty"`
	AdditionalAttributes []AuthenticateCredentialsAdditionalAttributes `json:"additionalAttributes,omitempty"`
}
