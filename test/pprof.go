package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "/tmp/cpuprofile", "write cpu profile to file")

/**
 * Analyze CPU Profile:
 * 1. generate report to a specific location, like /tmp/cpuprofile
 * 2. run command `go tool pprof /tmp/cpuprofile` [important]
 * 3. check top 10 threads by running `top 10`
 * 4. (pprof) list main.main - to display more info( optional )
 * 5. to generate a svg report
 *
 * Other profiling tools:
 * - CPU profiling( CPU usage, per 100ms )
 * - Memory profiling( memory usage )
 * - Block profiling( analyze deadlock )
 * - Goroutine profiling( goroutine information, scheduling relationship )
 */
func main() {
	fmt.Println("===Performance Profiling Starts===")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()

		var result int
		for i := 0; i < 10000000; i++ {
			result += i
		}
		log.Println("result:", result)
	}
	fmt.Println("===Performance Profiling Ends===")
}
