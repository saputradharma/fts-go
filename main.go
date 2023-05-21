package main

type Engine interface {
	LoadDocument(path string) (interface{}, error)
	Search(string) (interface{}, error)
}

func main() {

}
