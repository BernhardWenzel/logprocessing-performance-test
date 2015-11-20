package main
import (
	"time"
	"log"
	"os"
	"fmt"
	"runtime"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	// Allocate a logical processor for every available core.
	cpu := runtime.NumCPU()
	log.Printf("Num cpu: %v\n", cpu)
	// 4 => 1m11
	// 1 => 1m12
	runtime.GOMAXPROCS(1)

	start := time.Now()
	log.Printf("Started: %v", start)

	concurrentTotal := 4
	log.Printf("Number of concurrent calls: %v", concurrentTotal)

	for i := 0; i < concurrentTotal; i++ {
		func (concurrentCount int) {
			for j:= 0; j < 10000000000; j++ {
				if j % 100000000 == 0 {
					log.Printf("[%v] %v", concurrentCount, j)
				}
			}
			log.Printf("-- DONE Concurrent routine #%v", concurrentCount)
		}(i)
	}
	log.Println("------------------------------")
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}