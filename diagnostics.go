package diagnostics

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

func MemoryAtInterval(t time.Duration) chan runtime.MemStats {
	var m runtime.MemStats
	c := make(chan runtime.MemStats)

	go func() {
		for range time.Tick(t) {
			runtime.ReadMemStats(&m)

			select {
			case c <- m:
				break
			default:
				break

			}
		}
	}()

	return c
}

func LogMemoryAtInterval(t time.Duration) {
	c := MemoryAtInterval(t)

	go func() {
		var m runtime.MemStats
		for {
			m = <-c

			log.Printf(strings.Join([]string{
				"Alloc: %v",
				"TotalAlloc: %v",
				"Sys: %v",
				"NumGC: %v",
			}, " "), formatBytes(m.Alloc), formatBytes(m.TotalAlloc), formatBytes(m.Sys), m.NumGC)
		}
	}()
}

func formatBytes(b uint64) string {
	i := 0
	for {
		if b < 1024 {
			break
		}
		i++
		b = b / 1024
	}

	sizes := []string{
		"B",
		"KiB",
		"MiB",
		"GiB",
	}

	return fmt.Sprintf("%d %s", b, sizes[i])
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
