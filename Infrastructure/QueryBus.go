package Infrastructure

import (
	"context"
	"github.com/bordunosp/mediatr/Application/Query"
	"sync"
)

var _lockQueryBus = &sync.Mutex{}
var _queryBus Query.IQueryBus

func NewQueryBus() Query.IQueryBus {
	if _queryBus == nil {
		_lockQueryBus.Lock()
		defer _lockQueryBus.Unlock()

		if _queryBus == nil {
			_queryBus = &queryBus{}
		}
	}

	return _queryBus
}

type queryBus struct{}

func (bus *queryBus) Handle(ctx context.Context, query Query.IQuery) (value any, err error) {
	defer func() {
		if r := RecoverToError(recover()); r != nil {
			err = r
		}
	}()

	value, err = query.Handle(ctx)
	return
}

func (bus *queryBus) HandleAsync(ctx context.Context, query Query.IQuery) (replay chan Query.ReplayDTO) {
	replay = make(chan Query.ReplayDTO)

	go func(query Query.IQuery) {
		defer close(replay)
		value, err := query.Handle(ctx)

		replay <- *&Query.ReplayDTO{Value: value, Err: err}
	}(query)

	return replay
}

func (bus *queryBus) HandleAsyncAwait(ctx context.Context, query Query.IQuery) (any, error) {
	replay := <-bus.HandleAsync(ctx, query)
	return replay.Value, replay.Err
}
