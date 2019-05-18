package kvdb

import (
	"reflect"
	"testing"
)

func TestMem_Get(t *testing.T) {
	var tests = map[string]struct {
		key   string
		value string
		err   error
		want  kvPair
	}{
		"case1": {"/db/user", "admin", nil, kvPair{"/db/user", "admin"}},
		"case2": {"/db/pass", "foo", nil, kvPair{"/db/pass", "foo"}},
		"case3": {"/missing", "", ErrNotExist, kvPair{}},
	}
	db := NewMem()
	for name, test := range tests {
		// Put first
		if test.err == nil {
			db.Put(test.key, test.value)
		}

		t.Run(name, func(t *testing.T) {
			got, err := db.get(test.key)
			assertEqual(t, err, test.err)
			assertEqual(t, got, test.want)
		})
	}
}

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
			"✘got: %v\n\x1b[92m"+
			"want: %v\x1b[39m", got, want)
	}
}

var getAllTestInput = map[string]string{
	"/app/db/pass":               "foo",
	"/app/db/user":               "admin",
	"/app/port":                  "443",
	"/app/url":                   "app.example.com",
	"/app/vhosts/host1":          "app.example.com",
	"/app/upstream/host1":        "203.0.113.0.1:8080",
	"/app/upstream/host1/domain": "app.example.com",
	"/app/upstream/host2":        "203.0.113.0.2:8080",
	"/app/upstream/host2/domain": "app.example.com",
}

var getAllTests = map[string]struct {
	pattern string
	err     error
	want    kvPairs
}{
	"case1": {"/app/db/*", nil,
		kvPairs{
			kvPair{"/app/db/pass", "foo"},
			kvPair{"/app/db/user", "admin"}}},
	"case2": {"/app/*/host1", nil,
		kvPairs{
			kvPair{"/app/upstream/host1", "203.0.113.0.1:8080"},
			kvPair{"/app/vhosts/host1", "app.example.com"}}},

	"case3": {"/app/upstream/*", nil,
		kvPairs{
			kvPair{"/app/upstream/host1", "203.0.113.0.1:8080"},
			kvPair{"/app/upstream/host2", "203.0.113.0.2:8080"}}},
	"case4": {"[]a]", ErrNoMatched, nil},
	"case5": {"/app/missing/*", ErrNoMatched, nil},
}

func TestMem_GetAllMatched(t *testing.T) {
	db := NewMem()
	for key, value := range getAllTestInput {
		db.Put(key, value)
	}

	for name, test := range getAllTests {
		t.Run(name, func(t *testing.T) {
			got, err := db.getAllMatched(test.pattern)
			assertEqual(t, err, test.err)
			assertEqual(t, got, test.want)
		})
	}
}
