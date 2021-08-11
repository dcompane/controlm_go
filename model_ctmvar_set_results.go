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

type CtmvarSetResults struct {
	CompletedStatus *ResultsStatus `json:"completed_status,omitempty"`
	Error_ *ApiThrowable `json:"error,omitempty"`
	From int64 `json:"from,omitempty"`
	Results []CtmvarSetResultItem `json:"results,omitempty"`
	To int64 `json:"to,omitempty"`
}
