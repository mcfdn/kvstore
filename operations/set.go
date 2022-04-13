package operations

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/mcfdn/kvstore/store"
)

func NewSetOperation() *Operation {
	return &Operation{
		Name:   "set",
		Action: ExecSet(),
	}
}

func ExecSet() func(ctx context.Context, s store.Storer) (*Result, error) {
	return func(ctx context.Context, s store.Storer) (*Result, error) {
		args := ctx.Value("args").([]string)

		if len(args) != 3 {
			return &Result{}, errors.New("Expected 3 arguments to set")
		}

		ttl, _ := strconv.Atoi(args[2])
		if ttl < 0 {
			return &Result{}, errors.New("TTL must not be less than 0")
		}

		if time.Second*time.Duration(ttl) > time.Duration(time.Hour*87600) {
			return &Result{}, errors.New("TTL must not be more than 10 years")
		}

		s.Set(args[0], args[1], uint(ttl))

		return &Result{}, nil
	}
}
