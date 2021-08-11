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

type ComponentKeyWithStatusType struct {
	ComponentKey *ComponentMftKeyType `json:"componentKey,omitempty"`
	// Component status
	Status int32 `json:"status,omitempty"`
}