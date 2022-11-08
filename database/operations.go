package database

func (d *DatabaseInstance) Create(pair Pair) {
	d.memory[pair.Key] = pair.Value
}

func (d *DatabaseInstance) Read(key string) []byte {
	return d.memory[key]
}

func (d *DatabaseInstance) Update(pair Pair) {
	d.memory[pair.Key] = pair.Value
}

func (d *DatabaseInstance) Delete(key string) {
	delete(d.memory, key)
}
