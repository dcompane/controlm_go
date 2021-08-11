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

type VariableNames struct {
	// Array of pool variables in format %%\\\\PoolName\\AUTOVarInPool. HIDDEN.
	Variables []string `json:"variables,omitempty"`
}
