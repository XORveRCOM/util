package ticker

import (
	"fmt"
	"testing"
	"time"
)

// 周期実行
type worker struct {
	id  int
	num int
	sec int
}

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

// ビジネスロジック
func (w *worker) Before() {
	fmt.Println("[事前処理]", w, time.Now())
}
func (w *worker) Run() {
	w.num++
	fmt.Println("[start]", w, time.Now())
	time.Sleep(time.Duration(w.sec) * time.Millisecond)
	fmt.Println("[end]", w, time.Now())
}
func (w *worker) After() {
	fmt.Println("[事後処理]", w, time.Now())
}
func (w *worker) String() string {
	return fmt.Sprintf("&{id:%d, num:%d, sec:%d}", w.id, w.num, w.sec)
}

// TestNew は作成した直後にStopを呼びます。
func TestNew(t *testing.T) {
	tick := New()
	tick.Stop()
	tick.Wait()
	t.Log()
}

// 通常テスト
func TestStart(t *testing.T) {
	tick := New()
	tick.Start(&worker{id: 1, num: 0, sec: 500}, 300*time.Millisecond)
	tick.Start(&worker{id: 2, num: 0, sec: 500}, 500*time.Millisecond)
	tick.Start(&worker{id: 3, num: 0, sec: 500}, 600*time.Millisecond)
	time.Sleep(5 * time.Second)
	fmt.Println("tickStop()", time.Now())
	tick.Stop()
	tick.Wait()
	fmt.Println("exit", time.Now())
	t.Log()
}
