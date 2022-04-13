package operations

import (
	"context"
	"fmt"

	"github.com/mcfdn/kvstore/store"
)

type Operation struct {
	Name   string
	Action func(context.Context, store.Storer) (*Result, error)
}

type Result struct {
	Value string
}

type Router struct {
	operations []*Operation
}

func NewRouter() *Router {
	return &Router{}
}

func RegisterOperations(r *Router) {
	r.Register(
		NewSetOperation(),
		NewGetOperation(),
		NewDelOperation(),
	)
}

func (r *Router) Register(ops ...*Operation) {
	for _, op := range ops {
		r.operations = append(r.operations, op)
	}
}

func (r *Router) Route(ctx context.Context, s store.Storer, op string) (*Result, error) {
	for _, operation := range r.operations {
		if operation.Name == op {
			return operation.Action(ctx, s)
		}
	}

	return &Result{}, fmt.Errorf("No handler for operation: %s", op)
}
