package main

func createMap() map[string]func([]byte) []byte {
	funcMap := make(map[string]func([]byte) []byte)
	funcMap["Validate"] = validation
	return funcMap
}
