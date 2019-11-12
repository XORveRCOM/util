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
	"time"

	"github.com/xorvercom/util/pkg/easywork"
)

// Ticker は周期実行作業をするためのインタフェースです。
type Ticker interface {
	// Start は周期実行作業を開始します。
	// Logic.Run は Duration の周期で定期的に呼ばれる goroutine ですので、長時間の作業は行わないでください。
	// Duration の周期に TickerWork.Run が完了していなければ、その周期での実行はスキップされます。
	Start(Logic, time.Duration)
	// Stop は定期実行処理を停止します。
	Stop()
	// Wait は周期実行の停止が確認されるまで待機します。
	// Ticker が実行しているユーザロジックの中で呼び出すとデッドロックしますので注意してください。
	Wait()
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

// 周期実行を追加します。
func (t *ticker) Start(tl Logic, d time.Duration) {
	// 親子関係があり t.cancelFunc によって子供の ctx もキャンセルされるため fnc は不要
	// しかし ctx, _ := context.WithCancel(t.cancelCtx) だと go vet のバグでコンパイルエラーとなる
	ctx, fnc := context.WithCancel(t.cancelCtx)
	// https://github.com/golang/go/issues/29587 での回避策
	_ = fnc
	tw := &tickerLogic{ticker: time.NewTicker(d), cancelCtx: ctx, logic: tl}
	// easywork.WaitGroup として起動
	t.waitgroup.Start(tw)
}

// 周期実行を停止します。
func (t *ticker) Stop() {
	t.cancelFunc()
	// 念のため
	_ = <-t.cancelCtx.Done()
}

// 周期実行の停止が確認されるまで待機します。
// Ticker が実行しているユーザロジックの中で呼び出すとデッドロックしますので注意してください。
func (t *ticker) Wait() {
	// 終了待ち
	t.waitgroup.Wait()
}
