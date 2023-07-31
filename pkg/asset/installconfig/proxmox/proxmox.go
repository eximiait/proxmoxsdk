package proxmox_installconfig

import (
	proxmox "github.com/eximiait/proxmoxsdk/pkg/types/proxmox"
	proxmox_wrapper "github.com/eximiait/proxmoxsdk/pkg/wrapper"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const platformValidationMaxTries = 3

// Platform collects proxmox-specific configuration.
func Platform() (*proxmox.Platform, error) {
	p := proxmox.Platform{}
	var c *proxmox_wrapper.Connection
	proxmoxConfig, err := NewConfig()
	for tries := 0; tries < platformValidationMaxTries; tries++ {
		if err != nil {
			proxmoxConfig, err = engineSetup()
			if err != nil {
				logrus.Error(errors.Wrap(err, "Proxmox configuration failed"))
			}
		}

		if err == nil {
			c, err = proxmoxConfig.getValidatedConnection()
			if err != nil {
				logrus.Error(errors.Wrap(err, "failed to validate Proxmox configuration"))
			} else {
				break
			}
		}
	}
	if err != nil {
		// Last error is not nil, we don't have a valid config.
		return nil, errors.Wrap(err, "maximum retries for configuration exhausted")
	}
	defer c.Close()
	if err = proxmoxConfig.Save(); err != nil {
		return nil, err
	}

	err = askStorage(c, &p)
	if err != nil {
		return &p, err
	}

	err = askNetwork(c, &p)
	if err != nil {
		return &p, err
	}

	err = askVIPs(&p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}
