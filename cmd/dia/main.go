package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/doctordesh/diagnostics"
)

func main() {
	diagnostics.LogMemoryAtInterval(time.Millisecond * 250)

	a := [][]int{}
	for i := 0; i < 100; i++ {
		b := make([]int, 1024*1024)
		a = append(a, b)
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("Done allocating")

	a[0] = nil
	a[1] = nil
	a[2] = nil
	a[3] = nil
	a[4] = nil
	a[5] = nil
	a[6] = nil
	a[7] = nil
	a[8] = nil
	a[9] = nil

	runtime.GC()

	time.Sleep(time.Second * 5)

	runtime.GC()

	time.Sleep(time.Second * 5)
}
