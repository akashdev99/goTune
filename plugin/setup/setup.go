package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type SetupStage string

func makeConfiguration(config map[string]interface{}, dir string) error {
	var temp string
	file, err := os.Open(dir)
	if err != nil {
		log.Fatalf("Could not open file :%v", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		contains := false
		for property, value := range config {
			if strings.Contains(data, property) {
				contains = true
				temp = temp + fmt.Sprintf("%v = %v\n", property, value)
			}
		}

		if !contains {
			temp = temp + data + "\n"
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Something went wrong :", err)
		return err
	}

	f, err := os.OpenFile(dir, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal("Failed to open file for write :", err)
		return err
	}

	f.WriteString(temp)

	if err := f.Close(); err != nil {
		log.Fatal("Fialed to close file after write :", err)
		return err
	}
	return nil
}

func reloadHms() error {
	cmd := exec.Command("pmtool restartbyid hms")

	if err := cmd.Run(); err != nil {
		log.Fatal("hms restart failed")
		return err
	}
	return nil
}

func (s SetupStage) Setup(config map[string]interface{}, dir string) error {
	// 3)Configure hms with config
	if err := makeConfiguration(config, dir); err != nil {
		return err
	}

	if err := reloadHms(); err != nil {
		return err
	}
	//4) build hms binary

	//MOVE TO REPORTER
	// 1) run hms_tool
	//2)run pprof cpu start
	// 5)start HMS

	// fmt.Println("setup done")

	//loop through the config and run the daemon
	//stages
	return nil
}

func (s SetupStage) Description() {
	fmt.Println("the plugin will setup the stress system for perfromance testing")
}

var Setup SetupStage
