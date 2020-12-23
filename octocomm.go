package main

import (
	"os/exec"
	"fmt"
	"os"
	"strings"
)

func sendfile(nextinqueue string) {
	cmd := exec.Command("./OctoPrint/bin/octoprint", "client", "upload", "/api/files/local", nextinqueue)
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	if err != nil {
		panic(err)
	}
	os.Rename(nextinqueue, strings.Replace(nextinqueue, "queue", "completed", 1))
}
