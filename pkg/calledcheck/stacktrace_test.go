package calledcheck_test

import (
	"testing"

	cc "github.com/xorvercom/util/pkg/calledcheck"
)

func TestStackTrace(t *testing.T) {
	for i, s := range cc.StackTrace() {
		t.Log(i, s)
	}
	var ch chan struct{}
	go func(chan struct{}) {
		for i, s := range cc.StackTrace() {
			t.Log(i, s)
		}
		ch <- struct{}{}
	}(ch)
	<-ch
}
