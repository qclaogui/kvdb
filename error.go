package kvdb

const ErrNotExist = kvError("key does not exist")
const ErrNoMatch = kvError("no keys match")

type kvError string

func (e kvError) Error() string { return string(e) }
