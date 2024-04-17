package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

type OvalDefinitions struct {
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

func main() {
	resp, _ := http.Get("https://redos.red-soft.ru/support/secure/redos.xml")
	bytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	var s OvalDefinitions
	xml.Unmarshal([]byte(bytes), &s)
}
