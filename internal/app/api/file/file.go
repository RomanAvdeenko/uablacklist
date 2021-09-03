package file

import (
	"encoding/json"
	"fmt"
	"os"
	"uablacklist/internal/app/store"
)

type File struct{}

func New() *File { return &File{} }

func (f *File) Getter(uri string, store *store.Store) error {
	file, err := os.Open(uri)
	if err != nil {
		return fmt.Errorf("file open error: %w", err)
	}
	defer file.Close()

	// JSON response
	type blockedRecord struct {
		Urls []string `json:"urls"`
	}
	var blockedRecords blockedRecord

	if err = json.NewDecoder(file).Decode(&blockedRecords); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	for _, val := range blockedRecords.Urls {
		if err := store.AddDomainName(val); err != nil {
			return fmt.Errorf("AddDomainName error: %w", err)
		}
	}
	return nil
}
