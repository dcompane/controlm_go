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

type PasswordsObject struct {
	// user name HIDDEN
	User string `json:"user,omitempty"`
	// current password
	CurrentPassword string `json:"currentPassword,omitempty"`
	// new password
	NewPassword string `json:"newPassword,omitempty"`
}