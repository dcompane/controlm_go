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

type CtmvarGetResultItem struct {
	GetPoolVariablesErrorInfo *PoolVariablesErrorInfo `json:"get_pool_variables_error_info,omitempty"`
	PoolVariablesNameValue *PoolVariablesNameValue `json:"pool_variables_name_value,omitempty"`
}