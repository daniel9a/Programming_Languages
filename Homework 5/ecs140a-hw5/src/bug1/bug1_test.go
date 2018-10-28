package bug1

import (
	"testing"
	"sync"
	//"time"
)

func sumUp(n1, n2 int) Counter {
    var wg sync.WaitGroup
    counter := Counter(0)
    for i := 0; i < n1; i++ {
        wg.Add(1)
        go func() {
            for i := 0; i < n2; i++ {
                counter.Add(1)
            }
            wg.Done()
        }()
    }
    wg.Wait()
    return counter
}
/*
func sumUp(n1, n2 int) Counter {
	counter := Counter(0)
	for i := 0; i < n1; i++ {
		go func() {
			for i := 0; i < n2; i++ {
				counter.Add(1)
			}
		}()
	}

	time.Sleep(time.Second)
	return counter
}
*/
func TestBug1(t *testing.T) {
	tests := []struct {
		n1   int
		n2   int
		want Counter
	}{
		{100, 1000, 100 * 1000},
		{500, 4000, 500 * 4000},
	}
	for _, test := range tests {
		got := sumUp(test.n1, test.n2)
		if test.want != got {
			t.Errorf("bug1 failed on (%d, %d); want %d, got %d", test.n1, test.n2, test.want, got)
		}
	}
}
