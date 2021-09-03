package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (val *DomainName) Validate() error {
	err := validation.Validate(*val, validation.Required, is.DNSName)
	if err != nil {
		err = fmt.Errorf("validate error: %v %w", *val, validation.Validate(*val, validation.Required, is.DNSName))
	}
	return err
}
