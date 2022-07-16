package easywork_test

import (
	"testing"

	ew "github.com/xorvercom/util/pkg/easywork"
)

const (
	testPanic = "test"
)

// テスト用の仕事のベース
type workBase struct {
	t *testing.T
	i int
}

// テスト用の仕事１(easywork.Runnable)
type work1 workBase

// work1 のビジネスロジック
func (w *work1) Run() {
	w.t.Logf("%d\n", w.i)
}

// テスト用の仕事２(easywork.Runnable)
type work2 workBase

// work2 のビジネスロジック
func (w *work2) Run() {
	w.t.Logf("%d\n", w.i)
	panic(testPanic)
}

// TestEasyWork は全ての機能をテストします。
func TestEasyWork(t *testing.T) {
	// 生成
	eg := ew.NewGroup()

	// 仕事の開始
	eg.Start(&work1{t: t, i: 1})
	eg.Start(&work2{t: t, i: 2})

	// 終了を待つ
	eg.Wait()

	// 全ての結果を参照
	for n, w := range eg.Results() {
		if w.Result != nil {
			if w.Result != testPanic {
				t.Fail()
			}
			t.Logf("[%d] inst:%v result:%s", n, w.Instance, w.Result)
		} else {
			t.Logf("[%d] inst:%v success", n, w.Instance)
		}
	}
	// panic 終了した結果を参照
	for n, w := range eg.Panics() {
		t.Logf("[%d] inst:%v result:%s", n, w.Instance, w.Result)
	}

	// コーディングミスの検知
	func() {
		defer func() {
			err := recover()
			if err != ew.StartAfterWaitPanic {
				t.Fatalf("got %v\nwant %v", err, "ew.Append()")
			}
		}()
		// ew.Wait() した後に ew.Start() するのはコーディングミス
		eg.Start(&work1{t: t, i: 3})
	}()
}
