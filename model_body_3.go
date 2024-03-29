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
import (
	"os"
)

type Body3 struct {
	// A .ctmai file that contains definitions of jobtype to be deployed to the server.
	DefinitionsFile **os.File `json:"definitionsFile"`
	// A JSON file that contains specifications of an agent.
	PayloadFile **os.File `json:"payloadFile,omitempty"`
}
