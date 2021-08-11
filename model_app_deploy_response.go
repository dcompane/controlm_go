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

// Deployed App (job type) object response for AI deploy api
type AppDeployResponse struct {
	// status
	Status string `json:"status,omitempty"`
	// message
	Message string `json:"message,omitempty"`
	// agent name
	AgentName string `json:"agentName,omitempty"`
}