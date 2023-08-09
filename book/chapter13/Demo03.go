package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func Factorial2(n uint64) (result uint64) {
	if n > 0 {
		return n * Factorial2(n-1)
	}
	return 1
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	go compute()
	time.Sleep(10 * time.Second)

}

func compute() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(Factorial2(uint64(40)))
		}()
		time.Sleep(time.Millisecond)
	}
}
