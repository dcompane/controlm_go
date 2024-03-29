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

type SettingsMetadataProperties struct {
	Components []ComponentMetaDataProperties `json:"components,omitempty"`
	Error_ string `json:"error,omitempty"`
	SelectedComponentAfterActivation string `json:"selectedComponentAfterActivation,omitempty"`
	SelectedComponentBeforeActivation string `json:"selectedComponentBeforeActivation,omitempty"`
}
