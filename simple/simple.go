package simple

import (
	"encoding/xml"
	"os"
	"strings"
)

type simpleFts struct {
	documents []document
}

type xmlDocument struct {
	Documents []document `xml:"doc"`
}

type document struct {
	Title    string `xml:"title"`
	Url      string `xml:"url"`
	Abstract string `xml:"abstract"`
	ID       int
}

func NewSimpleFts(path string) (*simpleFts, error) {
	engine := simpleFts{}
	if err := engine.LoadDocument(path); err != nil {
		return nil, err
	}

	return &engine, nil
}

func (s *simpleFts) LoadDocument(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := xml.NewDecoder(file)
	xDoc := xmlDocument{}

	if err := decoder.Decode(&xDoc); err != nil {
		return err
	}

	s.documents = xDoc.Documents

	return nil
}

// naive search with string contains
func (s *simpleFts) Search(keyword string) []document {
	var result []document
	for _, d := range s.documents {
		if strings.Contains(d.Abstract, keyword) {
			result = append(result, d)
		}
	}

	return result
}
