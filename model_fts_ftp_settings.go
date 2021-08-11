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

// FTP/FTPS server parameters
type FtsFtpSettings struct {
	// Enable/Disable listening for FTP/S connection
	ServerEnabled bool `json:"serverEnabled,omitempty"`
	// FTP server port
	Port int32 `json:"port,omitempty"`
	// Authentication method being used to connect FTP server
	AuthenticationMethod string `json:"authenticationMethod,omitempty"`
	// Use FTP secure connection (SSL/TLS)
	Secured bool `json:"secured,omitempty"`
	// FTPS keystore file location
	KeystoreFilePath string `json:"keystoreFilePath,omitempty"`
	// Password being used to access the FTPS keystore
	KeystoreFilePassword string `json:"keystoreFilePassword,omitempty"`
	// Ftps server allowed cipher suites (comma-separated). Leave empty to allow all supported cipher suites.
	Ciphers string `json:"ciphers,omitempty"`
	// Implicit negotiation mode - requires that the entire FTP session must be encrypted
	ListenForImplicitConnection bool `json:"listenForImplicitConnection,omitempty"`
	// Passive FTP ports range - for PASV connections, the server will open a random port in this range for client to connect to
	PassivePorts string `json:"passivePorts,omitempty"`
}
