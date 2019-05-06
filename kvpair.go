package kvdb

// 有序键值对
type kvPair struct {
	Key   string
	Value string
}

type kvPairs []kvPair

func (ks kvPairs) Len() int {
	return len(ks)
}

func (ks kvPairs) Less(i, j int) bool {
	return ks[i].Key < ks[j].Key
}

func (ks kvPairs) Swap(i, j int) {
	ks[i], ks[j] = ks[j], ks[i]
}
