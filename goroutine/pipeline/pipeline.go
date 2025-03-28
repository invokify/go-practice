package pipeline

import "sync"

func repeatFn[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primeStream := make(chan int)

	go func() {
		defer close(primeStream)

		for {
			select {
			case <-done:
				return
			case randomInt := <-randIntStream:
				if isPrime(randomInt) {
					primeStream <- randomInt
				}
			}
		}
	}()

	return primeStream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	takeStream := make(chan T)

	go func() {
		defer close(takeStream)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-stream:
			}
		}
	}()

	return takeStream
}

func merge[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	mergedStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case mergedStream <- i:
			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(mergedStream)
	}()

	return mergedStream
}
