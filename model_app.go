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

// App (job type) object for AI deploy api
type App struct {
	// job type display name
	Displayname string `json:"displayname,omitempty"`
	// job type name
	Name string `json:"name,omitempty"`
	// description
	Desc string `json:"desc,omitempty"`
	// default type
	DefaultType string `json:"defaultType,omitempty"`
	// last modification date
	LastModified string `json:"lastModified,omitempty"`
	// creation date
	CreatedOn string `json:"createdOn,omitempty"`
	// creation author
	CreatedBy string `json:"createdBy,omitempty"`
}
