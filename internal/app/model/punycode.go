package model

import (
	"fmt"

	"golang.org/x/net/idna"
)

var inda = idna.New()

func (val *DomainName) Punycode() error {
	punyKey, err := inda.ToASCII(string(*val))
	if err != nil {
		return fmt.Errorf("punikey error: %w", err)
	}
	*val = DomainName(punyKey)
	return err
}
