package queryBuilder

import "strings"

func New() (qb QB) {
	qb = QB{}

	return qb
}

func (qb QB) GetQuery() (str string) {
	str += "SELECT " + get(qb.s, ", ")
	str += " "
	str += "FROM " + qb.f
	str += " "
	//str += get(qb.j, " ")
	str += " "
	str += "WHERE " + get(qb.w, " AND ")

	return
}

func (qb *QB) Select(s string) {
	qb.s[getNextKey(&qb.s)] = s
}
func (qb *QB) Where(s string) {
	qb.w[getNextKey(&qb.w)] = s
}

func (qb *QB) From(s string) {
	qb.f = s
}

func initMap(m *map[string]string) {
	//key := len(*m)
	//if key == 0 {
	if *m == nil {
		*m = make(map[string]string, 10)
	}
}

func get(items map[int]string, separator string) (str string) {
	for _, item := range items {
		str = str + item + separator
	}
	str = strings.TrimRight(str, separator)
	return str
}

func getNextKey(items *map[int]string) int {
	key := len(*items) + 1
	if key == 1 {
		*items = make(map[int]string)
	}

	return key
}
