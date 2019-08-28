package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//环境变量配置的spm配置文件路径
const SpmConfigEnvVar = "SPM_CONFIG"

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
	MaxAge time.Duration 		`json:"maxAge" yaml:"maxAge"`
	RotationTime time.Duration 	`json:"rotationTime" yaml:"rotationTime"`
}

type EnvironmentConfig struct {
	Profile string		`json:"profile" yaml:"profile"`

	Database Database 	`json:"database" yaml:"database"`

	Server Server 		`json:"server" yaml:"server"`

	Logger Logger 		`json:"logger" yaml:"logger"`
}

type ConfigContent struct {
	Environment string 			`json:"environment" yaml:"environment"`
	Configurations []EnvironmentConfig		`json:"configurations" yaml:"configurations"`
}

type Config struct {
	Environment string
	EnvironmentConfig
}

//全局配置
var Global = Config{}


func (c *Config) findEnvironmentConfig(content *ConfigContent, env string) {
	for _, config := range content.Configurations {
		if config.Profile==env {
			c.EnvironmentConfig = config
			return
		}
	}
}

func (c *Config) loadFromJson(path string) error{
	data, err := ioutil.ReadFile(path)
	if err!=nil {
		return err
	}
	content := &ConfigContent{}
	err = json.Unmarshal(data, content)
	if err!=nil {
		return err
	}
	c.Environment = content.Environment
	c.findEnvironmentConfig(content, content.Environment)
	return nil
}

func (c *Config) loadFromYml(path string) error{
	data, err := ioutil.ReadFile(path)
	if err!=nil {
		return err
	}
	content := &ConfigContent{}
	err = yaml.Unmarshal(data, content)
	if err!=nil {
		return err
	}
	c.Environment = content.Environment
	c.findEnvironmentConfig(content, content.Environment)
	return nil
}

func (c *Config) currentPath() (string, error) {
	dir, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(dir), nil
}

//从环境变量中获取配置文件路径
func (c *Config) fileFromEnvVar() string{
	spmConfigPath := os.Getenv(SpmConfigEnvVar)
	if spmConfigPath!="" {
		if IsExists(spmConfigPath) {
			return spmConfigPath
		}
	}
	return ""
}

func (c *Config) LoadFromEnvVar() (bool, error){
	configFile := c.fileFromEnvVar()
	if configFile=="" {
		return false, nil
	}
	fmt.Println("load config from environment variable:", SpmConfigEnvVar, configFile)
	extName := filepath.Ext(configFile)
	if extName==".yaml" || extName==".yml" {
		return true, c.loadFromYml(configFile)
	}else if extName==".json" {
		return true, c.loadFromJson(configFile)
	}
	return false, nil
}

//Load 加载配置文件，先尝试从环境变量中的配置文件位置中读取，再尝试从程序所在位置下的config目录中读取
func (c *Config) Load() error{
	//首先，尝试从环境变量中读取配置文件
	isLoad, err := c.LoadFromEnvVar()
	if err!=nil {
		return err
	}
	if isLoad {
		return nil
	}
	//从当前目录下读取
	currentPath, err := c.currentPath()
	if err!=nil {
		return err
	}
	fmt.Println("load config from:", currentPath)
	configFile := filepath.Join(currentPath, "config", "spm.yml")
	if IsExists(configFile) {
		return c.loadFromYml(configFile)
	}
	configFile = filepath.Join(currentPath, "config", "spm.yaml")
	if IsExists(configFile) {
		return c.loadFromYml(configFile)
	}
	configFile = filepath.Join(currentPath, "config", "spm.json")
	if IsExists(configFile) {
		return c.loadFromJson(configFile)
	}
	return errors.New("configuration file not found: " + configFile)
}

func (c *Config) IsDev() bool{
	return c.Environment=="dev"
}

func (c *Config) IsProd() bool{
	return c.Environment=="prod"
}

func (c *Config) IsTest() bool{
	return c.Environment=="test"
}

func (c *Config) Is(env string) bool{
	return c.Environment==env
}




