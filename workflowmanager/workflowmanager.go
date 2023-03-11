package workflowmanager

import (
	"gotune/configmanager"
	"gotune/workflowmanager/pluginloader"
	"log"
)

func Start() {
	_, err := configmanager.ConfigMgr.GetConfig()
	if err != nil {
		log.Fatalf("GOTUNE: Failed to getConfig :%v", err)
	}

	//plugin load setup
	for _, setupPlugins := range pluginloader.LoadSetupPlugins() {
		setupPlugins.Setup()
		setupPlugins.Description()
	}

	// 6) run event generator
	//kill the process
	//loop again

}
