package stress

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

type ReportStage struct {
	StressChan chan error
}

var wg sync.WaitGroup

func (s ReportStage) startHmsEventGenrator() {
	cmd := exec.Command("./event_generator")
	fmt.Println("Stressing started")
	if err := cmd.Run(); err != nil {
		log.Fatal("failed to run hms_tool", err)
		s.StressChan <- err
		wg.Done()
		return
	}
	s.StressChan <- nil
	wg.Done()
}

func (s ReportStage) Stress(life int) error {
	wg.Add(1)
	go s.startHmsEventGenrator()
	wg.Wait()
	fmt.Println("Done stress")
	return nil
}

var ReportInstance = ReportStage{
	StressChan: make(chan error, 1),
}
