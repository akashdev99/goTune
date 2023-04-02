package report

import (
	"log"
	"os/exec"
	"strconv"
	"time"
)

type ReportStage string

func startHmsTool(life int, interval int) {

	cmd := exec.Command("hms_tool", "-stats", "system,hm", "-format", "csv", "-interval", strconv.Itoa(interval), "-life", strconv.Itoa(life))

	if err := cmd.Run(); err != nil {
		log.Fatal("failed to run hms_tool", err)
	}
}

func (s ReportStage) Sample(life int, interval int) error {
	go startHmsTool(life, interval)
	time.Sleep(time.Second * time.Duration(life+5))

	return nil
}

var ReportInstance ReportStage
