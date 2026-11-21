package config

import (
	"encoding/json"
	"os"
)

const configFileName string = "gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	data, err := json.Marshal(&c)
	if err != nil {
		return err
	}

	err = os.WriteFile(homeDir+"/."+configFileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	var config Config
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return config, err
	}

	jsonUrl := homeDir + "/." + configFileName
	data, err := os.ReadFile(jsonUrl)

	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}
