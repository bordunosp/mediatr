# MediatR

A very simple package in the go language paradigm.

No magic, no reflection, no code generation, etc...


## Architecture Patterns
* [CQRS](https://martinfowler.com/bliki/CQRS.html): Martin Fowler
* [CQRS](https://learn.microsoft.com/en-us/previous-versions/msp-n-p/jj554200(v=pandp.10)): learn.microsoft
* [CQRS](https://www.redhat.com/architect/illustrated-cqrs): RedHat
* [CQS](https://en.wikipedia.org/wiki/Command%E2%80%93query_separation): wiki




## Installation

```bash
go get -u github.com/bordunosp/mediatr
```

## Getting Started

First we need to implement one of 3 possible interfaces

```go
// github.com/bordunosp/mediatr/Application/Command/ICommand
type ICommand interface {
    Execute(ctx context.Context) error
}

// github.com/bordunosp/mediatr/Application/Query/IQuery
type IQuery interface {
    Handle(ctx context.Context) (any, error)
}

// github.com/bordunosp/mediatr/Application/Event/IEvent
type IEvent interface {
    Dispatch(ctx context.Context) error
}
```

---

#### Command Example

```go
import (
    "context"
    "github.com/google/uuid"
)

type UserAddCommand struct {
    Id    uuid.UUID
    Name  string
    Email string
}

func (command *UserAddCommand) Execute(ctx context.Context) error {
    // save user to database
    return nil
}
```

---

#### Query Example

```go
import "github.com/google/uuid"

type UserDTO struct {
    Id    uuid.UUID
    Name  string
    Email string
}
```

```go
import "context"

type UserByEmailQuery struct {
    Email string
}

func (query *UserByEmailQuery) Handle(ctx context.Context) (any, error) {
    var user UserDTO
    // get user from database by email 'query.Email'
    return user, nil
}
```

---

#### Event Example

```go
import (
    "context"
    "github.com/google/uuid"
)

type UserWasCreatedEvent struct {
    Id uuid.UUID
}

func (event *UserWasCreatedEvent) Dispatch(ctx context.Context) error {
    // do stuff
    // it could be save log or send letter or something else 
    return nil
}
```

---

### Simple CommandBus Example

```go
package main

import (
    "context"
    "github.com/bordunosp/mediatr"
    "github.com/google/uuid"
)

func main() {
    cnt := context.Background()

    command := &UserAddCommand{
        Id:    uuid.New(),
        Name:  "User Name",
        Email: "user@email.com",
    }

    // 1st - standard way, just do job
    _ = mediatr.CommandBus.Execute(cnt, command)

    // 2nd - run job in coroutine, the chan with error will be returned 
    _ = <-mediatr.CommandBus.ExecuteAsync(cnt, command)

    // 3rd - run and await job in coroutine
    _ = mediatr.CommandBus.ExecuteAsyncAwait(cnt, command)
}
```

### Simple QueryBus Example

```go
package main

import (
    "context"
    "github.com/bordunosp/mediatr"
    "log"
)

func main() {
    cnt := context.Background()

    query := &UserByEmailQuery{
        Email: "user@email.com",
    }

    // 1st - standard way, just do job
    user, err := mediatr.QueryBus.Handle(cnt, query)
    if err != nil {
        panic(err)
    }
    log.Println(user.(UserDTO))

    // 2nd - run job in coroutine, 
    // the chan with Query.ReplayDTO will be returned 
    replayDTO := <-mediatr.QueryBus.HandleAsync(cnt, query)
    if replayDTO.Err != nil {
        panic(replayDTO.Err)
    }
    log.Println(replayDTO.Value.(UserDTO))

    // 3rd - run and await job in coroutine
    user, err = mediatr.QueryBus.HandleAsyncAwait(cnt, query)
    if err != nil {
        panic(err)
    }
    log.Println(user.(UserDTO))
}
```

### Simple EventBus Example

```go
package main

import (
    "context"
    "github.com/bordunosp/mediatr"
    "log"
)

func main() {
    cnt := context.Background()

    userId := uuid.New()

    query := &UserWasCreatedEvent{
        Id: userId,
    }

    // 1st - standard way, just do job
    _ = mediatr.EventBus.Dispatch(cnt, event)

    // 2nd - run job in coroutine, the chan with error will be returned 
    _ = <-mediatr.EventBus.DispatchAsync(cnt, event)

    // 3rd - run and await job in coroutine
    _ = mediatr.EventBus.DispatchAsyncAwait(cnt, event)
}
```

> If we need more then one Dispatch per Event
> We just call many events inside first Dispatch event
