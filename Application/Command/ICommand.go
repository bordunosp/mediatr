package Command

import "context"

type ICommand interface {
	Execute(ctx context.Context) error
}
