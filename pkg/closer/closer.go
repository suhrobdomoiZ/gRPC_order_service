package closer

import (
	"context"
	"sync"
)

type closeFn func(ctx context.Context) error

type item struct {
	name string
	fn   closeFn
}

type Closer struct {
	mu    sync.Mutex
	items map[string]item
}
