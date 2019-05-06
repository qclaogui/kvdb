package kvdb_test

import (
	"testing"

	"github.com/qclaogui/kvdb"

	"github.com/google/go-cmp/cmp"
)

func TestStoreX_GetValue(t *testing.T) {
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

	db := kvdb.NewDB()
	for name, test := range tests {
		// Set first
		if test.err == nil {
			db.Set(test.key, test.value)
		}
		t.Run(name, func(t *testing.T) {
			got, err := db.GetV(test.key)
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
	db := kvdb.NewDB()

	got, err := db.GetV("/db/user", "defaultValue")
	if df := cmp.Diff(err, nil); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
	if df := cmp.Diff(got, want); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
}

func TestGetValueWithEmptyDefault(t *testing.T) {
	want := ""
	db := kvdb.NewDB()

	got, err := db.GetV("/db/user", "")
	if df := cmp.Diff(err, nil); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
	if df := cmp.Diff(got, want); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
}

func TestDel(t *testing.T) {
	db := kvdb.NewDB()
	db.Set("/app/port", "8080")
	want := "8080"
	got, err := db.GetV("/app/port")
	if df := cmp.Diff(err, nil); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
	if df := cmp.Diff(got, want); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}

	db.Del("/app/port")
	want = ""
	got, err = db.GetV("/app/port")
	if df := cmp.Diff(err, kvdb.ErrNotExist); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}
}
