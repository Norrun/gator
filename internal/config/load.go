package config

import (
	"encoding/json"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

func Read() Config {
	configPath, err := getConfigFilePath()
	if err != nil {
		println(err.Error())
	}
	configRaw, err := os.ReadFile(configPath)
	configeration := Config{}
	if err != nil {
		println(err.Error())
	}
	err = json.Unmarshal(configRaw, &configeration)

	if err != nil {
		println(err.Error())
	}
	return configeration

}

func (c *Config) SetUser(user string) {
	c.CurrentUserName = user
	err := write(*c)
	if err != nil {
		println(err.Error())
	}
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := path.Join(homeDir, configFileName)

	return configPath, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
