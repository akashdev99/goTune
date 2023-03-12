package pluginloader

import (
	"gotune/workflowmanager/models"
	"log"
	"plugin"
)

var setupPluginList []string = []string{"./plugin/setup/setup.so"}
var setupVar []string = []string{"Setup"}

func LoadSetupPlugins() (pluginList []models.SetupPlugin, err error) {
	for i, setupPlugin := range setupPluginList {
		plugin, err := plugin.Open(setupPlugin)
		if err != nil {
			log.Fatal("plugin directory open failed: ", err)
			return nil, err
		}

		setupPlugin, err := plugin.Lookup(setupVar[i])
		if err != nil {
			log.Fatal("plugin lookup failed: ", err)
			return nil, err
		}

		setupInterface, ok := setupPlugin.(models.SetupPlugin)
		if !ok {
			log.Fatalf("unexpected type from module symbol")
			return nil, err
		}
		pluginList = append(pluginList, setupInterface)
	}
	return
}
