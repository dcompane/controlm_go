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

type StatisticsAverageInfo struct {
	// Average start time
	StartTime string `json:"startTime,omitempty"`
	// Average CPU time
	CpuTime string `json:"cpuTime,omitempty"`
	// Average time
	RunTime string `json:"runTime,omitempty"`
}
