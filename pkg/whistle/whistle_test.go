package whistle_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/xorvercom/util/pkg/whistle"
)

func TestRing(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wb := whistle.New()
	defer wb.Quit()
	w := wb.Child()
	defer w.Quit()
	wdummy := wb.Child()
	defer wdummy.Quit()
	wg := &sync.WaitGroup{}
	const gcount = 3
	wg.Add(gcount)
	for i := 0; i < gcount; i++ {
		go func(idx int, wh *whistle.Whistle) {
			// fmt.Printf("start %d %#v\n", idx, wh)
			wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case <-wh.Listen():
					wg.Done()
					// fmt.Println(idx, "receive")
					// t.Log(idx, "receive")
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
}

func TestChild(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	w := whistle.New()
	defer w.Quit()
	wg := &sync.WaitGroup{}
	const gcount1 = 2
	const gcount2 = 3
	const gcount = gcount1 * gcount2
	wg.Add(gcount)
	for i := 0; i < gcount1; i++ {
		go func(i int, w *whistle.Whistle) {
			for j := 0; j < gcount2; j++ {
				go func(i, j int, w *whistle.Whistle) {
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						case <-w.Listen():
							fmt.Println("receive", i, j)
							wg.Done()
						}
					}
				}(i, j, w.Child())
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
}
