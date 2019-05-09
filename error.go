package kvdb

const ErrNotExist = kvError("key does not exist")
const ErrNoMatched = kvError("no keys matched")

type kvError string

func (e kvError) Error() string { return string(e) }
