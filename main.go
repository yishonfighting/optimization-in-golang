package main

import "yishon.github.com/optimization-in-golang/dealFunc"

func main() {


	//待优化的原生func
	// run result:
	//"moby.txt": 428544 words
	//go run main.go moby.txt  1.02s user 1.17s system 106% cpu 2.060 total
	dealFunc.Func1()

	//优化一版的func
	//run result:
	//"moby.txt": 428544 words
	//go run main.go moby.txt  0.27s user 0.18s system 70% cpu 0.641 total
	dealFunc.Func2()

	//优化第二版的func
	//dealFunc.Func3()

}


