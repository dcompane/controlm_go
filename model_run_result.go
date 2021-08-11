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

type RunResult struct {
	// An ID that identifies running jobs and can be used to track their status.
	RunId string `json:"runId"`
	// A URI that can be used to get the status of the run jobs.
	StatusURI string `json:"statusURI,omitempty"`
	// A URI to a page displaying the workflow run live.
	MonitorPageURI string `json:"monitorPageURI,omitempty"`
}