package main

import (
	"log"
	"os"
	"time"
	"fmt"
	"github.com/bernhardwenzel/logprocessing-performance-go/logfile"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("Creating log file: " + logfile.LogFile)

	fo, err := os.Create(logfile.LogFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	now := time.Now()
	for i:=0; i < logfile.LogFileNumberOfLines; i++  {
		date := now.Add(time.Duration(i) * time.Second)
		_, err = fo.WriteString(fmt.Sprintf("%v -- %v\n", date, fmt.Sprintf("log entry #%v", i)))
		if err != nil {
			panic(err)
		}
		if i % 100000 == 0 {
			log.Printf("... created %v lines", i)
		}
	}
	log.Printf("Done writing %v lines", logfile.LogFileNumberOfLines)
}