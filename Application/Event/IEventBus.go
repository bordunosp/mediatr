package Event

import "context"

type IEventBus interface {
	Dispatch(ctx context.Context, event IEvent) error
	DispatchAsync(ctx context.Context, event IEvent) chan error
	DispatchAsyncAwait(ctx context.Context, event IEvent) error
}
