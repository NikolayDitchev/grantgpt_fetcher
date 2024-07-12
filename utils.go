package main

import (
	"bytes"
	"regexp"
	"strings"
)

func StripHTMLTags(content string) string {
	re := regexp.MustCompile(`(<[^>]+>)|(&[a-z]{2,5};)`)
	return re.ReplaceAllString(content, " ")
}

type fundingBuffer struct {
	content  *bytes.Buffer
	fileName string
}

func (fu *fundingBuffer) WriteString(data string) error {
	_, err := fu.content.WriteString(data)
	return err
}

func ConvertFieldName(s string) string {
	r := regexp.MustCompile(`(\\b|-|_|\\.)[a-z]`)
	return r.ReplaceAllStringFunc(s, func(t string) string {
		if len(t) == 1 {
			return strings.ToUpper(t)
		} else {
			return strings.ToUpper(string(t[1]))
		}

	})
}
