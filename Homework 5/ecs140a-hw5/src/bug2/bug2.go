package bug2

import (
	"fmt"
	"sync"
)

func bug2(n int, foo func()) {
	var wg sync.WaitGroup
//	wg.Add(n)
	for i := 0; i < n; i++ {
		fmt.Println(i)
		wg.Add(1)
		go func() {
		//	fmt.Println(i)
			foo()
			wg.Done()
		}()
	}
	wg.Wait()
}
