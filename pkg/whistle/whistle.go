package whistle

import (
	"sync"
)

// Whistle は Ring() メソッドを実行すると Child() で生成していた子供の Listen() の戻すチャネルに送信する同期オブジェクトです。
type Whistle struct {
	recv   chan int
	done   chan struct{}
	childs []*Whistle
	parent *Whistle
	mu     sync.RWMutex
}

const (
	recv_ring = 0
	recv_quit = 1
)

// 新規に Whistle を生成します。
func New() *Whistle {
	return new(nil)
}
func new(p *Whistle) *Whistle {
	w := &Whistle{parent: p}
	w.recv = make(chan int, 1)
	w.done = make(chan struct{}, 1)
	go w.run()
	return w
}

func (w *Whistle) run() {
	for {
		r := <-w.recv
		switch r {
		case recv_ring:
			// Listen() で渡したチャネルに通知
			if len(w.done) == 0 {
				w.done <- struct{}{}
			}
			// 更に子供に転送
			w.Ring()
		case recv_quit:
			// 子供に転送
			w.send(recv_quit)
			// 常駐解除
			return
		}
	}
}

// 親 Whistle で Ring() されたことを通知してくるチャネルを返します。
func (w *Whistle) Listen() <-chan struct{} {
	return w.done
}

// 子 Whistle を生成します。
func (w *Whistle) Child() *Whistle {
	child := new(w)
	w.mu.Lock()
	defer w.mu.Unlock()
	w.childs = append(w.childs, child)
	return child
}

func (w *Whistle) send(code int) {
	if w.childs == nil {
		return
	}
	w.mu.RLock()
	defer w.mu.RUnlock()
	for _, child := range w.childs {
		if len(child.recv) == 0 {
			child.recv <- code
		}
	}
}

// 子 Whistle に一斉通知します。
func (w *Whistle) Ring() {
	w.send(recv_ring)
}

// 子 Whistle に停止を一斉通知します。
func (w *Whistle) Quit() {
	// 親から自身の参照を除去
	if w.parent != nil {
		w.parent.mu.Lock()
		childs := []*Whistle{}
		for _, val := range w.parent.childs {
			if val != w {
				childs = append(childs, val)
			}
		}
		w.parent.childs = childs
		w.parent.mu.Unlock()
	}
	w.mu.Lock()
	w.childs = nil
	w.mu.Unlock()
	// 自身に停止を指示
	if len(w.recv) == 0 {
		w.recv <- recv_quit
	}
}
