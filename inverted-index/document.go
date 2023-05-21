package inverted_index

type xmlDocument struct {
	Documents []document `xml:"doc"`
}

type document struct {
	Title    string `xml:"title"`
	Url      string `xml:"url"`
	Abstract string `xml:"abstract"`
	ID       int
}
