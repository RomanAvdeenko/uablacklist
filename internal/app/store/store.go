package store

import (
	"uablacklist/internal/app/model"
)

type Store struct {
	domainNames map[model.DomainName]struct{}
}

func New() *Store {
	return &Store{domainNames: make(map[model.DomainName]struct{})}
}

func (s *Store) AddDomainName(val string) {
	s.domainNames[model.DomainName(val)] = struct{}{}
}

func (s *Store) GetDomainNames() map[model.DomainName]struct{} {
	return s.domainNames
}
