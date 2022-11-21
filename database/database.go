package database

type DatabaseInstance struct {
	Memory map[string][]byte `json:"memory"`
}

type Pair struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}

func NewDatabase() DatabaseInstance {
	memory := make(map[string][]byte)

	return DatabaseInstance{Memory: memory}
}

func (d *DatabaseInstance) Create(key string, value []byte) {
	d.Memory[key] = value
}

func (d *DatabaseInstance) Read(key string) []byte {
	return d.Memory[key]
}

func (d *DatabaseInstance) Update(key string, value []byte) {
	d.Memory[key] = value
}

func (d *DatabaseInstance) Delete(key string) {
	delete(d.Memory, key)
}

func (d *DatabaseInstance) OverwriteMemory(DBMemory map[string][]byte) {
	for key, element := range DBMemory {
		d.Memory[key] = element
	}
}

func (d *DatabaseInstance) GetKeyValuePairs() ([]string, [][]byte) {
	keys := make([]string, len(d.Memory))
	values := make([][]byte, len(d.Memory))

	i := 0
	for key, value := range d.Memory {
		keys[i] = key
		values[i] = value
	}

	return keys, values
}
