package file

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

// LoadJSON loads a json file into a struct
func LoadJSON[T any](filename string) (T, error) {
	var data T

	fileData, err := os.ReadFile(filename)
	if err != nil {
		return data, errors.Wrap(err, "failed to read file")
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return data, errors.Wrap(err, "failed to unmarshal json")
	}

	return data, nil
}
