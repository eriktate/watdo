package postgres

import (
	"fmt"
	"io/ioutil"
)

func getQuery(name string) string {
	filename := fmt.Sprintf("queries/%s.sql", name)
	query, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("%s does not exist: %s", filename, err))
	}

	return string(query)
}
