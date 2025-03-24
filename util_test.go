package blue_green_state_monitor_go

import (
	"sync"
	"time"
)

func waitWG(timeout time.Duration, groups ...*sync.WaitGroup) bool {
	finalWg := &sync.WaitGroup{}
	finalWg.Add(len(groups))
	c := make(chan struct{})
	for _, wg := range groups {
		go func(wg *sync.WaitGroup) {
			wg.Wait()
			finalWg.Done()
		}(wg)
	}
	go func() {
		finalWg.Wait()
		c <- struct{}{}
	}()

	timer := time.NewTimer(timeout)
	select {
	case <-c:
		timer.Stop()
		return true
	case <-timer.C:
		return false
	}
}
