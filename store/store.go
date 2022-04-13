package store

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	timeNow = time.Now
)

type StoreSetter interface {
	Set(k string, v string, ttl uint)
}

type StoreGetter interface {
	Get(k string) string
}

type StoreDeleter interface {
	Delete(k string)
}

type Storer interface {
	StoreSetter
	StoreGetter
	StoreDeleter
}

type Store struct {
	entries map[string]string
}

func New() *Store {
	return &Store{
		entries: make(map[string]string),
	}
}

func (s *Store) Set(k string, v string, ttl uint) {
	var exp int64 = 0

	if ttl > 0 {
		t := timeNow().Add(time.Second * time.Duration(ttl))
		exp = t.Unix()
	}

	s.entries[k] = fmt.Sprintf("%d:%s", exp, v)
}

func (s *Store) Get(k string) string {
	parts := strings.Split(s.entries[k], ":")
	exp, _ := strconv.ParseInt(parts[0], 10, 64)

	if exp > 0 && exp <= timeNow().Unix() {
		delete(s.entries, k)

		return ""
	}

	return strings.Join(parts[1:], ":")
}

func (s *Store) Delete(k string) {
	delete(s.entries, k)
}
