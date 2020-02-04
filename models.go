package queryBuilder

type (
	QB struct {
		fields  map[int]string
		table   string
		j       map[int]string
		w       map[int]string
		g       map[int]string
		o       map[int]string
		aliases map[string]string

		qt  int
		val map[int]string
	}
)

const (
	SELECT_TYPE = 1
	INSERT_TYPE = 2
	UPDATE_TYPE = 3
	DELETE_TYPE = 4
	LEFT_JOIN   = "LEFT JOIN"
	JOIN        = "JOIN"
)
