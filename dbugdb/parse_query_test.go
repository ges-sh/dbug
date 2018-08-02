package dbugdb

import "testing"

func TestParseQuery(t *testing.T) {
	testCases := []struct {
		query string
		args  []interface{}
		want  string
	}{
		{
			query: "SELECT id, name, elo FROM packages WHERE id = $1 AND name = $2 and created_at < $3",
			args:  []interface{}{1, "test", "2018-05-17"},
			want:  "SELECT id, name, elo FROM packages WHERE id = 1 AND name = test and created_at < 2018-05-17",
		},
	}
	for _, tt := range testCases {
		if query := parseQuery(tt.query, tt.args...); query != tt.want {
			t.Errorf("exp %v\ngot %v", tt.want, query)
		}
	}
}
