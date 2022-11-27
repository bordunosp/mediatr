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

### Simple Command Example

> First we need to implement the command interface

```go
// github.com/bordunosp/mediatr/Application/Command/ICommand

type ICommand interface {
    Execute(ctx context.Context) error
}
```
this will be our first Command

```go
import "context"

type UserAdd struct {
    Name  string
    Email string
}

func (command *UserAdd) Execute(ctx context.Context) error {
    // save user to database
    return nil
}
```

And after that we can use it with CommandBus
in one of 3 ways convenient for us


```go
package main

import (
    "github.com/bordunosp/mediatr"
)

func main() {
    cnt := context.Background()

    command := &UserAdd{
        Name:  "User Name",
        Email: "user@email.com",
    }

    // 1st - standard way, just do job
    err := mediatr.CommandBus.Execute(cnt, command)
    if err != nil {
        panic(err)
    }

    // 2nd - run job in coroutine, the chan with error will be returned 
    err = <-mediatr.CommandBus.ExecuteAsync(cnt, command)
    if err != nil {
        panic(err)
    }

    // 3rd - run and await job in coroutine
    err = mediatr.CommandBus.ExecuteAsyncAwait(cnt, command)
    if err != nil {
        panic(err)
    }
}
```





### Simple Query Example

> First we need to implement the query interface

```go
// github.com/bordunosp/mediatr/Application/Query/IQuery

type IQuery interface {
    Handle(ctx context.Context) (any, error)
}
```
this will be our first Query

```go
import "context"

type UserByEmail struct {
    Email string
}

func (query *UserByEmail) Handle(ctx context.Context) (any, error) {
    // get user from database by email 'query.Email'
    var user UserDTO
	
    return user, nil
}
```

And after that we can use it with QueryBus
in one of 3 ways convenient for us


```go
package main

import (
    "github.com/bordunosp/mediatr"
    "log"
)

func main() {
    cnt := context.Background()

    query := &UserByEmail{
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




### Simple Event Example

> First we need to implement the event interface

```go
// github.com/bordunosp/mediatr/Application/Event/IEvent

type IEvent interface {
    Dispatch(ctx context.Context) error
}
```
this will be our first Event

```go
import "context"

type UserWasCreated struct {
    Email string
}

func (event *UserWasCreated) Dispatch(ctx context.Context) error {
    // do stuff
    // it could be save log or send letter or something else 
    return nil
}
```

And after that we can use it with EventBus
in one of 3 ways convenient for us


```go
package main

import (
    "github.com/bordunosp/mediatr"
    "log"
)

func main() {
    cnt := context.Background()

    query := &UserWasCreated{
        Email: "user@email.com",
    }

    // 1st - standard way, just do job
    err := mediatr.EventBus.Dispatch(cnt, event)
    if err != nil {
        panic(err)
    }

    // 2nd - run job in coroutine, the chan with error will be returned 
    err = <-mediatr.EventBus.DispatchAsync(cnt, event)
    if err != nil {
        panic(err)
    }

    // 3rd - run and await job in coroutine
    err = mediatr.EventBus.DispatchAsyncAwait(cnt, event)
    if err != nil {
        panic(err)
    }
}
```

> If we need more then one Dispatch per Event
> We just call many events inside first Dispatch event