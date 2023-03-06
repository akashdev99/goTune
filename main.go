package main

import "gotune/workflowmanager"

func main() {
	//get user config for perfomance test
	workflowmanager.Start()

	//parallely run the plugins or In built performance tester
	//Finally make reports for each loop/config
}
