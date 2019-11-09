// Package ticker は問題のある sync.Ticker をラップし周期実行作業を改善します。
// sync.Ticker 単体では Stop() で停止しようとしても周期的なチャネル送信を止めるだけなので、goroutine は停止せずにリークします。
// ticker は周期のタイミングでのチャネルからの受信などをユーザロジックと分離し、Stop() を拡張して goroutine も停止させます。
// ユーザロジックは Logic インタフェースとして実装してください。
// ユーザロジックは最初に Before() が、一回の周期ごとに Run() が、Stop() された時に After() が呼び出されます。
// ユーザロジックで周期時間よりもかかった場合には、衝突した周期でのユーザロジックの呼び出しはキャンセルされます。
// キャンセルですので次回の周期には再び呼び出しが試行されます。
package ticker

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xorvercom/util/pkg/easywork"
)

// Logic は周期実行するユーザロジックを規定します。
type Logic interface {
	// 事前処理
	Before()
	// ユーザロジック
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
	// ユーザロジック
	logic Logic
}

// 周期的にユーザロジックを呼ぶ
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

	// ユーザロジックの実行中
	var mutex sync.Mutex
	// 実行中フラグ。boolはbyteでアトミックだと思いたいけど、至る所でスレッドセーフではないと書かれているので。
	running := int32(0)

	// 周期処理
	for {
		select {
		case <-w.cancelCtx.Done():
			// <-w.ticker.C の停止要求
			w.ticker.Stop()
			// ユーザロジックが実行中だった場合には終了を待つ
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
					// ユーザロジックでのパニックは無視
					// 捕捉したいならユーザロジック側でrecover()する
					_ = recover()
					atomic.StoreInt32(&running, 0)
					mutex.Unlock()
				}()
				// ユーザロジック実行
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
