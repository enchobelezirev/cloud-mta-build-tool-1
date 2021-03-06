package artifacts

import (
	"github.com/pkg/errors"

	"github.com/SAP/cloud-mta-build-tool/internal/fs"
	"github.com/SAP/cloud-mta-build-tool/internal/logs"
	"github.com/SAP/cloud-mta/validations"
)

// ExecuteValidation - executes validation of MTA
func ExecuteValidation(source, desc, mode string, getWorkingDir func() (string, error)) error {
	logs.Logger.Info("validating the MTA project")
	loc, err := dir.Location(source, "", desc, getWorkingDir)
	if err != nil {
		return errors.Wrap(err, "validation failed when initializing the location")
	}
	validateSchema, validateProject, err := validate.GetValidationMode(mode)
	if err != nil {
		return errors.Wrap(err, "validation failed when analyzing the validation mode")
	}
	err = validate.MtaYaml(source, loc.GetMtaYamlFilename(), validateSchema, validateProject)
	if err != nil {
		return err
	}
	return nil
}
