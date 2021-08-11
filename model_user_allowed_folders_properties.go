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

type UserAllowedFoldersProperties struct {
	AllowedFoldersNames []string `json:"allowedFoldersNames,omitempty"`
	As2CertificateAlias string `json:"as2CertificateAlias,omitempty"`
	As2Id string `json:"as2Id,omitempty"`
	As2PublicKeyCertificate string `json:"as2PublicKeyCertificate,omitempty"`
	As2TargetFolder string `json:"as2TargetFolder,omitempty"`
	Company string `json:"company,omitempty"`
	Description string `json:"description,omitempty"`
	Email string `json:"email,omitempty"`
	FullName string `json:"fullName,omitempty"`
	HashedPassword string `json:"hashedPassword,omitempty"`
	IsLdapAuth int32 `json:"isLdapAuth,omitempty"`
	Name string `json:"name,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	SshPublicKey string `json:"sshPublicKey,omitempty"`
	IsLocked bool `json:"isLocked,omitempty"`
	LockReason string `json:"lockReason,omitempty"`
	ChangePasswordAtNextLogin bool `json:"changePasswordAtNextLogin,omitempty"`
	PasswordNeverExpires bool `json:"passwordNeverExpires,omitempty"`
	LastSuccessfulLoginTime string `json:"lastSuccessfulLoginTime,omitempty"`
}
