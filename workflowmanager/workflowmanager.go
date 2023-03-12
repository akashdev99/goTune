package workflowmanager

import (
	"fmt"
	"gotune/configmanager"
	"gotune/workflowmanager/pluginloader"
	"log"
)

func Start() {
	configMgr, err := configmanager.GetInstance()
	if err != nil {
		log.Fatalf("GOTUNE: Failed load confgi manager :%v", err)
		return
	}

	configList := configMgr.Config
	dir := configMgr.ConfigDirectory
	for _, config := range configList {

		//SETUP STAGE
		//plugin load setup
		setup(config, dir)

	}
	// 6) run event generator
	//kill the process
	//loop again
}

func setup(config map[string]interface{}, dir string) {
	setupPluginList, err := pluginloader.LoadSetupPlugins()
	if err != nil {
		return
	}

	for _, setupPlugins := range setupPluginList {
		setupPlugins.Setup(config, dir)
	}

	fmt.Println("DONE")
}
