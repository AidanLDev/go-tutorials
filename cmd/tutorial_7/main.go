package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var m = sync.Mutex{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 :=
		time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	var delay = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	fmt.Println("The result from the database is:", dbData[i])
	wg.Done()
}
