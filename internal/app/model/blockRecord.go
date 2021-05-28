package model

var BlockRecords map[string]blockRecord

type blockRecord struct {
	Alias string   `json:"alias"`
	Term  string   `json:"term"`
	Urls  []string `json:"urls"`
	IPs   []string `json:"ips"`
}
