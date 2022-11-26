package Event

import "context"

type IEvent interface {
	Dispatch(ctx context.Context) error
}
