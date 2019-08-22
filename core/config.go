package core

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Database struct {
	Name string 	`json:"name" yaml:"name"`
	Url string 		`json:"url" yaml:"url"`
}

type Server struct {
	Port string 	`json:"port" yaml:"port"`
}

type Logger struct {
	Level string 		`json:"level" yaml:"level"`
	Path string 		`json:"path" yaml:"path"`
	Filename string 	`json:"filename" yaml:"filename"`
	MaxAge int32 		`json:"maxAge" yaml:"maxAge"`
	RotationTime int32 	`json:"rotationTime" yaml:"rotationTime"`
}

type Config struct {
	Database *Database 	`json:"database" yaml:"database"`

	Server *Server 		`json:"server" yaml:"server"`

	Logger *Logger 		`json:"logger" yaml:"logger"`
}

//全局配置
var Global = Config{}

func (c *Config) loadFromJson(path string) error{
	data, err := ioutil.ReadFile(path)
	if err!=nil {
		return err
	}
	return json.Unmarshal(data, c)
}

func (c *Config) loadFormYml(path string) error{
	data, err := ioutil.ReadFile(path)
	if err!=nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}

func (c *Config) currentPath() (string, error) {
	dir, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(dir), nil
}

func (c *Config) defaultConfig() {

}

func (c *Config) Load() error{
	currentPath, err := c.currentPath()
	if err!=nil {
		return err
	}
	configFile := filepath.Join(currentPath, "spm.yml")
	if IsExists(configFile) {
		return c.loadFormYml(configFile)
	}
	configFile = filepath.Join(currentPath, "spm.json")
	if IsExists(configFile) {
		return c.loadFromJson(configFile)
	}
	return errors.New("configuration file not found: " + configFile)
}




