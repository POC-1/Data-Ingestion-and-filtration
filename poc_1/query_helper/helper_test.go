package query_helper

import (
	"reflect"
	"strings"
	"testing"
)

func TestConstructQuery(t *testing.T) {
	type args struct {
		q     string
		size  int
		score int
	}
	tests := []struct {
		name string
		args args
		want *strings.Reader
	}{
		// TODO: Add test cases.
		{ name: "test123",
		args:args{q: `"match" : {"dept": "Computer Science"}`,
				  size: 10,
				  score: 1 },
				  want: strings.NewReader(`{"min_score":1, "query": {"match" : {"dept": "Computer Science"}}, "size": 10}`),
		},
		{ name: "test121",
		args:args{q: `"match" : {"dept": "Computer Application"}`,
				  size: 10,
				  score: 1 },
				  want: strings.NewReader(`{"min_score":1, "query": {"match" : {"dept": "Computer Application"}}, "size": 10}`),
		},
		{ name: "test122",
		args:args{q: `"match" : {"dept": "Computer"}`,
				  size: 10,
				  score: 0 },
				  want: strings.NewReader(`{"min_score":0, "query": {"match" : {"dept": "Computer"}}, "size": 10}`),
		},
		{ name: "test123",
		args:args{q: `"match" : {"dept": "Application"}`,
				  size: 100,
				  score: 1 },
				  want: strings.NewReader(`{"min_score":1, "query": {"match" : {"dept": "Application"}}, "size": 100}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstructQuery(tt.args.q, tt.args.size, tt.args.score); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConstructQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
