// Package easywork は行うべき仕事のサポートを主眼とした sync.WaitGroup のラッパーです。
// sync.WaitGroup の操作は本来のロジックとは関係ありません。
// easywork は sync.WaitGroup の操作とユーザロジックを分離します。
// ユーザロジックは Running インタフェースとして実装してください。
package easywork

import (
	"sync"
)

// Runnable は非同期実行する仕事です。
type Runnable interface {
	// Run にはユーザロジックを実装します。
	Run()
}

// WaitGroup は仕事(Runnable)の完了を待ちます。
type WaitGroup interface {
	// Start は Runnable を実行します
	Start(Runnable)
	// Wait は実行した Runnable が終了するのを待ちます
	Wait()
	// Results は実行した Runnable の一覧を返します
	Results() []*Work
	// Results は panic を起こした Runnable の一覧を返します
	Panics() []*Work
}

// Work は一つの仕事(Runnable)を格納します。
type Work struct {
	// Runnable の格納用
	Instance Runnable
	// recover() の結果
	Result interface{}
}

// easyWait は sync.WaitGroup のラッパーです。
type easyWait struct {
	// sync.WaitGroup
	wg sync.WaitGroup
	// easyWait の排他制御
	mutex sync.Mutex
	// 仕事の一覧
	works []*Work
	// WaitGroup の完了状態
	done bool
}

// NewGroup は新しい WaitGroup を生成します。
func NewGroup() WaitGroup {
	return &easyWait{works: []*Work{}}
}

const (
	startAfterWaitPanic = "Start() after Wait()"
)

// Start は仕事を開始します。
func (eg *easyWait) Start(ew Runnable) {
	// eg の排他
	eg.mutex.Lock()
	defer eg.mutex.Unlock()
	if eg.done {
		panic(startAfterWaitPanic)
	}

	// sync.WaitGroup に登録
	eg.wg.Add(1)
	wk := &Work{Instance: ew}
	eg.works = append(eg.works, wk)
	go func(wk *Work) {
		defer func(wk *Work) {
			// panic() の捕捉
			wk.Result = recover()
			// 完了
			eg.wg.Done()
		}(wk)
		// 仕事のユーザロジックの実行
		ew.Run()
	}(wk)
}

// Wait は Start() した全ての仕事が完了するまで待機します。
func (eg *easyWait) Wait() {
	// 待機
	eg.wg.Wait()

	// eg の排他
	eg.mutex.Lock()
	defer eg.mutex.Unlock()

	eg.done = true
}

// Results は仕事の結果の一覧を返します。
func (eg *easyWait) Results() []*Work {
	// eg の排他
	eg.mutex.Lock()
	defer eg.mutex.Unlock()

	res := []*Work{}
	res = append(res, eg.works...)
	return res
}

// Panics は panic で終了した仕事の結果の一覧を返します。
func (eg *easyWait) Panics() []*Work {
	// eg の排他
	eg.mutex.Lock()
	defer eg.mutex.Unlock()

	res := []*Work{}
	for _, w := range eg.works {
		if w.Result != nil {
			res = append(res, w)
		}
	}
	return res
}
