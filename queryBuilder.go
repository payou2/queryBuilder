package queryBuilder

import "fmt"

func New() (qb QB) {
	qb = QB{}

	return qb
}

func (qb QB) GetQuery() (str string) {
	str += "SELECT " + get(qb.s, ", ")
	str += " "
	str += "FROM " + qb.f
	str += " "
	str += get(qb.j, " ")
	str += " "
	str += "WHERE " + get(qb.w, " AND ")

	return
}

func (qb *QB) Select(s string) {
	qb.s[getNextKey(&qb.s)] = s
}

func (qb *QB) From(t string, a string) {
	qb.f = fmt.Sprintf("%s as %s", t, a)
}

func (qb *QB) Join(t string, a string, c string) {
	qb.join(JOIN, t, a, c)
}

func (qb *QB) LeftJoin(t string, a string, c string) {
	qb.join(LEFT_JOIN, t, a, c)
}

func (qb *QB) join(jt string, t string, a string, c string) {
	qb.j[getNextKey(&qb.j)] = fmt.Sprintf("%s %s as %s ON %s", jt, t, a, c)
}

func (qb *QB) Where(s string) {
	qb.w[getNextKey(&qb.w)] = s
}

func get(items map[int]string, separator string) (str string) {
	cnt := len(items)
	for i := 1; i < cnt; i++ {
		str = str + items[i] + separator
	}
	str = str + items[cnt]

	return str
}

func getNextKey(items *map[int]string) int {
	key := len(*items) + 1
	if key == 1 {
		*items = make(map[int]string)
	}

	return key
}

func initMap(m *map[string]string) {
	//key := len(*m)
	//if key == 0 {
	if *m == nil {
		*m = make(map[string]string, 10)
	}
}
