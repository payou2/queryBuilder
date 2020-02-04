package queryBuilder

import (
	"fmt"
	"strings"
)

func New() (qb QB) {
	qb = QB{}

	return qb
}

func (qb QB) GetQuery() (str string) {
	switch qb.qt {
	case SELECT_TYPE:
		str = qb.GetSelectQuery()
	case INSERT_TYPE:
		str = qb.GetInsertQuery()
	default:
		str = ""
	}

	return str
}

func (qb QB) GetInsertQuery() (str string) {
	str += "INSERT"
	str += " "
	str += "INTO " + qb.table
	str += get(qb.j, " ")
	str += " ("
	str += get(qb.fields, ", ")
	str += ")"
	str += " "
	str += "VALUES ("
	str += get(qb.val, ", ")
	str += ")"
	if w := get(qb.w, " AND "); w != "" {
		str += " WHERE " + w
	}

	str += ";"

	return str
}

func (qb QB) GetSelectQuery() (str string) {
	str += "SELECT " + get(qb.fields, ", ")
	str += " "
	str += "FROM " + qb.table
	str += " "
	str += get(qb.j, " ")
	if w := get(qb.w, " AND "); w != "" {
		str += " WHERE " + w
	}

	str += ";"

	return str
}

func (qb *QB) Into(t string, a string) {
	qb.qt = INSERT_TYPE
	qb.From(t, a)
}

func (qb *QB) Insert(s string) {
	qb.fields[getNextKey(&qb.fields)] = s
}
func (qb *QB) Values(val string) {
	qb.val[getNextKey(&qb.val)] = val
}

func (qb *QB) Select(s string) {
	qb.qt = SELECT_TYPE
	qb.fields[getNextKey(&qb.fields)] = s
}

func (qb *QB) From(t string, a string) {
	if strings.TrimSpace(a) == "" {
		qb.table = t
	} else {
		qb.table = fmt.Sprintf("%s as %s", t, a)

	}
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
