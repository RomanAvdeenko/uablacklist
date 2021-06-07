package uablacklist

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/RomanAvdeenko/uablacklist/internal/app/model"
	"golang.org/x/net/idna"
)

const generationMessage = "; This code is generated!!!\n"

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

// Fetch JSON
func (s *server) getJson(url string, target interface{}) error {
	resp, err := s.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// Make config file
func (s *server) makeFile(tplFileName, outFileName, dropURL string) error {
	// Template manipulation
	tpl, err := template.ParseFiles(tplFileName)
	if err != nil {
		return err
	}

	file, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	file.WriteString(generationMessage)

	var tplData struct {
		SERIAL   string
		DROP_URL string
		RECORDS  []string
	}
	tplData.SERIAL = time.Now().Format("2006010215")
	tplData.DROP_URL = dropURL
	tplData.RECORDS = make([]string, 0, len(model.BlockedRecords))

	// punycode
	p := idna.New()
	for key := range model.BlockedRecords {
		punyKey, err := p.ToASCII(key)
		if err != nil {
			continue
		}
		tplData.RECORDS = append(tplData.RECORDS, punyKey)
	}

	err = tpl.Execute(file, tplData)
	if err != nil {
		return err
	}

	file.WriteString(generationMessage)

	return nil
}
