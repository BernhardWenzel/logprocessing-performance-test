package main

import (
	"time"
	"log"
	"fmt"
	"bufio"
	"os"
	"regexp"
)

//var regex = regexp.MustCompile(".*#[15][15]1110+$")
var regex = regexp.MustCompile("^[^#]*#[15][15]1110+$")


func main() {
	start := time.Now()
	log.Printf("Started: %v", start)
	inputFile, err := os.Open("log-sample.txt")
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer inputFile.Close()
	reader := bufio.NewReader(inputFile)
	for line, _, err := reader.ReadLine(); err==nil;
		line, _, err = reader.ReadLine(){
		if regex.Match(line) {
			fmt.Printf("Match: %v\n", string(line))
		}
	}
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}