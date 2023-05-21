package main

import "testing"

func IGNORETestSearch(t *testing.T) {
	data, _ := LoadDocument(dataPath)
	keyWords := []string{
		"Country",
		"Harry Potter",
		"Indonesia",
	}

	for _, k := range keyWords {
		search(data, k)
	}
}
