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

type ErrorData struct {
	// The error message that describes the problem.
	Message string `json:"message,omitempty"`
	// An internal identifier of the error.
	Id string `json:"id,omitempty"`
	// Reference to the item this error is referring to.
	Item string `json:"item,omitempty"`
	// The file this error occurred in.
	File string `json:"file,omitempty"`
	// The number of line in the file this error occurred in.
	Line int32 `json:"line,omitempty"`
	// The number of column in the file this error occurred in.
	Col int32 `json:"col,omitempty"`
}
