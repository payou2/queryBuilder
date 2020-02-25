package queryBuilder

import "strings"

func ArgsCount(cnt int) (args string) {
	if cnt <= 0 {
		args = ""
	} else {
		args = "?" + strings.Repeat(",?", cnt-1)
	}

	return args
}
