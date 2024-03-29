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

type CertificateSigningRequestData struct {
	// The organization HIDDEN
	Organization string `json:"organization,omitempty"`
	// The organizationUnit HIDDEN
	OrganizationUnit string `json:"organizationUnit,omitempty"`
	// The cityLocality HIDDEN
	CityLocality string `json:"cityLocality,omitempty"`
	// The stateProvince HIDDEN
	StateProvince string `json:"stateProvince,omitempty"`
	// The country HIDDEN
	Country string `json:"country,omitempty"`
	// The emailAddress HIDDEN
	EmailAddress string `json:"emailAddress,omitempty"`
}
