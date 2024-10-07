package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFileName), nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return errors.New("Unable to write db: " + err.Error())
	}
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func Read() (*Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var con Config
	err = json.Unmarshal(data, &con)
	if err != nil {
		return &Config{}, errors.New("Unable to load db: " + err.Error())
	}
	return &con, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg, err := Read()
	if err != nil {
		return err
	}
	cfg.CurrentUserName = user
	return write(*cfg)
}
