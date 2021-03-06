package kvdb

// Storage 接口
type Storage interface {
	Put(string, string)
	PutMany(map[string]string)
	Exists(string) bool
	Get(string, ...string) (string, error)
	GetMany(string) ([]string, error)
	Del(string)
	Flush()
}
