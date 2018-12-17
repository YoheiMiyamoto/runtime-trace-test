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

// work関数は引数のTaskを実行し、その実行結果を返します
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
			return err // ここでエラーが起きたことはctx.Done()経由で他のgoroutineに送信される
		}
		fmt.Fprintf(os.Stdout, "done %s\n", t.name)
		return nil
	case <-ctx.Done(): // 他のGoroutineからキャンセルのための終了通知を受け取った場合
		fmt.Fprintf(os.Stdout, "%s キャンセル\n", t.name)
		return ctx.Err()
	}
}
