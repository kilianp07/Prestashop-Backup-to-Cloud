package config

import (
	"encoding/json"
	"os"
)

// GetSSHConfig returns a ssh.ClientConfig struct.
func (c *Config) GetSSHConfig() *SshConf {
	return &c.SshConfig
}

// Function New returns a filled Config struct.
func New() (*Config, error) {
	return readConfig()
}

// readConfig opens the config.json file and decodes it into a Config struct.
func readConfig() (*Config, error) {
	config := &Config{}
	// open config file
	file, err := os.Open("/conf/config.json")
	if err != nil {
		return nil, err
	}
	// decode json into config struct
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
