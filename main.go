package main

import (
	"fmt"
	"time"

	"github.com/solf1re2/monsterteaparty/console"
	"github.com/solf1re2/monsterteaparty/gametime"
)

var startTime int64
var maintimemanager gametime.Manager

// main start
func main() {
	//startGame()
	console.StartConsole()
	//accessStartTime()
}

func startGame() {
	maintimemanager = gametime.Init(time.Now().Unix())
	time.Sleep(10000 * time.Millisecond)
	gametime.GetTimeElapsed(maintimemanager)
}

func accessStartTime() {
	fmt.Println(gametime.GetStartTime(maintimemanager))
}
