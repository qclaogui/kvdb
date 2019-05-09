package kvdb

// ErrNotExist key does not exist
const ErrNotExist = kvError("key does not exist")

// ErrNoMatched no keys matched
const ErrNoMatched = kvError("no keys matched")

type kvError string

func (e kvError) Error() string { return string(e) }
