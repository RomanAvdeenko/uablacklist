package uablacklist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"uablacklist/internal/app/store"
)

type Uablacklist struct{}

func New() *Uablacklist { return &Uablacklist{} }

func (u *Uablacklist) Get(uri string, store *store.Store) error {
	// JSON response
	type blockedRecord struct {
		Alias string   `json:"alias"`
		Term  string   `json:"term"`
		Urls  []string `json:"urls"`
		IPs   []string `json:"ips"`
	}

	var blockedRecords map[string]blockedRecord

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(uri)
	if err != nil {
		return fmt.Errorf("get api error: %w", err)
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&blockedRecords); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	for key := range blockedRecords {
		if err := store.AddDomainName(key); err != nil {
			return fmt.Errorf("AddDomainName error: %w", err)
		}
	}

	return nil
}
