package Command

import "context"

type ICommandBus interface {
	Execute(ctx context.Context, command ICommand) error
	ExecuteAsync(ctx context.Context, command ICommand) chan error
	ExecuteAsyncAwait(ctx context.Context, command ICommand) error
}
