package whistle

import (
	"sync"
	"sync/atomic"
)

type Whistle struct {
	recv   chan int
	ringed atomic.Value
	childs []*Whistle
	mu     sync.RWMutex
}

const (
	recv_ring = 0
	recv_quit = 1
)

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})
var openedchan = make(chan struct{})

func init() {
	close(closedchan)
}

func New() *Whistle {
	w := &Whistle{}
	w.ringed.Store(openedchan)
	w.recv = make(chan int)
	go w.run()
	return w
}

func (w *Whistle) run() {
	for {
		r := <-w.recv
		switch r {
		case recv_ring:
			w.ringed.Store(closedchan)
			w.Ring()
		case recv_quit:
			return
		}
	}
}

func (w *Whistle) Listen() <-chan struct{} {
	return w.ringed.Swap(openedchan).(chan struct{})
}

func (w *Whistle) Child() *Whistle {
	child := New()
	w.mu.Lock()
	defer w.mu.Unlock()
	w.childs = append(w.childs, child)
	return child
}

func (w *Whistle) send(code int) {
	for _, child := range w.childs {
		child.recv <- code
		// // non block
		// select {
		// case child.recv <- code:
		// default:
		// 	// The channel is blocked for some reason.
		// }
	}
}

func (w *Whistle) Ring() {
	w.mu.RLock()
	defer w.mu.RUnlock()
	w.send(recv_ring)
}

func (w *Whistle) Quit() {
	w.mu.RLock()
	defer w.mu.RUnlock()
	w.send(recv_quit)
	w.recv <- recv_quit
	// // non block
	// select {
	// case w.recv <- recv_quit:
	// default:
	// 	// The channel is blocked for some reason.
	// }
}
