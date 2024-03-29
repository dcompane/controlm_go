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

type ResourceObj struct {
	// Resource name
	Name string `json:"name,omitempty"`
	// Control-M Server hosting the resource
	Ctm string `json:"ctm,omitempty"`
	// The resource q current quantity.
	Available string `json:"available,omitempty"`
	// The resource q max usage value.
	Max int32 `json:"max,omitempty"`
	// Workload Policy.
	WorkloadPolicy string `json:"workloadPolicy,omitempty"`
}
