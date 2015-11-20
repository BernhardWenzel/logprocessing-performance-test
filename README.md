# logprocessing-performance-test

Comparing the performance of Go with Java when filtering the lines of a large log file using regular expressions.

The results:

Testing machine: MacBook Pro, Intel Core i5, 2.4GHz, 2 cores, 8GB Ram

Input file: plain-text log file with 20.000.000 lines (1.24GB)

| Go implementation | Settings  | Execution time  | 
|--:|---|---|
| Looping with counter  | n/a  	|  1.8s |
| Sequential processing  | n/a  |  1m2s |
| Parallel unbuffered channel 	| number of goroutines: 4  	| 54s  |
| Parallel unbuffered channel 	| number of goroutines: 8  	| 37s  |
| Parallel unbuffered channel	| number of goroutines: 16  | 33s  |
| Parallel unbuffered channel	| number of goroutines: 32  | 31s  |
| Parallel unbuffered channel	| number of goroutines: 64  | 31s  |
| Parallel unbuffered channel	| number of goroutines: 128 | 31s  |


| Java implementation | Settings  | Execution time  | 
|--:|---|---|
| Looping with counter  | n/a  	|  5s |
| Sequential processing  | n/a  |  18s |
| Parallel BlockingQueue 	| number of threads: 2  | 18s  |
| Parallel BlockingQueue	| number of threads: 4  | 22s  |
| Parallel BlockingQueue	| number of threads: 8  | 26s  |

Read about the performance test here: <>