package ticker

import (
	"fmt"
	"time"
)

// 周期実行（ユーザロジック）
type tick struct {
	// カウンタ
	count int
	// 停止用
	ticker Ticker
}

// 事前処理（ユーザロジック）
func (t *tick) Before() {
	t.count = 0
	fmt.Printf("before count:%d\n", t.count)
}

// 周期実行処理（ユーザロジック）
func (t *tick) Run() {
	t.count++
	fmt.Printf("run count:%d\n", t.count)
	if t.count > 2 {
		// 停止
		t.ticker.Stop()
	}
}

// 事後処理（ユーザロジック）
func (t *tick) After() {
	fmt.Printf("after count:%d\n", t.count)
}

func Example() {
	t := &tick{}
	t.ticker = New()
	t.ticker.Start(t, time.Duration(100)*time.Millisecond)
	t.ticker.Wait()
	fmt.Println("done")
	// Output:
	// before count:0
	// run count:1
	// run count:2
	// run count:3
	// after count:3
	// done
}
