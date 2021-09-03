package main

import (
	"flag"
	"log"
	"uablacklist/internal/app/generator"
)

var (
	jsonURI,
	blockedFileName,
	tplFileName,
	resultFileName,
	dropIP string
)

func init() {
	flag.StringVar(&jsonURI, "json", "https://uablacklist.net/all.json", "JSON URI address")
	flag.StringVar(&blockedFileName, "file", "./configs/blocked_sites.json", "blocked sites file name")
	flag.StringVar(&tplFileName, "tpl", "./configs/db.rpz.zone.tpl", "template file name")
	flag.StringVar(&resultFileName, "out", "./db.rpz.zone", "result file name")
	flag.StringVar(&dropIP, "drop", "127.0.0.1", "rpz drop IP address")
	flag.Parse()
}
func main() {
	cfg := generator.NewConfig(jsonURI, blockedFileName, tplFileName, resultFileName, dropIP)

	if err := generator.Start(*cfg); err != nil {
		log.Fatalf("error: %v\n", err)

	}
}
