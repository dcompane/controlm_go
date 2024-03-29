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

type MonitoringPrivilegeCategory struct {
	// Alerts access level (None, Browse, Update, Full)
	Alert string `json:"Alert,omitempty"`
	// Archived Viewpoints access level (None, Browse, Update, Full)
	ViewpointArchive string `json:"ViewpointArchive,omitempty"`
}
