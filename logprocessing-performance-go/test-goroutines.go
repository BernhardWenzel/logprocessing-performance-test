package main
import (
	"time"
	"log"
	"os"
	"fmt"
	"sync"
	"runtime"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	cpu := runtime.NumCPU()
	log.Printf("Num cpu: %v\n", cpu)
	// 4 => 37s
	// 1 => 1m13
	// 64 => 37s
	runtime.GOMAXPROCS(64)

	start := time.Now()
	log.Printf("Started: %v", start)

	var waitGroup sync.WaitGroup
	concurrentTotal := 4
	waitGroup.Add(concurrentTotal)
	log.Printf("Number of concurrent calls: %v", concurrentTotal)

	for i := 0; i < concurrentTotal; i++ {
		go func (concurrentCount int) {
			defer waitGroup.Done()
			for j:= 0; j < 10000000000; j++ {
				if j % 100000000 == 0 {
					log.Printf("[%v] %v", concurrentCount, j)
				}
			}
			log.Printf("-- DONE Concurrent routine #%v", concurrentCount)
		}(i)
	}
	waitGroup.Wait()
	log.Println("------------------------------")
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}