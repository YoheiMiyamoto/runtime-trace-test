// シンプルな単一処理

package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"

	"golang.org/x/sync/errgroup"
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

	var eg errgroup.Group

	for _, t := range ts {
		t := t
		eg.Go(func() error {
			return work(t)
		})
	}

	// すべての並列処理が完了するまで待機
	if err := eg.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}
