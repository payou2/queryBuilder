package queryBuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetQuery(t *testing.T) {
	qb := New()
	qb.Select("country.id")
	qb.Select("lang.id as langId")
	qb.Select("country.title as title")
	qb.From("countries as country")
	qb.Where("country.id = 1")
	qb.Where("lang.id = 2")

	s := "SELECT country.id, lang.id as langId, country.title as title "
	f := "FROM countries as country "
	w := "WHERE country.id = 1 AND lang.id = 2"
	expected := s + f + w

	actual := qb.GetQuery()
	assert.Equal(t, expected, actual, "")
}
