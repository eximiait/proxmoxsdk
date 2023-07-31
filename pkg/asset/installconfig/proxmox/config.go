package proxmox_installconfig

import (
	"crypto/x509"
	"os"
	"path/filepath"

	proxmox_wrapper "github.com/eximiait/proxmoxsdk/pkg/wrapper"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var defaultProxmoxConfigEnvVar = "PROXMOX_CONFIG"
var defaultProxmoxConfigPath = filepath.Join(os.Getenv("HOME"), ".proxmox", "proxmox-config.yaml")

// Config holds Proxmox api access details
type Config struct {
	URL      string `yaml:"proxmox_url"`
	FQDN     string `yaml:"proxmox_fqdn"`
	Username string `yaml:"proxmox_username"`
	Password string `yaml:"proxmox_password"`
}

// clientHTTP struct - Hold info about http calls
type clientHTTP struct {
	urlAddr    string // URL or Address
	skipVerify bool   // skipt cert validatin in the http call
	certPool   *x509.CertPool
}

// LoadProxmoxConfig from the following location (first wins):
// 1. PROXMOX_CONFIG env variable
// 2  $defaultProxmoxConfigPath
// See #@Config for the expected format
func LoadProxmoxConfig() ([]byte, error) {
	data, err := os.ReadFile(discoverPath())
	if err != nil {
		return nil, err
	}
	return data, nil
}

// NewConfig will return an Config by loading
// the configuration from locations specified in @LoadProxmoxConfig
func NewConfig() (Config, error) {
	c := Config{}
	in, err := LoadProxmoxConfig()
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(in, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func discoverPath() string {
	path, _ := os.LookupEnv(defaultProxmoxConfigEnvVar)
	if path != "" {
		return path
	}

	return defaultProxmoxConfigPath
}

// Save will serialize the config back into the locations
// specified in @LoadProxmoxConfig, first location with a file, wins.
func (c *Config) Save() error {
	out, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	path := discoverPath()
	err = os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0600)
}

// getValidatedConnection will create a connection and validate it before returning.
func (c *Config) getValidatedConnection() (*proxmox_wrapper.Connection, error) {
	connection, err := getConnection(*c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build configuration for proxmox connection validation")
	}

	if err := connection.Test(); err != nil {
		_ = connection.Close()
		return nil, err
	}
	return connection, err
}
