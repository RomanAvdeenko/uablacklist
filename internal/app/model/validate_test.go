package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate(t *testing.T) {
	tests := []struct {
		name    string
		provide DomainName
		isValid bool
	}{
		{
			"Правильное доменное имя",
			"www.domain.com",
			true,
		},
		{
			"Доменное имя на русском",
			"82.мвд.рф",
			false,
		},
		{
			"Пустое имя",
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isValid {
				assert.Nil(t, tt.provide.Validate())
			} else {
				assert.NotNil(t, tt.provide.Validate())
			}
		})
	}
}
