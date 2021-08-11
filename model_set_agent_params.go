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

type SetAgentParams struct {
	// parameter name
	Name string `json:"name,omitempty"`
	// parameter value
	Value string `json:"value,omitempty"`
	// parameter type
	Type_ string `json:"type,omitempty"`
}
