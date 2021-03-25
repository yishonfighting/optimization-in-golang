package main

import (
	"os"
	"runtime/pprof"
	"yishon.github.com/optimization-in-golang/dealFunc"
)

func main() {

	//pprof open code
	cpuProfile, _ := os.Create("cpu_profile")
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()


	//待优化的原生func
	dealFunc.Func1()

	//优化一版的func
	//run result:
	//"moby.txt": 428544 words
	//go run main.go moby.txt  0.27s user 0.18s system 70% cpu 0.641 total
	//dealFunc.Func2()

	//优化第二版的func
	//dealFunc.Func3()

}


