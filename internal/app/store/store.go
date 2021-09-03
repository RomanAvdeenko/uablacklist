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

func (s *Store) AddDomainName(val string) error {
	domainName := model.DomainName(val)
	if err := domainName.Punycode(); err != nil {
		return err
	}
	if err := domainName.Validate(); err != nil {
		return err
	}
	s.domainNames[domainName] = struct{}{}
	return nil
}

func (s *Store) GetDomainNames() map[model.DomainName]struct{} {
	return s.domainNames
}
