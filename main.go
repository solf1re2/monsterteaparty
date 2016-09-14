package main

import (
	"time"

	"github.com/solf1re2/monsterteaparty/gametime"
)

var startTime int64

// main start
func main() {
	var maintimemanager = gametime.Init(time.Now().Unix())
	time.Sleep(10000 * time.Millisecond)
	gametime.GetTimeElapsed(maintimemanager)
}
