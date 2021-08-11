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

type AlertParam struct {
	// alertIds. HIDDEN.
	AlertIds []int32 `json:"alertIds"`
	// modify urgency. HIDDEN.
	Urgency string `json:"urgency,omitempty"`
	// modify comment. HIDDEN.
	Comment string `json:"comment,omitempty"`
}