package main

import (
	"flag"
	"log"

	"github.com/RomanAvdeenko/uablacklist/internal/app/uablacklist"
)

var (
	jsonURL,
	resultFileName string
)

func init() {
	flag.StringVar(&jsonURL, "url", "https://uablacklist.net/all.json", "JSON URL address")
	flag.StringVar(&resultFileName, "out", "out.txt", "result file name")
	flag.Parse()
}
func main() {
	cfg := uablacklist.NewConfig(jsonURL, resultFileName)

	if err := uablacklist.Start(*cfg); err != nil {
		log.Fatalf("error: %v\n", err)

	}
}
