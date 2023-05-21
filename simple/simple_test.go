package simple

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	engine *simpleFts
}

func (s *testSuite) SetupTest() {
	e, err := NewSimpleFts("./../data/enwikinews-20230501-abstract.xml")
	s.engine = e
	s.NoError(err)
}

func (s *testSuite) TestSearch() {
	keyWords := []string{
		"Marvel",
		"Movie",
		"Stan Lee",
	}

	for _, k := range keyWords {
		s.engine.Search(k)
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func BenchmarkSimpleFts(b *testing.B) {
	keyWords := []string{
		"Marvel",
		"Movie",
		"Stan Lee",
	}

	s := new(testSuite)
	s.SetT(&testing.T{})
	// loading data
	s.SetupTest()

	for _, k := range keyWords {
		// reset timer to limit to
		// search operation
		b.ResetTimer()
		b.Run(fmt.Sprintf("Searching Keyword_%s", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				s.engine.Search(k)
				b.StopTimer()
			}
		})
	}

}
