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

type AssociateData struct {
	// associate typed REQUIRED
	Type_ string `json:"type,omitempty"`
	// associate name REQUIRED
	Name string `json:"name,omitempty"`
}
