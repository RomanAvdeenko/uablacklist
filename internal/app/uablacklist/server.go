package uablacklist

import (
	"log"
	"net/http"
	"time"

	"uablacklist/internal/app/model"
)

type server struct {
	client *http.Client
}

func newServer() (*server, error) {
	s := new(server)
	s.client = &http.Client{Timeout: 3 * time.Second}
	return s, nil
}

// Start work
func Start(cfg Config) error {
	s, err := newServer()
	if err != nil {
		return err
	}
	log.Println("Server start")
	//
	if err := s.getJson(cfg.url, &model.BlockedRecords); err != nil {
		return err
	}
	//
	if err := s.makeFile(cfg.tplFileName, cfg.outFileName, cfg.dropURL); err != nil {
		return err
	}
	log.Print("Server finish")
	return nil
}
