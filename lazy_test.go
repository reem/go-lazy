package lazy

import "testing"
import "sync"

type Data struct {
	x int
}

func TestForce(t *testing.T) {
	// It is only legal to access data after thunk.Force has been called.
	data := &Data{0}
	thunk := Defer(func() {
		data.x = 45
	})

	// "Expensive computation run!" will be printed once
	// some time after this.

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		thunk.Force()
		if data.x != 45 {
			t.Fatal("data.x has an unexpected value of:", data.x)
		}
	}()

	go func() {
		defer wg.Done()

		thunk.Force()
		if data.x != 45 {
			t.Fatal("data.x has an unexpected value of:", data.x)
		}
	}()

	wg.Wait()
}

func TestRunsWorkOnec(t *testing.T) {
	counter := 0
	thunk := Defer(func() {
		counter += 1
	})

	var wg sync.WaitGroup
	wg.Add(10)

	for range make([]struct{}, 10) {
		go func() {
			defer wg.Done()

			thunk.Force()
			thunk.Force()
			thunk.Force()
			thunk.Force()
			thunk.Force()
		}()
	}

	wg.Wait()
	if counter != 1 {
		t.Fatal("The lazy workload was run more than once.")
	}
}
