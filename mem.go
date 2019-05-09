package kvdb

import (
	"path"
	"sort"
	"sync"
)

type mem struct{ m *sync.Map }

// NewMem new a map to store
func NewMem() *mem { return &mem{new(sync.Map)} }

func (s *mem) Put(key, value string) { s.m.Store(key, kvPair{key, value}) }

func (s *mem) PutMany(m map[string]string) {
	for k, v := range m {
		s.Put(k, v)
	}
}

func (s *mem) Exists(key string) bool {
	if _, err := s.get(key); err != nil {
		return false
	}
	return true
}

func (s *mem) get(key string) (kvPair, error) {
	v, ok := s.m.Load(key)
	if !ok {
		return kvPair{}, ErrNotExist
	}
	return v.(kvPair), nil
}

// Get 获取key的相应value
func (s *mem) Get(key string, defaultValue ...string) (string, error) {
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

func (s *mem) getAllMatched(pattern string) (kvPairs, error) {
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
		return nil, ErrNoMatched
	}
	sort.Sort(kvs)
	return kvs, nil
}

// GetMany 获取匹配到pattern的所有keys的value
func (s *mem) GetMany(pattern string) ([]string, error) {
	vs := make([]string, 0)
	kvs, err := s.getAllMatched(pattern)
	if err != nil {
		return nil, err
	}
	for _, kv := range kvs {
		vs = append(vs, kv.Value)
	}
	sort.Strings(vs)
	return vs, nil
}

func (s *mem) Del(key string) { s.m.Delete(key) }

func (s *mem) Flush() {
	s.m.Range(func(key, _ interface{}) bool {
		s.m.Delete(key)
		return true
	})
}
