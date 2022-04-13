package operations

import (
	"context"
	"testing"

	"github.com/mcfdn/kvstore/store"
)

func TestRegisteredOperationsAreRouted(t *testing.T) {
	r := NewRouter()

	r.Register(&Operation{
		Name: "test",
		Action: func(ctx context.Context, s store.Storer) (*Result, error) {
			return &Result{
				Value: "success",
			}, nil
		},
	})

	result, _ := r.Route(context.Background(), &store.Store{}, "test")

	if result.Value != "success" {
		t.Fatalf("Result was not expected value")
	}
}

func TestRouteReturnsErrorForUnknownOperation(t *testing.T) {
	r := NewRouter()

	// Register a control Operation to be sure it doesn't get returned accidentally
	r.Register(&Operation{
		Name: "test",
		Action: func(ctx context.Context, s store.Storer) (*Result, error) {
			return &Result{
				Value: "success",
			}, nil
		},
	})

	_, err := r.Route(context.Background(), &store.Store{}, "op")

	if err == nil {
		t.Fatalf("An error was expected but nil was returned")
	}

	if err.Error() != "No handler for operation: op" {
		t.Fatalf("Unexpected error was returned")
	}
}
