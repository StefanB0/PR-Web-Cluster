package database

type DatabaseInstance struct {
	memory map[string][]byte
}

type Pair struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}

func NewDatabase() DatabaseInstance {
	memory := make(map[string][]byte)

	return DatabaseInstance{memory: memory}
}
