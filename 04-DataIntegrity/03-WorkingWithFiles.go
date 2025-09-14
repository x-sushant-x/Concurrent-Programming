package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var writer chan bool
var rwLock sync.RWMutex

func writeFile(i int) {
	rwLock.RLock()

	os.WriteFile("test.txt", []byte(strconv.FormatInt(int64(i), 10)), 0x777)
	rwLock.RUnlock()

	writer <- true
}

func writeMain() {
	writer = make(chan bool)

	for i := range 10 {
		go writeFile(i)
	}

	<-writer

	fmt.Println("Done")
}
