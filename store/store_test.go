package store

import (
	"testing"
	"time"
)

func TestSetGet(t *testing.T) {
	s := New()
	s.Set("key", "value", 0)

	val := s.Get("key")

	if val != "value" {
		t.Fatalf("Expected '%s', received '%s'", "value", val)
	}
}

func TestGetMissingKey(t *testing.T) {
	s := New()

	val := s.Get("key")

	if val != "" {
		t.Fatalf("Expected empty value, received '%s'", val)
	}
}

func TestSetGetWithValidTtl(t *testing.T) {
	timeNow = func() time.Time {
		time, _ := time.Parse(time.RFC3339, "2022-01-14T01:04:00Z")
		return time
	}

	s := New()
	s.Set("key", "value", 10)

	timeNow = func() time.Time {
		time, _ := time.Parse(time.RFC3339, "2022-01-14T01:04:09Z")
		return time
	}

	val := s.Get("key")

	if val != "value" {
		t.Fatalf("Expected '%s', received '%s'", "value", val)
	}
}

func TestSetGetWithExpiredTtl(t *testing.T) {
	timeNow = func() time.Time {
		time, _ := time.Parse(time.RFC3339, "2022-01-14T01:04:00Z")
		return time
	}

	s := New()
	s.Set("key", "value", 10)

	timeNow = func() time.Time {
		time, _ := time.Parse(time.RFC3339, "2022-01-14T01:04:10Z")
		return time
	}

	val := s.Get("key")

	if val != "" {
		t.Fatalf("Received unexpected value")
	}
}

func TestDel(t *testing.T) {
	s := New()
	s.Set("key", "value", 0)

	s.Delete("key")

	if val := s.Get("key"); val != "" {
		t.Fatalf("Expected empty value, received '%s'", val)
	}
}
