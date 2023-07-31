package proxmox_installconfig

import (
	proxmox_wrapper "github.com/eximiait/proxmoxsdk/pkg/wrapper"
	"github.com/pkg/errors"
)

// getConnection is a convenience method to get a connection to ovirt api
// form a Config Object.
func getConnection(proxmoxConfig Config) (*proxmox_wrapper.Connection, error) {
	con, err := proxmox_wrapper.NewConnectionBuilder().
		URL(proxmoxConfig.URL).
		Username(proxmoxConfig.Username).
		Password(proxmoxConfig.Password).
		Client(proxmoxConfig.URL).
		Build()
	if err != nil {
		return nil, err
	}
	return con, nil
}

// NewConnection returns a new client connection to oVirt's API endpoint.
// It is the responsibility of the caller to close the connection.
func NewConnection() (*proxmox_wrapper.Connection, error) {
	proxmoxConfig, err := NewConfig()
	if err != nil {
		return nil, errors.Wrap(err, "getting Engine configuration")
	}
	con, err := getConnection(proxmoxConfig)
	if err != nil {
		return nil, errors.Wrap(err, "establishing Engine connection")
	}
	return con, nil
}
