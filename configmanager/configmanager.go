package configmanager

import (
	"errors"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type configmap map[string][]map[string]interface{}

type ConfigManager struct {
	config configmap
}

var (
	ConfigMgr ConfigManager
	err       error
)

func init() {
	//Init or singleton load ???
	err = load()
}

const FILE_DIR = "./configmanager/config.yaml"

func load() error {
	file, err := ioutil.ReadFile(FILE_DIR)
	if err != nil {
		log.Fatal(err)
	}

	config := make(map[string][]map[string]interface{})
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	ConfigMgr.config = config
	if !ConfigMgr.validateConfig() {
		return errors.New("config load failed")
	}
	return nil
}

func (configmgr ConfigManager) GetConfig() (configmap, error) {
	if err != nil {
		return nil, err
	}
	return configmgr.config, nil
}

func (configmgr ConfigManager) validateConfig() bool {
	//check if all config object has equal number of keys - uniform config (maye blater make it take default vlaues)
	return isUniformPropertCount(configmgr.config["config"])
}

func isUniformPropertCount(configList []map[string]interface{}) bool {
	propCount := len(configList[0])

	tempCount := 0
	for _, config := range configList {
		tempCount = len(config)
		if tempCount != propCount {
			return false
		}
	}
	return true
}

// func printData(data map[string][]map[string]interface{}) {
// 	for _, v := range data {
// 		for _, arrayItem := range v {
// 			for k1, v1 := range arrayItem {
// 				fmt.Printf("%s -> %d\n", k1, v1)
// 			}
// 		}
// 	}
// }
