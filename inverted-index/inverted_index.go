package inverted_index

import (
	"encoding/xml"
	"os"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/kljensen/snowball"
)

type invertedIndex struct {
	index     map[string][]int
	documents map[int]document
}

func New(path string) (*invertedIndex, error) {
	engine := invertedIndex{
		index:     make(map[string][]int),
		documents: make(map[int]document),
	}

	err := engine.LoadDocument(path)
	if err != nil {
		return nil, err
	}

	return &engine, nil
}

// LoadDocument is to load the xml document data
// the data also will be added to the index list
func (e *invertedIndex) LoadDocument(path string) error {
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

	for k, d := range xDoc.Documents {
		d.ID = k
		e.add(d)
	}

	return nil
}

// Search searching documents based on a keyword
func (e *invertedIndex) Search(keyword string) []document {

	tokens, _ := e.tokenize(keyword)

	queryResult := make(map[int]document)
	for _, t := range tokens {
		if ids, ok := e.index[t]; ok {
			for _, id := range ids {
				// only add non-exists documentId to queryResult
				if _, ok := queryResult[id]; !ok {
					queryResult[id] = e.documents[id]
				}
			}
		}
	}
	result := []document{}
	for _, d := range queryResult {
		result = append(result, d)
	}
	return result
}

func (e *invertedIndex) add(d document) {
	e.documents[d.ID] = d
	e.addIndex(d.Abstract, d.ID)
}

func (e *invertedIndex) addIndex(text string, documentId int) error {

	// parse text to tokens
	tokens, err := e.tokenize(text)
	if err != nil {
		return err
	}
	// add new tokens to index
	for _, st := range tokens {
		// exists
		if ids, ok := e.index[st]; ok {
			// only add non-exists documentId in the index
			if ids[len(ids)-1] != documentId {
				e.index[st] = append(ids, documentId)
			}
		} else {
			e.index[st] = []int{documentId}
		}
	}

	return nil
}

func (e *invertedIndex) tokenize(text string) ([]string, error) {
	// removeStopWords
	text = stopwords.CleanString(text, "en", false)

	// tokenize split texts based on whitespace
	tokens := strings.Fields(text)

	// stemm
	stems := []string{}
	for _, t := range tokens {
		st, err := snowball.Stem(t, "english", false) // TODO: re-assess the stemmer library correctness
		if err != nil {
			return []string{}, err
		}

		stems = append(stems, st)
	}

	return stems, nil
}
