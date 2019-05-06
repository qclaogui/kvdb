package kvdb

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStoreX_Get(t *testing.T) {
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
	db := NewDB()
	for name, test := range tests {
		// Set first
		if test.err == nil {
			db.Set(test.key, test.value)
		}

		t.Run(name, func(t *testing.T) {
			got, err := db.get(test.key)
			if df := cmp.Diff(err, test.err); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
			if df := cmp.Diff(got, test.want); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
		})
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
	"case4": {"[]a]", ErrNoMatch, nil},
	"case5": {"/app/missing/*", ErrNoMatch, nil},
}

func TestStoreX_GetAll(t *testing.T) {
	db := NewDB()
	for key, value := range getAllTestInput {
		db.Set(key, value)
	}

	for name, test := range getAllTests {
		t.Run(name, func(t *testing.T) {
			got, err := db.getAll(test.pattern)
			if df := cmp.Diff(err, test.err); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
			if df := cmp.Diff(got, test.want); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
		})
	}
}
