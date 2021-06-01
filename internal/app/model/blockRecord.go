package model

var BlockedRecords map[string]blockedRecord

type blockedRecord struct {
	Alias string   `json:"alias"`
	Term  string   `json:"term"`
	Urls  []string `json:"urls"`
	IPs   []string `json:"ips"`
}
