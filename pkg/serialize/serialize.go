package serialize

func Do(obj interface{}) (interface{}, error) {
	// Implemented in serialize.lua
	return nil, nil
}

func Undo(from, to interface{}) error {
	// Implemented in serialize.lua
	return nil
}
