package operations

import (
	"context"
	"testing"

	"github.com/mcfdn/kvstore/storetest"
)

func TestExecGetName(t *testing.T) {
	o := NewGetOperation()

	if o.Name != "get" {
		t.Fatalf("Expected operation name to be 'get', received %s", o.Name)
	}
}

func TestExecGetCallsCorrectStoreMethod(t *testing.T) {
	args := make([]string, 1)
	args[0] = "key"

	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", args)
	s := &storetest.TestStore{}

	NewGetOperation().Action(ctx, s)

	if s.K != "key" {
		t.Fatalf("Get was not called with expected key")
	}
}

func TestExecGetReturnsValue(t *testing.T) {
	args := make([]string, 1)
	args[0] = "key"

	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", args)
	s := &storetest.TestStore{
		V: "thevalue",
	}

	result, _ := NewGetOperation().Action(ctx, s)

	if result.Value != "thevalue" {
		t.Fatalf("Get did not return the expected value")
	}
}

func TestExecGetReturnsErrorIfArgsIsEmpty(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", make([]string, 0))
	s := &storetest.TestStore{}

	_, err := NewGetOperation().Action(ctx, s)

	if err == nil {
		t.Fatalf("Expected error was not thrown")
	}

	if err.Error() != "Expected 1 argument to get" {
		t.Fatalf("Unexpected error thrown")
	}
}

func TestExecGetReturnsErrorIfArgsHasMoreThanOneEntry(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", make([]string, 2))
	s := &storetest.TestStore{}

	_, err := NewGetOperation().Action(ctx, s)

	if err == nil {
		t.Fatalf("Expected error was not thrown")
	}

	if err.Error() != "Expected 1 argument to get" {
		t.Fatalf("Unexpected error thrown")
	}
}
