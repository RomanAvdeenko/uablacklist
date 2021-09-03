package api

import "uablacklist/internal/app/store"

type Getter interface {
	Get(string, *store.Store) error
}
