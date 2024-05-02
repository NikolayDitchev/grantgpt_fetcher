package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Document struct {
	NameDoc            string //`json:"nameDoc"`
	TypeDoc            string //`json:"typeDoc"`
	LanguageDoc        string //`json:"languageDoc"`
	FinalPublicDocDate string //`json:"finalPublicDocDate"`
	DocURL             string //`json:"docUrl"`
}

func (doc *Document) DownloadFile(folderPath string) error {

	filePath := filepath.Join(folderPath, doc.NameDoc+"."+doc.TypeDoc)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(doc.DocURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	//fmt.Println("here")

	return nil
}
