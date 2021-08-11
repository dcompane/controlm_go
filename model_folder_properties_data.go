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

type FolderPropertiesData struct {
	// The name of the folder. REQUIRED:addMFTFolder | HIDDEN
	Name string `json:"name,omitempty"`
	// Authorized Internal Users. HIDDEN
	AuthorizedInternalUsers []string `json:"authorizedInternalUsers,omitempty"`
	// Authorized External Users And User Groups. HIDDEN
	AuthorizedExternalUsersAndGroups []string `json:"authorizedExternalUsersAndGroups,omitempty"`
	// Delete file after downloaded from incoming folder. HIDDEN
	DeleteFilesAfterDownload bool `json:"deleteFilesAfterDownload,omitempty"`
	// Send email notification to external users when a new file arrives. HIDDEN
	NotifyByEmailWhenFileArrive bool `json:"notifyByEmailWhenFileArrive,omitempty"`
	// Retention Time in hours. HIDDEN
	RetentionPolicy int64 `json:"retentionPolicy,omitempty"`
	// Size limit for folder (in Gigabyte). HIDDEN
	SizeLimit int64 `json:"sizeLimit,omitempty"`
	// allowed file pattern wildcard. HIDDEN
	AllowedFilePattern string `json:"allowedFilePattern,omitempty"`
	// blocked file pattern wildcard. HIDDEN
	BlockedFilePattern string `json:"blockedFilePattern,omitempty"`
	// Folder's access type (Limited, Unlimited). HIDDEN
	AccessType string `json:"accessType,omitempty"`
}
