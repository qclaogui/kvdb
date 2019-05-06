package kvdb

import (
	"path"
	"sort"
	"sync"
)

type store struct{ m *sync.Map }

func NewDB() *store { return &store{new(sync.Map)} }

func (s *store) Set(key, value string) { s.m.Store(key, kvPair{key, value}) }

func (s *store) Exists(key string) bool {
	if _, err := s.get(key); err != nil {
		return false
	}
	return true
}

func (s *store) get(key string) (kvPair, error) {
	if v, ok := s.m.Load(key); !ok {
		return kvPair{}, ErrNotExist
	} else {
		return v.(kvPair), nil
	}
}

func (s *store) GetV(key string, defaultValue ...string) (string, error) {
	kv, err := s.get(key)
	if err != nil {
		// 如果有设置默认值,将返回defaultValue中的第一个作为默认值
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return "", err
	}
	return kv.Value, nil
}

func (s *store) getAll(pattern string) (kvPairs, error) {
	kvs := make(kvPairs, 0)
	s.m.Range(func(_, value interface{}) bool {
		kv := value.(kvPair)
		if matched, _ := path.Match(pattern, kv.Key); matched {
			kvs = append(kvs, kv)
		}
		return true
	})
	// 查看是否匹配到
	if len(kvs) == 0 {
		return nil, ErrNoMatch
	}
	sort.Sort(kvs)
	return kvs, nil
}

func (s *store) GetVs(pattern string) ([]string, error) {
	vs := make([]string, 0)
	kvs, err := s.getAll(pattern)
	if err != nil {
		return nil, err
	}
	for _, kv := range kvs {
		vs = append(vs, kv.Value)
	}
	sort.Strings(vs)
	return vs, nil
}

func (s *store) Del(key string) { s.m.Delete(key) }

func (s *store) purge() {
	s.m.Range(func(key, _ interface{}) bool {
		s.m.Delete(key)
		return true
	})
}
