package ticker

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
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

// 周期実行作業
type tickerLogic struct {
	// 周期
	ticker *time.Ticker
	// キャンセルcontext
	cancelCtx context.Context
	// ユーザロジック
	logic Logic
}

// 周期的にユーザロジックを呼ぶ
func (w *tickerLogic) Run() {
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
