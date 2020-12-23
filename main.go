package main

import (
	"time"
)

func main() {
	go button()
	for {
		if ready {
			nextinqueue := nextjob("/home/pi/queue/")
			if nextinqueue != ""{
			sendfile(nextinqueue)
			ready = false
			}
		}
		time.Sleep(15 * time.Second)
	}
}