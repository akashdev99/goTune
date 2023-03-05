package main

import (
	"fmt"
	"gotune/configmanager"
)

func main() {
	//get user config for perfomance test
	fmt.Println(configmanager.ConfigMgr.GetConfig())

	//loop through the config and run the daemon
	//parallely run the plugins or In built performance tester
	//Finally make reports for each loop/config
}
