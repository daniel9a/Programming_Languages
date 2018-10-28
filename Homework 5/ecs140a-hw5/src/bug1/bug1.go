
package bug1

import (
	"sync"
)

type Counter int64 
var mutex = &sync.Mutex{}

func (c *Counter) Add(x int64) {
    mutex.Lock()
	*c++
    mutex.Unlock()
}
