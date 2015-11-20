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

	reader := bufio.NewReader(inputFile)
	i := 0
	for _, _, err := reader.ReadLine(); err==nil;
	_, _, err = reader.ReadLine(){
		i++;
	}
	fmt.Printf("Count: %v\nElapsed time: %v\n", i, time.Since(start))
}