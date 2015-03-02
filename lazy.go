package lazy

import "sync"

// A type which lazily evaluates a workload, represented
// as a 0-argument function.
//
// Constructed using Defer
type Lazy struct {
	once  *sync.Once
	work  func()
}

// Create a new instance of Lazy which manages a func() workload.
func Defer(work func()) *Lazy {
	return &Lazy{&sync.Once{}, work}
}

// Force an existing Lazy, this blocks until the workload has been
// finished at least once.
func (l *Lazy) Force() {
    l.once.Do(l.work)
}

