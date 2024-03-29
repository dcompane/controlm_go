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

// the properties of the new Control-M Agent installed thing to be connected to SaaS Control-M
type AgentThingProperties struct {
	// the logical name to be used for new agent (equivilant to NodeId)
	Name string `json:"name,omitempty"`
	// the physical name of the thing the agent is installed on
	PhysicalName string `json:"physicalName,omitempty"`
	// the agent tag to be associated with the new agent (CMS RBA permissions tag that is)
	Tag string `json:"tag,omitempty"`
}
