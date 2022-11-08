package database

func (d *DatabaseInstance) OverwriteMemory(DBMemory map[string][]byte) (responseCode int) {
	for key, element := range DBMemory {
		d.memory[key] = element
	}

	return OK
}