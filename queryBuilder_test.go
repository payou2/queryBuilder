package queryBuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSelectQuery(t *testing.T) {
	qb := New()
	qb.Select("country.id")
	qb.Select("lang.id as langId")
	qb.Select("country.title as title")
	qb.From("countries", "country")
	qb.Join("languages", "lang", "country.code=lang.code")
	qb.LeftJoin("phones", "phone", "country.code=phone.code")
	qb.Where("country.id = 1")
	qb.Where("lang.id = 2")

	s := "SELECT country.id, lang.id as langId, country.title as title "
	f := "FROM countries as country "
	j := "JOIN languages as lang ON country.code=lang.code "
	lj := "LEFT JOIN phones as phone ON country.code=phone.code "
	w := "WHERE country.id = 1 AND lang.id = 2;"
	expected := s + f + j + lj + w
	actual := qb.GetQuery()
	assert.Equal(t, expected, actual, "")
}

func TestGetSelectQueryAll(t *testing.T) {
	qb := New()
	qb.Select("id")
	qb.Select("title as title")
	qb.From("countries", "")

	s := "SELECT id, title as title "
	f := "FROM countries "
	w := ";"
	expected := s + f + w
	actual := qb.GetQuery()
	assert.Equal(t, expected, actual, "")
}

func TestGetInsertQuery(t *testing.T) {
	qb := New()
	qb.Insert("code")
	qb.Insert("title")
	qb.Into("countries", "")
	qb.Values("'ua'")
	qb.Values("'Ukraine'")

	i := "INSERT INTO countries (code, title) "
	v := "VALUES ('ua', 'Ukraine');"

	expected := i + v
	actual := qb.GetQuery()
	assert.Equal(t, expected, actual, "")
}
