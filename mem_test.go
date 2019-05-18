package kvdb_test

import (
	"reflect"
	"testing"

	"github.com/qclaogui/kvdb"
)

func TestMem_Get(t *testing.T) {
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

			assertEqual(t, err, test.err)
			assertEqual(t, got, test.want)
		})
	}
}

func TestGetWithDefault(t *testing.T) {
	want := "defaultValue"
	db := kvdb.NewMem()

	got, err := db.Get("/db/user", "defaultValue")

	assertEqual(t, err, nil)
	assertEqual(t, got, want)
}

func TestGetValueWithEmptyDefault(t *testing.T) {
	want := ""
	db := kvdb.NewMem()

	got, err := db.Get("/db/user", "")

	assertEqual(t, err, nil)
	assertEqual(t, got, want)
}

func TestDel(t *testing.T) {
	db := kvdb.NewMem()
	db.Put("/app/port", "8080")
	want := "8080"
	got, err := db.Get("/app/port")

	assertEqual(t, err, nil)
	assertEqual(t, got, want)

	db.Del("/app/port")
	want = ""
	got, err = db.Get("/app/port")

	assertEqual(t, err, kvdb.ErrNotExist)
	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
			"✘got: %v\n\x1b[92m"+
			"want: %v\x1b[39m", got, want)
	}
}
