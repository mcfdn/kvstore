package operations

import (
	"context"
	"testing"

	"github.com/mcfdn/kvstore/storetest"
)

func TestExecSetName(t *testing.T) {
	o := NewSetOperation()

	if o.Name != "set" {
		t.Fatalf("Expected operation name to be 'set', received %s", o.Name)
	}
}

func TestExecSetCallsCorrectStoreMethod(t *testing.T) {
	args := make([]string, 3)
	args[0] = "thekey"
	args[1] = "thevalue"
	args[2] = "30"

	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", args)
	s := &storetest.TestStore{}

	NewSetOperation().Action(ctx, s)

	if s.K != "thekey" {
		t.Fatalf("Set was not called with expected key")
	}

	if s.V != "thevalue" {
		t.Fatalf("Set was not called with expected value")
	}
}

func TestExecSetReturnsErrorIfArgsIsEmpty(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", make([]string, 0))
	s := &storetest.TestStore{}

	_, err := NewSetOperation().Action(ctx, s)

	if err == nil {
		t.Fatalf("Expected error was not thrown")
	}

	if err.Error() != "Expected 3 arguments to set" {
		t.Fatalf("Unexpected error thrown")
	}
}

func TestExecSetReturnsErrorIfArgsHasMoreThanThreeEntries(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", make([]string, 4))
	s := &storetest.TestStore{}

	_, err := NewSetOperation().Action(ctx, s)

	if err == nil {
		t.Fatalf("Expected error was not thrown")
	}

	if err.Error() != "Expected 3 arguments to set" {
		t.Fatalf("Unexpected error thrown")
	}
}

func TestExecSetReturnsErrorIfTtlIsLessThanZero(t *testing.T) {
	args := make([]string, 3)
	args[0] = "thekey"
	args[1] = "thevalue"
	args[2] = "-1"

	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", args)
	s := &storetest.TestStore{}

	_, err := NewSetOperation().Action(ctx, s)

	if err == nil {
		t.Fatalf("Expected error was not thrown")
	}

	if err.Error() != "TTL must not be less than 0" {
		t.Fatalf("Unexpected error thrown")
	}
}

func TestExecSetReturnsErrorIfTtlIsMoreThanTenYears(t *testing.T) {
	args := make([]string, 3)
	args[0] = "thekey"
	args[1] = "thevalue"
	args[2] = "315360001"

	ctx := context.Background()
	ctx = context.WithValue(ctx, "args", args)
	s := &storetest.TestStore{}

	_, err := NewSetOperation().Action(ctx, s)

	if err == nil {
		t.Fatalf("Expected error was not thrown")
	}

	if err.Error() != "TTL must not be more than 10 years" {
		t.Fatalf("Unexpected error thrown")
	}
}
