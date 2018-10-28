package bug3

// Merge sends the output from channels ch1 and ch2 to out.
func merge(ch1, ch2, out chan uint) {
	defer close(out)

	for {
		select {
		case x := <-ch1: // Either input from ch1
			out <- x
		case x := <-ch2: // or input from ch2
			out <- x
		}
	}
}

func bug3(producer1 func(chan uint), producer2 func(chan uint)) chan uint {
	// Create channels for producers
	ch1, ch2 := make(chan uint), make(chan uint)
	// Create channel for consumer
	out := make(chan uint)

	// Spawn each goroutine
	go producer1(ch1)
	go producer2(ch2)
	go merge(ch1, ch2, out)

    res := <-out
	// Return the output channel
	return res
}
