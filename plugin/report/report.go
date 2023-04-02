package report

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"sync"
)

type ReportStage struct {
	SampleChan chan error
}

var wg sync.WaitGroup

func (s ReportStage) startHmsTool(life int, interval int, config map[string]interface{}) {
	fileName := fmt.Sprintf("HMS_TOOL__PROCCESSOR_%v_BUCKET_%v", config["MAX_CONCURRENT_PROCESSORS"], config["PROCESSOR_BUCKET_SIZE"])

	cmd := exec.Command("hms_tool", "-stats", "system,hm", "-format", "csv", "-interval", strconv.Itoa(interval), "-life", strconv.Itoa(life), "-name", fileName)
	fmt.Println("Sampling started")
	if err := cmd.Run(); err != nil {
		log.Fatal("failed to run hms_tool", err)
		s.SampleChan <- err
		wg.Done()
		return
	}
	s.SampleChan <- nil
	wg.Done()
}

func (s ReportStage) Sample(life int, interval int, config map[string]interface{}) error {
	wg.Add(1)
	go s.startHmsTool(life, interval, config)
	wg.Wait()
	return nil
}

var ReportInstance = ReportStage{
	SampleChan: make(chan error, 1),
}
