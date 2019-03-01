package dbugdb

import (
	"fmt"
	"strings"
)

// ParseQuery returns query with dollars replaced by arguments
func ParseQuery(query string, args ...interface{}) string {
	for i, v := range args {
		query = strings.Replace(query, fmt.Sprintf("$%d", i+1), fmt.Sprintf("%v", v), 1)
	}
	return query
}
