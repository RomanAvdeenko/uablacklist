package uablacklist

import (
	"encoding/json"
	"os"
	"text/template"
	"time"

	"uablacklist/internal/app/model"

	"golang.org/x/net/idna"
)

const generationMessage = "; This code is generated!!!\n"

func writeGenMessage(file *os.File) error {
	_, err := file.WriteString(generationMessage)
	if err != nil {
		return err
	}
	return nil
}

// Make output config file
func (s *server) makeFile(tplFileName, outFileName, dropURL string) error {
	// Template manipulation
	tpl, err := template.ParseFiles(tplFileName)
	if err != nil {
		return err
	}

	// If the file exists then rename it.
	if _, err = os.Stat(outFileName); err == nil {
		if err = os.Rename(outFileName, outFileName+".orig"); err != nil {
			return err
		}

	}

	//
	file, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	if err := writeGenMessage(file); err != nil {
		return err
	}

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

	if err := writeGenMessage(file); err != nil {
		return err
	}

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
