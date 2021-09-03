package model

type DomainName string

func (val *DomainName) Validate() (DomainName, bool) {
	return *val, true
}
