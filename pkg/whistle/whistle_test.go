package whistle_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	//	"time"

	"github.com/xorvercom/util/pkg/whistle"
)

func TestRing(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	w := whistle.New()
	wg := &sync.WaitGroup{}
	const gcount = 3
	wg.Add(gcount)
	for i := 0; i < gcount; i++ {
		go func(idx int, wh *whistle.Whistle) {
			fmt.Printf("start %d %#v\n", idx, wh)
			wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case <-wh.Listen():
					wg.Done()
					fmt.Println(idx, "receive")
					t.Log(idx, "receive")
				}
			}
		}(i, w.Child())
	}
	wg.Wait()
	fmt.Println("running")
	for i := 0; i < 5; i++ {
		fmt.Println("Whistle Ring", i)
		wg.Add(gcount)
		w.Ring()
		wg.Wait()
		fmt.Println("Whistle Ringed", i)
	}
	w.Quit()
}
