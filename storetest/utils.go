package storetest

type TestStore struct {
	K   string
	V   string
	Ttl uint
}

func (s *TestStore) Set(k string, v string, ttl uint) {
	s.K = k
	s.V = v
	s.Ttl = ttl
}

func (s *TestStore) Get(k string) string {
	s.K = k

	return s.V
}

func (s *TestStore) Delete(k string) {
	s.K = k
}
