package main

import (
	"flag"
	"log"
	"uablacklist/internal/app/uablacklist"
)

var (
	jsonURL,
	tplFileName,
	resultFileName,
	dropURL string
)

func init() {
	flag.StringVar(&jsonURL, "url", "https://uablacklist.net/all.json", "JSON URL address")
	flag.StringVar(&tplFileName, "tpl", "./configs/db.rpz.zone.tpl", "template file name")
	flag.StringVar(&resultFileName, "out", "db.rpz.zone", "result file name")
	flag.StringVar(&dropURL, "drop", "127.0.0.1", "rpz drop URL")
	flag.Parse()
}
func main() {
	cfg := uablacklist.NewConfig(jsonURL, tplFileName, resultFileName, dropURL)

	if err := uablacklist.Start(*cfg); err != nil {
		log.Fatalf("error: %v\n", err)

	}
}
