package libraries

import "context"

type CommandHandler interface {
	Handle(ctx context.Context, event interface{}) (interface{}, error)
}

type LambdaHandler interface {
	Handler(ctx context.Context, event interface{}) (interface{}, error)
}

type Middleware func(LambdaHandler) LambdaHandler

// Constructor types
type CommandHandlerConstructor[R any] func(repo R) CommandHandler
type HandlerConstructor func(cmd CommandHandler) LambdaHandler

func CreateHandler[R any](
	newCommandHandler CommandHandlerConstructor[R],
	newHandler HandlerConstructor,
	repo R,
	middleware ...Middleware,
) LambdaHandler {
	cmdHandler := newCommandHandler(repo)
	handler := newHandler(cmdHandler)

	if len(middleware) > 0 {
		return middleware[0](handler)
	}

	return handler
}
