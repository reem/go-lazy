# Lazy [![Build Status](https://travis-ci.org/reem/go-future.svg?branch=master)](https://travis-ci.org/reem/go-future)

> A Lazy type for synchronization of lazily-evaluated data.

Lazy controls a 0-argument function which can be used to initialize a
captured pointer. To work around the lack of generics, Lazy can "control"
this accompanying pointer, like Mutex.

## Example

```go
package main

import lazy "github.com/reem/go-lazy"
import "fmt"

type Data struct {
    int x
}

func main() {
    // It is only legal to access data after thunk.Force has been called.
    data := &Data{0}
    thunk := lazy.Defer(func() {
        fmt.Println("Expensive computation run!")
        data.x = 45
    })

    // "Expensive computation run!" will be printed once
    // some time after this.

    go func() {
        thunk.Force()
        fmt.Println("data.x:", data.x)
    }()

    go func() {
        thunk.Force()
        fmt.Println("data.x:", data.x)
    }()
}
```

## Author

[Jonathan Reem](https://medium.com/@jreem) is the primary author and maintainer of future.

## License

MIT

