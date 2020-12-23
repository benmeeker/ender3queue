package main

import (
	"github.com/nathany/bobblehat/sense/stick"

	"os/signal"
	"os"
	"fmt"
	"flag"
)

var ready bool

func button() {
	var path string

	flag.StringVar(&path, "path", "/dev/input/event0", "path to the event device")

	flag.Parse()

	input, err := stick.Open(path)
	if err != nil {
		fmt.Printf("Unable to open input device: %s\nError: %v\n", path, err)
		os.Exit(1)
	}
	// Set up a signals channel (stop the loop using Ctrl-C)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	for {
		select {
		case <-signals:
			fmt.Println("")

			return
		case e := <-input.Events:

			switch e.Code {
			case stick.Enter:
				print("Enter")
				ready = true
			}
		}
	}
}

