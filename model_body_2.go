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

type Body2 struct {
	// A file that contains definitions to be compiled. Can be either a JSON definition file, or a zip file that contains multiple JSON definition files.
	DefinitionsFile **os.File `json:"definitionsFile"`
	// Deploy Descriptor JSON file.
	DeployDescriptorFile **os.File `json:"deployDescriptorFile,omitempty"`
}
