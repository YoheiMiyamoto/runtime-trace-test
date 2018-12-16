package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

type Task struct {
	name     string        // taskの名前
	duration time.Duration // work完了までに必要とする時間
	bug      bool          // trueの場合処理時にエラーが発生する
}

func work(ctx context.Context, t *Task) error {
	errCh := make(chan error, 1)  // 追加
	go func(errCh chan<- error) { // work処理を別のgoroutineで実行。エラーはerrChに送信。
		time.Sleep(t.duration)
		if t.bug {
			errCh <- fmt.Errorf("err %s\n", t.name)
			return
		}
		errCh <- nil
	}(errCh)

	// errCh もしくは、ctx.Done() を受け取るまで待機
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "done %s\n", t.name)
		return nil
	case <-ctx.Done():
		fmt.Fprintf(os.Stdout, "%s 強制終了\n", t.name)
		return ctx.Err()
	}
}
