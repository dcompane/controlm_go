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

type SystemSetting struct {
	Saml2IdentityProvider *Saml2IdentityProvider `json:"saml2IdentityProvider,omitempty"`
	NewDayTime *SystemSettingProperty `json:"newDayTime,omitempty"`
	FirstDayOfTheWeek *SystemSettingProperty `json:"firstDayOfTheWeek,omitempty"`
	EnvironmentBannerColor *SystemSettingProperty `json:"environmentBannerColor,omitempty"`
	EnvironmentTitle *SystemSettingProperty `json:"environmentTitle,omitempty"`
	EnvironmentDescription *SystemSettingProperty `json:"environmentDescription,omitempty"`
	EnforceSiteStandards *SystemSettingProperty `json:"enforceSiteStandards,omitempty"`
	Strictnesslevel *SystemSettingProperty `json:"strictnesslevel,omitempty"`
	HistoryRetentionDays *SystemSettingProperty `json:"historyRetentionDays,omitempty"`
	Annotations *[]string `json:"annotations,omitempty"`
	Errors []ErrorData `json:"errors,omitempty"`
}
