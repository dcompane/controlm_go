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

type ArchiveJobsList struct {
	Jobs []Job `json:"jobs,omitempty"`
	Returned int32 `json:"returned,omitempty"`
	AdditionalJobsMatchSearchCriteria bool `json:"additionalJobsMatchSearchCriteria,omitempty"`
}