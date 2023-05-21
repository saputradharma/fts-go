# (WIP) Full-text Search Go Implementation

Searching an information is one of most important operation in computer science. Finding relevant information in a bunch of news article, books chapters, academic papers just some of the the examples.
 
This is an exploration to implement Full-text search in Go with different techniques:
- Simple approach using string.Contains
- Inverted Index
- TF-IDF 
- Vector Search

Each of the techniques will be explained and we will bechmark the performance of the approach. For the simplicity we will limit the operations to two big parts: Loading the data (loading, transforming, indexing, etc.) and Searching the data.
We will use [wikipedia dump data](https://dumps.wikimedia.org/enwikinews/20230501/enwikinews-20230501-abstract.xml.gz) as sample document data source. Let's begin with the easiest technique by using string.Contains.

## Simple approach using string.Contains




## References
- https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/
- https://weaviate.io/blog/why-is-vector-search-so-fast
- https://github.com/bbalet/stopwords
- https://github.com/go-nlp/tfidf


