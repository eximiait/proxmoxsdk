package proxmox_installconfig

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/openshift/installer/pkg/types"
)

// Validate executes proxmox specific validation
func Validate(ic *types.InstallConfig) error {
	allErrs := field.ErrorList{}
	//proxmoxPlatformPath := field.NewPath("platform", "proxmox")

	/*if ic.Platform.Name() == "" {
		return errors.New(field.Required(
			proxmoxPlatformPath,
			"validation requires a Proxmox platform configuration").Error())
	}

	allErrs = append(
		allErrs,
		validation.ValidatePlatform(ic.Platform.None, proxmoxPlatformPath)...)*/

	con, err := NewConnection()
	if err != nil {
		return err
	}
	defer con.Close()
	return allErrs.ToAggregate()
}
