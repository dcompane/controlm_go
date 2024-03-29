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

type GatewayData struct {
	// gateway host name REQUIRED
	Host string `json:"host,omitempty"`
	// gateway port REQUIRED
	Port string `json:"port,omitempty"`
	// gateway status HIDDEN
	Status string `json:"status,omitempty"`
	// general message HIDDEN
	Message string `json:"message,omitempty"`
}
