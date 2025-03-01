package config

import (
	"fmt"
	"os"
	"path/filepath"
	"statd/pkg/utils"

	"gopkg.in/yaml.v3"
)

const EnvVar = "STATD_CONF"

var Cfg *Config
var DefaultPath = "~/.config/statd/config.yaml"

var filePath string

func FilePath() string {
	return filePath
}
func setFilePath(path string) error {
	path, err := utils.NormalizePath(path)
	if err != nil {
		return err
	}
	filePath = path
	return nil
}

func SelectConfigFile(arg string) error {
	if arg != "" {
		if !utils.FileExists(arg) {
			err := createDefaultConfig(arg)
			if err != nil {
				return fmt.Errorf("Failed to select config file. %w", err)
			}
		}
		setFilePath(arg)
		return nil
	}

	env := os.Getenv(EnvVar)
	if env != "" {
		if !utils.FileExists(env) {
			err := createDefaultConfig(env)
			if err != nil {
				return fmt.Errorf("Failed to select config file. %w", err)
			}
		}
		setFilePath(env)
		return nil
	}

	if !utils.FileExists(DefaultPath) {
		createDefaultConfig(DefaultPath)
	}
	setFilePath(DefaultPath)
	return nil
}

func LoadConfig() error {
	if FilePath() == "" {
		err := SelectConfigFile("")
		if err != nil {
			return fmt.Errorf("Faild to load config file. %w", err)
		}
	}
	var err error
	Cfg, err = readConfig(FilePath())
	if err != nil {
		return fmt.Errorf("Faild to load config file. %w", err)
	}
	return nil
}

func readConfig(path string) (*Config, error) {
	path, err := utils.NormalizePath(path)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func saveConfig(cfg *Config, path string) error {
	path, err := utils.NormalizePath(path)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func createDefaultConfig(path string) error {
	path, err := utils.NormalizePath(path)
	if err != nil {
		return err
	}
	return saveConfig(&DefaultConfig, path)
}
