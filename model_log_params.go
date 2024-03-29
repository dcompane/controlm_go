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

type LogParams struct {
	// The name of the Control-M server in which the job was ordered from. HIDDEN.
	Ctm string `json:"ctm,omitempty"`
	// Job's order id. HIDDEN.
	OrderId string `json:"orderId,omitempty"`
	// Job's rerun number. HIDDEN.
	NumberOfRuns int32 `json:"numberOfRuns,omitempty"`
	// Job's log table name, accepted as a value in search request. HIDDEN.
	LogTable string `json:"logTable,omitempty"`
}
