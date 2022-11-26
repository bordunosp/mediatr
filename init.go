package mediatr

import (
	"github.com/bordunosp/mediatr/Application/Command"
	"github.com/bordunosp/mediatr/Application/Event"
	"github.com/bordunosp/mediatr/Application/Query"
	"github.com/bordunosp/mediatr/Infrastructure"
)

var CommandBus Command.ICommandBus
var QueryBus Query.IQueryBus
var EventBus Event.IEventBus

func init() {
	CommandBus = Infrastructure.NewCommandBus()
	QueryBus = Infrastructure.NewQueryBus()
	EventBus = Infrastructure.NewEventBus()
}
