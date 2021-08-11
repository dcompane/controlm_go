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

type ComponentMetaDataProperties struct {
	DisplayName string `json:"displayName,omitempty"`
	DisplayNameID string `json:"displayNameID,omitempty"`
	Name string `json:"name,omitempty"`
	Sections []SectionMetadataProperties `json:"sections,omitempty"`
}