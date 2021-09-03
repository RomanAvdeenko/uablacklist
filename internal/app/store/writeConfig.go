package store

import (
	"html/template"
	"os"
	"time"
)

//
func writeGenMessage(file *os.File) error {
	const generationMessage = "; This code is generated!!!\n"
	_, err := file.WriteString(generationMessage)
	if err != nil {
		return err
	}
	return nil
}

// Make output config file
func (s *Store) MakeFile(tplFileName, outFileName, dropURL string) error {
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
	//
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

	tplData.RECORDS = make([]string, 0, len(s.GetDomainNames()))
	for key := range s.GetDomainNames() {
		tplData.RECORDS = append(tplData.RECORDS, string(key))
	}

	err = tpl.Execute(file, tplData)
	if err != nil {
		return err
	}
	//
	if err := writeGenMessage(file); err != nil {
		return err
	}

	return nil
}
