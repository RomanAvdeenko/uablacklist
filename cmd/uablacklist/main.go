package main

import (
	"flag"
	"log"

	"github.com/RomanAvdeenko/uablacklist/internal/app/uablacklist"
)

var (
	jsonURL,
	tplFileName,
	resultFileName string
)

func init() {
	flag.StringVar(&jsonURL, "url", "https://uablacklist.net/all.json", "JSON URL address")
	flag.StringVar(&tplFileName, "tpl", "db.rpz.zone.tpl", "template file name")
	flag.StringVar(&resultFileName, "out", "db.rpz.zone", "result file name")
	flag.Parse()
}
func main() {
	cfg := uablacklist.NewConfig(jsonURL, tplFileName, resultFileName)

	if err := uablacklist.Start(*cfg); err != nil {
		log.Fatalf("error: %v\n", err)

	}
	// test CI
}
