package main

import (
	"log"
	"os"
	"time"
	"fmt"
)

func init() {
	log.SetOutput(os.Stdout)
}

const numberOfLines = 20000000
const logFile = "log-sample.txt"

func main() {
	log.Println("Creating log file: " + logFile)

	fo, err := os.Create(logFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	now := time.Now()
	for i:=0; i < numberOfLines; i++  {
		date := now.Add(time.Duration(i) * time.Second)
		_, err = fo.WriteString(fmt.Sprintf("%v -- %v\n", date, fmt.Sprintf("log entry #%v", i)))
		if err != nil {
			panic(err)
		}
		if i % 100000 == 0 {
			log.Printf("... created %v lines", i)
		}
	}
	log.Printf("Done writing %v lines", logFile)
}