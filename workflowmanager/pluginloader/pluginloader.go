package pluginloader

import (
	"fmt"
	"gotune/workflowmanager/models"
	"os"
	"plugin"
)

var setupPluginList []string = []string{"./plugin/setup/setup.so"}

func LoadSetupPlugins() (pluginList []models.SetupPlugin) {
	for _, setupPlugin := range setupPluginList {
		plugin, err := plugin.Open(setupPlugin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//name convention
		//upper case first letter of plugin name
		setupPlugin, err := plugin.Lookup("Setup")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		setupInterface, ok := setupPlugin.(models.SetupPlugin)
		if !ok {
			fmt.Println("unexpected type from module symbol")
			os.Exit(1)
		}
		pluginList = append(pluginList, setupInterface)
	}
	return
}
