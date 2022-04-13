package operations

import (
	"context"
	"errors"

	"github.com/mcfdn/kvstore/store"
)

func NewDelOperation() *Operation {
	return &Operation{
		Name:   "del",
		Action: ExecDel(),
	}
}

func ExecDel() func(ctx context.Context, s store.Storer) (*Result, error) {
	return func(ctx context.Context, s store.Storer) (*Result, error) {
		args := ctx.Value("args").([]string)

		if len(args) != 1 {
			return &Result{}, errors.New("Expected 1 argument to del")
		}

		s.Delete(args[0])

		return &Result{}, nil
	}
}
