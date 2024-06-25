package DIP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDIP(t *testing.T) {
	persons := []Person{
		{"Person1"},
		{"Person2"},
		{"Person3"},
	}
	storage := PersonStorage{persons}

	qm := NewQueryManager(storage)
	result := qm.GetPersonByName("Person1")
	assert.Equal(t, persons[0], *result)
}
