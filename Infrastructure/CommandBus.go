package Infrastructure

import (
	"context"
	"github.com/bordunosp/mediatr/Application/Command"
	"sync"
)

var _lockCommandBus = &sync.Mutex{}
var _commandBus Command.ICommandBus

func NewCommandBus() Command.ICommandBus {
	if _commandBus == nil {
		_lockCommandBus.Lock()
		defer _lockCommandBus.Unlock()

		if _commandBus == nil {
			_commandBus = &commandBus{}
		}
	}

	return _commandBus
}

type commandBus struct{}

func (bus *commandBus) Execute(ctx context.Context, command Command.ICommand) (err error) {
	defer func() {
		if r := RecoverToError(recover()); r != nil {
			err = r
		}
	}()

	err = command.Execute(ctx)
	return
}

func (bus *commandBus) ExecuteAsync(ctx context.Context, command Command.ICommand) chan error {
	c := make(chan error)

	go func(ctx context.Context, command Command.ICommand) {
		defer close(c)
		c <- command.Execute(ctx)
	}(ctx, command)

	return c
}

func (bus *commandBus) ExecuteAsyncAwait(ctx context.Context, command Command.ICommand) error {
	return <-bus.ExecuteAsync(ctx, command)
}
