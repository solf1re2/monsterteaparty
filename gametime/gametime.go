package gametime

import (
	"fmt"
	"time"
)

type manager struct {
	startTime int64
}

// Init - initiate a gametime.Manager
func Init(t int64) manager {
	return manager{startTime: t}
}

// SetStartTime - set the startTime variable.
func SetStartTime(t manager) {
	t.startTime = time.Now().Unix()
}

// GetTimeElapsed - get the time from start of program
func GetTimeElapsed(t manager) {
	var currentTime = time.Now().Unix()
	fmt.Println("Start Time: ", t.startTime)
	fmt.Println("Current Time: ", currentTime)
	fmt.Println("Time Difference (secs): ", currentTime-t.startTime)
}
