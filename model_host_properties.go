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

type HostProperties struct {
	// Host Name HIDDEN
	HostName string `json:"hostName,omitempty"`
	// OS Type (Windows/Unix/MVS/OS400/Tandem/OpenVMS/OS2200) HIDDEN
	OsType string `json:"osType,omitempty"`
	// User Name HIDDEN
	User string `json:"user,omitempty"`
	// Password HIDDEN
	Password string `json:"password,omitempty"`
	// Protocol Name (Local/FTP/SFTP) HIDDEN
	Protocol string `json:"protocol,omitempty"`
	// Port HIDDEN
	Port int32 `json:"port,omitempty"`
	// FTP Over SSL/TLS HIDDEN
	Ssl bool `json:"ssl,omitempty"`
	// SSL Implicit HIDDEN
	SslImplicit bool `json:"sslImplicit,omitempty"`
	// Clear Command Channel (CCC) HIDDEN
	CccCommand bool `json:"cccCommand,omitempty"`
	// Clear Data Channel (CDC) HIDDEN
	CdcCommand bool `json:"cdcCommand,omitempty"`
	// SSL Security Level (No Authentication/Server Authentication/Client-Server Authentication) HIDDEN
	SslLevel string `json:"sslLevel,omitempty"`
	// FTP Passive HIDDEN
	Passive bool `json:"passive,omitempty"`
	// Substitute IP Address HIDDEN
	SubstituteIpAddress bool `json:"substituteIpAddress,omitempty"`
	// Extrernded Passive Mode (EPSV) HIDDEN
	ExtendedPassiveMode bool `json:"extendedPassiveMode,omitempty"`
	// SSH Compression HIDDEN
	SshCompression bool `json:"sshCompression,omitempty"`
	// Private Key Name HIDDEN
	LogicalKeyName string `json:"logicalKeyName,omitempty"`
	// Key Passphrase HIDDEN
	Passphrase string `json:"passphrase,omitempty"`
}