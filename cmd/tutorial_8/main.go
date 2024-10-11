package main

// Concurrency
import (
	"fmt"
	"sync"
	"time"
)

var m = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var wg = sync.WaitGroup{}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	// fmt.Printf("\nTotal results are: %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// save(dbData[i])
	// log()
	wg.Done()
}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}
	wg.Done()
}

// func save(result string) {
// 	m.Lock()
// 	results = append(results, result)
// 	m.Unlock()
// }

// func log() {
// 	m.RLock()
// 	fmt.Printf("\nThe current results are: %v", results)
// 	m.RUnlock()
// }
