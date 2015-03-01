package lazy

import oncemutex "github.com/reem/go-once-mutex"

const (
	start      uint32 = 0
	evaluating uint32 = 1
	complete   uint32 = 2
)

// A type which lazily evaluates a workload, represented
// as a 0-argument function.
//
// Constructed using Defer
type Lazy struct {
	state uint32
	mu    *oncemutex.OnceMutex
	work  func()
}

// Create a new instance of Lazy which manages a func() workload.
func Defer(work func()) *Lazy {
	return &Lazy{start, oncemutex.NewOnceMutex(), work}
}

// Force an existing Lazy, this blocks until the workload has been
// finished at least once.
func (l *Lazy) Force() {
	defer l.mu.Unlock()

	if !l.mu.Lock() {
		// We were never locked before, so can mutate.
		l.work()

		// Cleanup to allow GC.
		l.work = nil
	}
}
