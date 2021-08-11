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

type SystemParameter struct {
	// System parameter name
	Name string `json:"name,omitempty"`
	// System parameter value
	Value string `json:"value,omitempty"`
	// System parameter type
	Type_ string `json:"type,omitempty"`
	// System parameter description
	Description string `json:"description,omitempty"`
}