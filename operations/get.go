package operations

import (
	"context"
	"errors"

	"github.com/mcfdn/kvstore/store"
)

func NewGetOperation() *Operation {
	return &Operation{
		Name:   "get",
		Action: ExecGet(),
	}
}

func ExecGet() func(ctx context.Context, s store.Storer) (*Result, error) {
	return func(ctx context.Context, s store.Storer) (*Result, error) {
		args := ctx.Value("args").([]string)

		if len(args) != 1 {
			return &Result{}, errors.New("Expected 1 argument to get")
		}

		val := s.Get(args[0])

		return &Result{Value: val}, nil
	}
}
