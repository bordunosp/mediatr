package Infrastructure

import (
	"context"
	"github.com/bordunosp/mediatr/Application/Event"
	"sync"
)

var _lockEventBus = &sync.Mutex{}
var _eventBus Event.IEventBus

func NewEventBus() Event.IEventBus {
	if _eventBus == nil {
		_lockEventBus.Lock()
		defer _lockEventBus.Unlock()

		if _eventBus == nil {
			_eventBus = &eventBus{}
		}
	}

	return _eventBus
}

type eventBus struct{}

func (bus *eventBus) Dispatch(ctx context.Context, event Event.IEvent) (err error) {
	defer func() {
		if r := RecoverToError(recover()); r != nil {
			err = r
		}
	}()

	err = event.Dispatch(ctx)
	return
}

func (bus *eventBus) DispatchAsync(ctx context.Context, event Event.IEvent) chan error {
	c := make(chan error)

	go func(ctx context.Context, event Event.IEvent) {
		defer close(c)
		c <- event.Dispatch(ctx)
	}(ctx, event)

	return c
}

func (bus *eventBus) DispatchAsyncAwait(ctx context.Context, event Event.IEvent) error {
	return <-bus.DispatchAsync(ctx, event)
}
