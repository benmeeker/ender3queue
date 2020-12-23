package main

import (
    "log"
    "os"
    "time"
)

func nextjob(dirname string) string {
    var timestamp = time.Now().Local().Add(24 * time.Hour).Unix()
    var nextinqueue string
    f, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        if file.ModTime().Unix() < timestamp {
            timestamp = file.ModTime().Unix()
            nextinqueue = file.Name()
        }
    }
    if nextinqueue == ""{
        return nextinqueue
    }
    return dirname+nextinqueue
}