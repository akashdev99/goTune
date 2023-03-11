package main

import "fmt"

type SetupStage string

func (s SetupStage) Setup() {
	fmt.Println("setup done")

	//loop through the config and run the daemon
	//stages
	// 1) run hms_tool
	//2)run pprof cpu start
	// 3)Configure hms with config
	//4) build hms binary
	// 5)start HMS
}

func (s SetupStage) Description() {
	fmt.Println("the plugin will setup the stress system for perfromance testing")
}

var Setup SetupStage
