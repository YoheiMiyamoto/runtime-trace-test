package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	name     string        // taskの名前
	duration time.Duration // task完了までに必要とする時間
	bug      bool          // trueの場合処理時にエラーが発生する
}

// work関数は引数のTaskを実行し、その実行結果を返します
func work(t *Task) error {
	time.Sleep(t.duration) // 擬似的に待機させます。

	// taskにbugが含まれている場合はエラーを発生させます
	if t.bug {
		return fmt.Errorf("err %s", t.name)
	}

	fmt.Fprintf(os.Stdout, "done %s\n", t.name)
	return nil
}
