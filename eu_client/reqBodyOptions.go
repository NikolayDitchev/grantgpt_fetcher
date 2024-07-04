package eu_client

import (
	"encoding/json"
	"mime/multipart"
	"net/textproto"
)

type RequestBodyOption func(*multipart.Writer) error

func WithQuery(q *Query) RequestBodyOption {
	return func(writer *multipart.Writer) error {
		return AddWriterPart("query", q, writer)
	}
}

// ISO 639-1 language nomenclature
func WithLanguages(languages ...string) RequestBodyOption {
	return func(writer *multipart.Writer) error {
		return AddWriterPart("languages", languages, writer)
	}
}

func AddWriterPart(name string, content any, writer *multipart.Writer) error {

	json, err := json.Marshal(content)
	if err != nil {
		return err
	}

	var header textproto.MIMEHeader = make(textproto.MIMEHeader)
	header.Add("Content-Disposition", `form-data; name="`+name+`";`)
	header.Add("Content-Type", "application/json")

	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}

	part.Write(json)

	return nil
}
