// Package ticker は周期実行作業を実現します。
// sync.Ticker では Stop() で停止しようとしても周期的なチャネル送信を止めるだけなので、goroutine は停止せずにリークします。
// 周期のタイミングでのチャネルからの受信などは Ticker の内部で実装され、周期と周期実行のロジックだけを Logic として指定します。
package ticker

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xorvercom/util/pkg/easywork"
)

// Logic は周期実行するビジネスロジックを規定します。
type Logic interface {
	// 事前処理
	Before()
	// ビジネスロジック
	Run()
	// 事後処理
	After()
}

// Ticker は周期実行作業をするためのインタフェースです。
type Ticker interface {
	// Start は周期実行作業を開始します。
	// TickerWork.Run は Duration の周期で定期的に呼ばれる goroutine ですので、長時間の作業は行わないでください。
	// Duration の周期に TickerWork.Run が完了していなければ、その周期での実行はスキップされます。
	Start(Logic, time.Duration)
	// Stop は定期実行処理を停止して、停止が確認されるまで待機します。
	Stop()
}

// 周期実行作業
type tickerWork struct {
	// 周期
	ticker *time.Ticker
	// キャンセルcontext
	cancelCtx context.Context
	// 停止関数
	cancelFunc context.CancelFunc
	// ビジネスロジック
	logic Logic
}

// 周期的にビジネスロジックを呼ぶ
func (w *tickerWork) Run() {
	// 事前処理
	func() {
		defer func() {
			_ = recover()
		}()
		// 事前処理
		w.logic.Before()
	}()

	// 事後処理
	defer func() {
		defer func() {
			_ = recover()
		}()
		w.logic.After()
	}()

	// ビジネスロジックの実行中
	var mutex sync.Mutex
	// 実行中フラグ。boolはbyteでアトミックだと思いたいけど、至る所でスレッドセーフではないと書かれているので。
	running := int32(0)

	// 周期処理
	for {
		select {
		case <-w.cancelCtx.Done():
			// <-w.ticker.C の停止要求
			w.ticker.Stop()
			// ビジネスロジックが実行中だった場合には終了を待つ
			mutex.Lock()
			mutex.Unlock()
			return
		case <-w.ticker.C:
			if atomic.LoadInt32(&running) != 0 {
				// 実行中だったのでスキップ
				continue
			}
			mutex.Lock()
			atomic.StoreInt32(&running, 1)
			go func() {
				defer func() {
					// ビジネスロジックでのパニックは無視
					// 捕捉したいならビジネスロジック側でrecover()する
					_ = recover()
					atomic.StoreInt32(&running, 0)
					mutex.Unlock()
				}()
				// ビジネスロジック実行
				w.logic.Run()
			}()
		}
	}
}

// Ticker は time.Tick のラッパーです。
// time.Tick は Stop させると ticker はタイマーを止めるだけなので goroutine がリークします。
// 定期実行処理を止めると goroutine も止まる機構を実装します。
type ticker struct {
	// 周期実行の親コンテキスト
	cancelCtx context.Context
	// 停止関数
	cancelFunc context.CancelFunc
	// 待機
	waitgroup easywork.WaitGroup
}

// New は Ticker を作成します。
func New() Ticker {
	ctx, fnc := context.WithCancel(context.Background())
	return &ticker{cancelCtx: ctx, cancelFunc: fnc, waitgroup: easywork.NewGroup()}
}

// 周期実行を追加
func (t *ticker) Start(tl Logic, d time.Duration) {
	// ctx, _ := context.WithCancel(t.cancelCtx) だと go vet のバグなのかコンパイルエラーとなる
	ctx, fnc := context.WithCancel(t.cancelCtx)
	tw := &tickerWork{ticker: time.NewTicker(d), cancelCtx: ctx, cancelFunc: fnc, logic: tl}
	// easywork.WaitGroup として起動
	t.waitgroup.Start(tw)
}

// 周期実行を停止
func (t *ticker) Stop() {
	t.cancelFunc()
	// 念のため
	_ = <-t.cancelCtx.Done()
	// 終了待ち
	t.waitgroup.Wait()
}
