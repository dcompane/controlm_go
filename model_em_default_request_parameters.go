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

type EmDefaultRequestParameters struct {
	AnnotationCategory string `json:"annotation_category,omitempty"`
	AnnotationNotes string `json:"annotation_notes,omitempty"`
	CtmName string `json:"ctm_name,omitempty"`
	NetName string `json:"net_name,omitempty"`
}