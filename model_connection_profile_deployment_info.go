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

type ConnectionProfileDeploymentInfo struct {
	// The logical name of Control-M/Server
	CtmName string `json:"ctmName,omitempty"`
	// The deployment status of connection profile
	Status string `json:"status,omitempty"`
	// The deployment status code of connection profile
	StatusCode int32 `json:"statusCode,omitempty"`
	// UTC date of the modification
	LastUpdate string `json:"lastUpdate,omitempty"`
	// Status information
	Message string `json:"message,omitempty"`
}