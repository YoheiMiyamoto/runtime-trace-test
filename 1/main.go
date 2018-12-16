// シンプルな単一処理

package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	// trace処理
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	ts := []*Task{
		&Task{"one", 1 * time.Second, false},
		&Task{"two", 2 * time.Second, true},
		&Task{"three", 3 * time.Second, false},
	}

	for _, t := range ts {
		err := work(t)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
