package main

import (
	"time"
	"log"
	"os"
	"fmt"
	"bufio"
)

func main() {
	start := time.Now()
	log.Printf("Started: %v", start)
	inputFile, err := os.Open("log-sample.txt")
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		i++
		if line == nil {} // we need to use line variable to not get a compilation error
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Printf("Count: %v\nElapsed time: %v\n", i, time.Since(start))
}