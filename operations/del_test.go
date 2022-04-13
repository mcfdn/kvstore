package operations

import (
	"context"
	"testing"

	"github.com/mcfdn/kvstore/storetest"
)

func TestExecDelName(t *testing.T) {
	o := NewDelOperation()

	if o.Name != "del" {
		t.Fatalf("Expected operation name to be 'del', received %s", o.Name)
	}
}

func TestExecDelCallsCorrectStoreMethod(t *testing.T) {
	args := make([]string, 1)
	args[0] = "key"

	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", args)
	s := &storetest.TestStore{}

	NewDelOperation().Action(ctx, s)

	if s.K != "key" {
		t.Fatalf("Delete was not called with expected key")
	}
}

func TestExecDelReturnsErrorIfArgsIsEmpty(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", make([]string, 0))
	s := &storetest.TestStore{}

	_, err := NewDelOperation().Action(ctx, s)

	if err == nil {
		t.Fatalf("Expected error was not thrown")
	}

	if err.Error() != "Expected 1 argument to del" {
		t.Fatalf("Unexpected error thrown")
	}
}

func TestExecDelReturnsErrorIfArgsHasMoreThanOneEntry(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", make([]string, 2))
	s := &storetest.TestStore{}

	_, err := NewDelOperation().Action(ctx, s)

	if err == nil {
		t.Fatalf("Expected error was not thrown")
	}

	if err.Error() != "Expected 1 argument to del" {
		t.Fatalf("Unexpected error thrown")
	}
}
