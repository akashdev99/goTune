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

// func (s ReportStage) waitAndKill(life int) {
// 	time.Sleep(time.Second * time.Duration(life+5))

// 	s.SampleChan <- nil
// 	wg.Done()
// 	fmt.Println("Done and killed")
// }

func (s ReportStage) startHmsTool(life int, interval int) {
	cmd := exec.Command("hms_tool", "-stats", "system,hm", "-format", "csv", "-interval", strconv.Itoa(interval), "-life", strconv.Itoa(life))
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

func (s ReportStage) Sample(life int, interval int) error {
	wg.Add(1)
	go s.startHmsTool(life, interval)
	wg.Wait()
	return nil
}

var ReportInstance = ReportStage{
	SampleChan: make(chan error, 1),
}
