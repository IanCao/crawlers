package main

import (
	"crawlers/crawler-core/core"
	"sync"
)

func main() {
	maxConcurrentCount := 2
	wg := sync.WaitGroup{}
	wg.Add(maxConcurrentCount)
	go func() {
		defer wg.Done()
		core.ProcessBeiKeErShouFang()
	}()
	go func() {
		defer wg.Done()
		go core.ProcessBeiKeZuFang()
	}()
	wg.Wait()
}
