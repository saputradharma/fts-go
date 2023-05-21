package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"
)

type Engine interface {
	LoadDocument(path string) (interface{}, error)
	Search(string) (interface{}, error)
}

const dataPath = "./data/enwiki-latest-abstract1.xml"

type xmlDocument struct {
	Documents []document `xml:"doc"`
}

type document struct {
	Title    string `xml:"title"`
	Url      string `xml:"url"`
	Abstract string `xml:"abstract"`
	ID       int
}

func LoadDocument(path string) ([]document, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := xml.NewDecoder(file)
	xDoc := xmlDocument{}

	if err := decoder.Decode(&xDoc); err != nil {
		return nil, err
	}

	// for i := range xDoc.Documents {
	// 	xDoc.Documents[i].ID = i
	// }

	return xDoc.Documents, nil

}

// naive search with string contains
func search(docs []document, keyword string) []document {
	var result []document
	start := time.Now()
	for _, d := range docs {
		if strings.Contains(d.Abstract, keyword) {
			result = append(result, d)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Searching keyword: %v took %s\n", keyword, elapsed)
	return result
}

func main() {

}
