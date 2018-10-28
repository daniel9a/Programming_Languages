package bug2

func noop() {
}

// https://golang.org/pkg/testing/#hdr-Examples
func ExampleBug2() {
	bug2(5, noop)
	// Unordered Output: 0
	// 1
	// 2
	// 3
	// 4
}
