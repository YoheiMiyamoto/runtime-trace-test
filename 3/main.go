package main

import (
	"context"
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
	eg, ctx := errgroup.WithContext(context.Background()) // 追加

	for _, t := range ts {
		t := t
		eg.Go(func() error {
			err := work(ctx, t) // ctxを追加
			if err != nil {
				return err
			}
			return nil
		})
	}

	// すべての並列処理が完了するまで待機
	err = eg.Wait()
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
	}
}
