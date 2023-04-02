package workflowmanager

import (
	"fmt"
	"gotune/configmanager"
	"gotune/plugin/report"
	"gotune/plugin/stress"
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

		//REPORT STAGE -PARALLELIZE
		//get error via channels

		// stressChan := make(chan error)
		go report.ReportInstance.Sample(configMgr.ReporterLife, configMgr.ReporterInterval)
		go stress.ReportInstance.Stress(configMgr.ReporterLife)

		if err := <-report.ReportInstance.SampleChan; err != nil {
			log.Fatalf("GOTUNE: Sampling failed :%v", err)
			return err
		} else {
			fmt.Println("Sampling Done")
		}

		if err := <-stress.ReportInstance.StressChan; err != nil {
			log.Fatalf("GOTUNE: Stressing failed :%v", err)
			return err
		} else {
			fmt.Println("Stressing Done")
		}

	}

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
