package bug1

type Counter int64

func (c *Counter) Add(x int64) {
	*c++
}
