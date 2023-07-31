package proxmox_installconfig

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// checkURLResponse performs a GET on the provided urlAddr to ensure that
// the url actually exists. Users can set skipVerify as true or false to
// avoid cert validation. In case of failure, returns error.
func (c *clientHTTP) checkURLResponse() error {

	logrus.Debugf("checking URL response... urlAddr: %s skipVerify: %s", c.urlAddr, strconv.FormatBool(c.skipVerify))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: c.skipVerify,
			RootCAs:            c.certPool,
		},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(c.urlAddr)
	if err != nil {
		return errors.Wrapf(err, "error checking URL response")
	}
	defer resp.Body.Close()

	return nil
}

// askPassword will ask the password to connect to the Engine API.
// The password provided will be added in the Config struct.
// If an error happens, it will ask again username for users.
func askPassword(c *Config) error {
	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Password{
				Message: "Proxmox password",
				Help:    "Password for the chosen username, Press Ctrl+C to change username",
			},
			//Validate: survey.ComposeValidators(survey.Required, authenticated(c)),
		},
	}, &c.Password)

	if err != nil {
		return err
	}
	return nil
}

// askUsername will ask username to connect to the Engine API.
// The username provided will be added in the Config struct.
// Returns Config and error if failure.
func askUsername(c *Config) error {
	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Proxmox username",
				Help:    "The username to connect to the Proxmox API",
				Default: "root@pam",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &c.Username)
	if err != nil {
		return err
	}

	return nil
}

// askCredentials will handle username and password for connecting with Engine
func askCredentials(c Config) (Config, error) {
	loginAttempts := 3
	logrus.Debugf("login attempts available: %d", loginAttempts)
	for loginAttempts > 0 {
		err := askUsername(&c)
		if err != nil {
			return c, err
		}

		err = askPassword(&c)
		if err != nil {
			loginAttempts = loginAttempts - 1
			logrus.Debugf("login attempts now: %d", loginAttempts)
			if loginAttempts == 0 {
				return c, err
			}
		} else {
			break
		}
	}
	return c, nil
}

// engineSetup will ask users: FQDN, execute validations and about
// the credentials. In case of failure, returns Config and error
func engineSetup() (Config, error) {
	engineConfig := Config{}
	httpResource := clientHTTP{}

	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Proxmox FQDN[:PORT]",
				Help:    "The Proxmox FQDN[:PORT] (proxmox.example.com:8006)",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &engineConfig.FQDN)
	if err != nil {
		return engineConfig, err
	}
	logrus.Debug("Proxmox FQDN: ", engineConfig.FQDN)

	// Set c.URL with the API endpoint
	engineConfig.URL = fmt.Sprintf("https://%s/api2/json", engineConfig.FQDN)
	logrus.Debug("Proxmox URL: ", engineConfig.URL)

	// Start creating clientHTTP struct for checking if Engine FQDN is responding
	httpResource.skipVerify = true
	httpResource.urlAddr = engineConfig.URL
	err = httpResource.checkURLResponse()
	if err != nil {
		return engineConfig, err
	}
	return askCredentials(engineConfig)
}
