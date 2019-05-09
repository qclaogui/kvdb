package kvdb_test

import (
	"testing"

	"github.com/qclaogui/kvdb"

	"github.com/google/go-cmp/cmp"
)

func TestMem_GetValue(t *testing.T) {
	var tests = map[string]struct {
		key   string
		value string
		err   error
		want  string
	}{
		"case1": {"/db/user", "admin", nil, "admin"},
		"case2": {"/db/pass", "foo", nil, "foo"},
		"case3": {"/missing", "", kvdb.ErrNotExist, ""},
	}

	db := kvdb.NewMem()
	for name, test := range tests {
		// Put first
		if test.err == nil {
			db.Put(test.key, test.value)
		}
		t.Run(name, func(t *testing.T) {
			got, err := db.Get(test.key)
			if df := cmp.Diff(err, test.err); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
			if df := cmp.Diff(got, test.want); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
		})
	}
}

func TestGetValueWithDefault(t *testing.T) {
	want := "defaultValue"
	db := kvdb.NewMem()

	got, err := db.Get("/db/user", "defaultValue")
	if df := cmp.Diff(err, nil); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
	if df := cmp.Diff(got, want); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
}

func TestGetValueWithEmptyDefault(t *testing.T) {
	want := ""
	db := kvdb.NewMem()

	got, err := db.Get("/db/user", "")
	if df := cmp.Diff(err, nil); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
	if df := cmp.Diff(got, want); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
}

func TestDel(t *testing.T) {
	db := kvdb.NewMem()
	db.Put("/app/port", "8080")
	want := "8080"
	got, err := db.Get("/app/port")
	if df := cmp.Diff(err, nil); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
	if df := cmp.Diff(got, want); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}

	db.Del("/app/port")
	want = ""
	got, err = db.Get("/app/port")
	if df := cmp.Diff(err, kvdb.ErrNotExist); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
}
