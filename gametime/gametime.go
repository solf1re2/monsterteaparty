package gametime

import (
	"fmt"
	"time"
)

type Manager struct {
	startTime int64
}

// Init - initiate a gametime.Manager
func Init(t int64) Manager {
	return Manager{startTime: t}
}

// SetStartTime - set the startTime variable.
func SetStartTime(t Manager) {
	t.startTime = time.Now().Unix()
}

func GetStartTime(t Manager) int64 {
	return t.startTime
}

// GetTimeElapsed - get the time from start of program
func GetTimeElapsed(t Manager) {
	var currentTime = time.Now().Unix()
	fmt.Println("Start Time: ", t.startTime)
	fmt.Println("Current Time: ", currentTime)
	fmt.Println("Time Difference (secs): ", currentTime-t.startTime)
}
