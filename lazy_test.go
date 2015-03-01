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

