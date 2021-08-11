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

type Availability struct {
	// the type of the availability stat
	Type_ string `json:"Type,omitempty"`
	// the name of the stat
	Name string `json:"Name,omitempty"`
	// the current status
	Status string `json:"Status,omitempty"`
	// A message representing the problem
	Message string `json:"Message,omitempty"`
}