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

// File Transfer Server general parameters
type FtsGeneralSettings struct {
	// Root path where transferred files are stored. If you want to use a different directory for each logged in user, you must add /${userName} to the path.
	HomeDirectory string `json:"homeDirectory,omitempty"`
	// Allow multiple open sessions
	MultipleLoginAllowed bool `json:"multipleLoginAllowed,omitempty"`
	// Maximum concurrent open sessions
	MaxOpenSessions int32 `json:"maxOpenSessions,omitempty"`
	// Number of retries before the server closes the connection
	MaxLoginFailures int32 `json:"maxLoginFailures,omitempty"`
	// Time of waiting before letting the user to do another login in seconds
	DelayAfterLoginFailure int32 `json:"delayAfterLoginFailure,omitempty"`
	// Allow bandwidth throttling
	ThrottlingActivated bool `json:"throttlingActivated,omitempty"`
	// Maximum simultaneos uploads
	MaxSimultaneousUploads int32 `json:"maxSimultaneousUploads,omitempty"`
	// Maximum simultaneos downloads
	MaxSimultaneousDownloads int32 `json:"maxSimultaneousDownloads,omitempty"`
}