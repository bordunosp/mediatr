package Query

import "context"

type IQueryBus interface {
	Handle(ctx context.Context, query IQuery) (any, error)
	HandleAsync(ctx context.Context, query IQuery) chan ReplayDTO
	HandleAsyncAwait(ctx context.Context, query IQuery) (any, error)
}
