package generator

import (
	"fmt"
	"log"
	"uablacklist/internal/app/api/file"
	"uablacklist/internal/app/api/uablacklist"
	"uablacklist/internal/app/store"
)

// Start work
func Start(cfg Config) error {
	log.Println("Start")
	store := store.New()
	//
	uablacklist := uablacklist.New()
	if err := uablacklist.Get(cfg.uri, store); err != nil {
		return fmt.Errorf("json.Get() errr: %w", err)
	}
	//
	file := file.New()
	if err := file.Getter(cfg.blockedFileName, store); err != nil {
		return fmt.Errorf("file.Get() err: %w", err)
	}
	//
	if err := store.MakeFile(cfg.tplFileName, cfg.outFileName, cfg.dropIP); err != nil {
		return err
	}
	log.Print("Finish")
	return nil
}
