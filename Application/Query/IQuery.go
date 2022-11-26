package Query

import "context"

type IQuery interface {
	Handle(ctx context.Context) (any, error)
}
