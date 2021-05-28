package main

import (
	"flag"

	"github.com/RomanAvdeenko/uablacklist/internal/app/uablacklist"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "url", "https://uablacklist.net/all.json", "JSON URL address")
	flag.Parse()
}
func main() {
	cfg := uablacklist.NewConfig(configPath)
	uablacklist.Start(*cfg)
}
