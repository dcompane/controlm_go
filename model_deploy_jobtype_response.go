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

// Response of deployed jobtype
type DeployJobtypeResponse struct {
	// The name of the deployed .ctmai file
	DeploymentFile string `json:"deploymentFile,omitempty"`
	SuccessfulDeploys []JobtypeAgent `json:"successfulDeploys,omitempty"`
}
