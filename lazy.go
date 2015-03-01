package lazy

import oncemutex "github.com/reem/go-once-mutex"

const (
	start      uint32 = 0
	evaluating uint32 = 1
	complete   uint32 = 2
)

type Lazy struct {
	state uint32
	mu    *oncemutex.OnceMutex
	work  func()
}

func Defer(work func()) *Lazy {
	return &Lazy{start, oncemutex.NewOnceMutex(), work}
}

func (l *Lazy) Force() {
	defer l.mu.Unlock()

	if !l.mu.Lock() {
		// We were never locked before, so can mutate.
		l.work()

		// Cleanup to allow GC.
		l.work = nil
	}
}
