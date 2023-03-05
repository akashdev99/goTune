package configmanager

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type configmap map[string][]map[string]interface{}

type ConfigManager struct {
	config configmap
}

var ConfigMgr ConfigManager

func init() {
	load()
}

const FILE_DIR = "./configmanager/config.yaml"

func load() {
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
}

func (configmgr ConfigManager) GetConfig() configmap {
	return configmgr.config
}

func (configmgr ConfigManager) validateConfig() bool {
	//check if all config object has equal number of keys - uniform config (maye blater make it take default vlaues)
	return false
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
