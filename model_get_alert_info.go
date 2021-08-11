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

type GetAlertInfo struct {
	Id int32 `json:"id,omitempty"`
	Notes string `json:"notes,omitempty"`
	Severity string `json:"severity,omitempty"`
	Status string `json:"status,omitempty"`
}