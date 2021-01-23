package config

import (
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

type Info struct {
	apiVersion     string
	apiKey         string
	apiUrl		   string
	configFilePath string
	homeDir        string
}

func (c *Info) GetApiKey() string {
	return c.apiKey
}

func (c *Info) GetConfigFilePath() string {
	return c.configFilePath
}

func (c *Info) GetHomeDir() string {
	return c.homeDir
}

func (c *Info) SetApiKey(apiKey string) *Info {
	c.apiKey = apiKey
	return c
}

func (c *Info) GetUrl() string {
	return c.apiUrl
}


func (c *Info) Complete() error {
	c.homeDir, _ = defaultHomeDir()
	c.configFilePath = defaultConfigFilePath(c.homeDir)
	c.apiUrl = baseUrl()

	return nil
}

func (c *Info) SaveConfig() (bool, error) {
	if _, err := os.Stat(c.homeDir); os.IsNotExist(err) {
		err = os.MkdirAll(c.homeDir, os.ModePerm)
		if err != nil {
			return false, err
		}
	}
	file, err := os.Create(c.configFilePath)
	if err != nil {
		return false, err
	}

	defer file.Close()

	configFileStruct := configFile{
		ApiVersion: c.apiVersion,
		ApiKey:     c.apiKey,
		ApiUrl:		c.apiUrl,
	}

	configFileStr, err := yaml.Marshal(&configFileStruct)
	if err != nil {
		return false, err
	}

	_, err = file.Write(configFileStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func baseUrl() string {
	return "https://napi.arvancloud.com/cdn/4.0"
}

func defaultHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir + "/.arvan", nil
}

func defaultConfigFilePath(homeDir string) string {
	return homeDir + "/config.yml"
}
