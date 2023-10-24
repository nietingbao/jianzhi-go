package listnode

import (
	"fmt"
	"sync"
	"time"
)

func PrintNumber() {
	ch := make(chan struct{}, 1)
	i := 0
	ch <- struct{}{}
	print := func() {
		for {
			if i > 100 {
				return
			}
			<-ch
			fmt.Println(i)
			i++
			ch <- struct{}{}
		}
	}

	go print()
	go print()
	time.Sleep(time.Minute)
}

func LimitRoutine(limit int) {
	ch := make(chan struct{}, limit)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		i := i
		wg.Add(1)
		ch <- struct{}{}
		go func() {
			defer func() {
				<-ch
				wg.Done()
			}()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
