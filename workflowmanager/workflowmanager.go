package workflowmanager

import (
	"fmt"
	"gotune/configmanager"
	"gotune/plugin/report"
	"log"
)

func Start() error {
	configMgr, err := configmanager.GetInstance()
	if err != nil {
		log.Fatalf("GOTUNE: Failed load confgi manager :%v", err)
		return err
	}

	configList := configMgr.Config
	dir := configMgr.ConfigDirectory
	for _, config := range configList {

		//SETUP STAGE
		//plugin load setup
		if err := setupStage(config, dir); err != nil {
			return err
		}

		//REPORT STAGE
		if err := report.ReportInstance.Sample(configMgr.ReporterLife, configMgr.ReporterInterval); err != nil {
			return err
		}

	}
	// 6) run event generator
	//kill the process
	//loop again
	return nil
}

func setupStage(config map[string]interface{}, dir string) error {
	// setupPluginList, err := pluginloader.LoadSetupPlugins()
	// if err != nil {
	// 	return err
	// }
	// for _, setupPlugins := range setupPluginList {
	// 	if err := setupPlugins.Setup(config, dir); err != nil {
	// 		return err
	// 	}
	// }

	// if err := setup.SetupInstance.Setup(config, dir); err != nil {
	// 	return err
	// }

	fmt.Println("DONE")
	return nil
}

func reportStage() {

}
