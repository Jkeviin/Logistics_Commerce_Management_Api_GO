package loader

import (
	"encoding/json"
	"log"
	"os"
)

// JSONLoader is a struct that represents a generic JSON loader
type JSONLoader[T any, KeyType comparable] struct {
	Path string // Path to the JSON file
}

// NewJSONLoader creates a new instance of JSONLoader
func NewJSONLoader[T any, KeyType comparable](path string) *JSONLoader[T, KeyType] {
	return &JSONLoader[T, KeyType]{Path: path}
}

// LoadToMap loads data from the JSON file into a map using a key extractor function
func (l *JSONLoader[T, KeyType]) LoadToMap(keyExtractor func(T) KeyType) (map[KeyType]T, error) {
	// Open the file
	file, err := os.Open(l.Path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			log.Printf("Error closing file: %v", closeErr)
		}
	}(file)
	// Decode the JSON file into a slice of T
	var data []T
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}
	// Transform the slice into a map
	result := make(map[KeyType]T)
	for _, item := range data {
		key := keyExtractor(item)
		result[key] = item
	}
	return result, nil
}

// SaveToJSON saves the map data to the JSON file
func (l *JSONLoader[T, KeyType]) SaveToJSON(data map[KeyType]T) error {
	// Convert map to slice
	var items []T
	for _, item := range data {
		items = append(items, item)
	}

	// Create or truncate the file
	file, err := os.Create(l.Path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			log.Printf("Error closing file: %v", closeErr)
		}
	}(file)

	// Encode the data to JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(items)
}
