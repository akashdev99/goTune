package configmanager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

type ConfigList []map[string]interface{}

type ConfigManager struct {
	Config          ConfigList `yaml:"config"`
	ConfigDirectory string     `yaml:"configDirectory"`
}

var (
	ConfigMgr *ConfigManager
	err       error
	once      sync.Once
)

func GetInstance() (*ConfigManager, error) {
	var err error

	once.Do(func() {
		ConfigMgr = &ConfigManager{}
		err = ConfigMgr.load()
	})
	return ConfigMgr, err
}

var FILE_DIR = "./configmanager/config.yaml"

func (c *ConfigManager) load() error {
	file, err := ioutil.ReadFile(FILE_DIR)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Fatal("Not compatible", err)
	}
	if isValid, err := c.validateConfig(); !isValid {
		c = &ConfigManager{}
		return fmt.Errorf("config load failed :%v", err)
	}
	return nil
}

func (configmgr ConfigManager) validateConfig() (bool, error) {
	//check if all config object has equal number of keys - uniform config (maye blater make it take default vlaues)
	return configmgr.isUniformPropertCount()
}

func (configmgr ConfigManager) isUniformPropertCount() (bool, error) {
	configList := configmgr.Config
	propCount := len(configList[0])

	tempCount := 0
	for _, config := range configList {
		tempCount = len(config)
		if tempCount != propCount {
			return false, errors.New("uniform number of properties not found")
		}
	}
	return true, nil
}
