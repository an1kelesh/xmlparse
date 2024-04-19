package xmlparse

import (
	"encoding/xml"
	"io"
	"net/http"
)

type redosovalxml struct {
	Definitions struct {
		Definition []struct {
			Metadata struct {
				Title     string `xml:"title"`
				Reference []struct {
					RefID  string `xml:"ref_id,attr"`
					RefURL string `xml:"ref_url,attr"`
				} `xml:"reference"`
				Description string `xml:"description"`
				Advisory    struct {
					Severity string `xml:"severity"`
				} `xml:"advisory"`
			} `xml:"metadata"`
			Criteria struct {
				Criterion struct {
					Comment string `xml:"comment,attr"`
				} `xml:"criterion"`
			} `xml:"criteria"`
		} `xml:"definition"`
	} `xml:"definitions"`
}

var s redosovalxml

func Xmldownload(urlredxml string) redosovalxml {
	resp, _ := http.Get(urlredxml)
	bytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal([]byte(bytes), &s)
	return s
}
