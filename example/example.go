package main

import (
	"fmt"
	"time"

	"github.com/prashantgupta24/activity-tracker/pkg/tracker"
)

func main() {
	fmt.Println("starting activity tracker")

	timeToCheck := 5

	activityTracker := &tracker.Instance{
		TimeToCheck: time.Duration(timeToCheck),
	}
	heartbeatCh, quitActivityTracker := activityTracker.Start()

	timeToKill := time.NewTicker(time.Second * 30)

	for {
		select {
		case heartbeat := <-heartbeatCh:
			if !heartbeat.IsActivity {
				fmt.Printf("no activity detected in the last %v seconds\n\n", int(timeToCheck))
			} else {
				fmt.Printf("activity detected in the last %v seconds. ", int(timeToCheck))
				fmt.Printf("Activity type:\n")
				for activity, time := range heartbeat.Activity {
					fmt.Printf("%v at %v\n", activity.ActivityType, time)
				}
				fmt.Println()
			}
		case <-timeToKill.C:
			fmt.Println("time to kill app")
			quitActivityTracker <- struct{}{}
			return
		}
	}
}