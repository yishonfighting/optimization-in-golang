package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"yishon.github.com/optimization-in-golang/dealFunc"
)

func main() {

	//pprof open code
	//cpuProfile, _ := os.Create("cpu_profile")
	//pprof.StartCPUProfile(cpuProfile)
	//defer pprof.StopCPUProfile()


	// pprof mem
	memProfile, _ := os.Create("mem_profile")
	runtime.MemProfileRate = 1

	defer func() {
	pprof.WriteHeapProfile(memProfile)
	memProfile.Close()
	}()



	//待优化的原生func
	//dealFunc.Func1()

	//优化一版的func
	dealFunc.Func2()

	//优化第二版的func
	//dealFunc.Func3()

}


