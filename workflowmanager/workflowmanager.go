package workflowmanager

import (
	"gotune/configmanager"
	"log"
)

type SetupPlugin interface {
	Description()
	Setup()
}

type StressPlugin interface {
	Description()
	Stress()
}

type CleanUpPlugin interface {
	Description()
	Clean()
}

func Start() {
	_, err := configmanager.ConfigMgr.GetConfig()
	if err != nil {
		log.Fatalf("GOTUNE: Failed to getConfig :%v", err)
	}

	//loop through the config and run the daemon
	//stages
	// 1) run hms_tool
	//2)run pprof cpu start
	// 3)Configure hms with config
	//4) build hms binary
	// 5)start HMS
	// 6) run event generator

	//kill the process
	//loop again

}
