package main
import (
	"time"
	"log"
	"os"
	"fmt"
	"bufio"
	"sync"
	"regexp"
	//"runtime"
)

const numberWorkers = 32
var waitGroup sync.WaitGroup
var regex = regexp.MustCompile("^[^#]*#[15][15]1110+$")

func main() {
	//cpu := runtime.NumCPU()
	//log.Printf("Num cpu: %v\n", cpu)
	//runtime.GOMAXPROCS(cpu)

	start := time.Now()
	log.Printf("Started: %v", start)
	inputFile, err := os.Open("log-sample.txt")
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer inputFile.Close()

	waitGroup.Add(numberWorkers)

	queue := make(chan []byte)
	for gr := 1; gr <= numberWorkers; gr++ {
		go worker(queue, gr)
	}

	reader := bufio.NewReader(inputFile)
	for line, _, err := reader.ReadLine(); err==nil;
	line, _, err = reader.ReadLine(){
		queue <- line
	}

	close(queue)
	waitGroup.Wait()
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}

func worker(queue chan []byte, id int) {
	defer waitGroup.Done()
	for {
		line, ok := <-queue
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", id)
			return
		}
		if regex.Match(line) {
			fmt.Printf("[%v] Match: %v\n", id, string(line))
		}
	}
}




