package elasticsearch_services

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestConstructQuery(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want *strings.Reader
	}{

		{name: "test101",
			args: args{query: `{"query":{"match_all":{}}}`},
			want: strings.NewReader(`{"query":{"match_all":{}}}`),
		},
		{name: "test102",
			args: args{query: `{
				"query": {
				  "nested": {
					"path": "address",
					"query": {
					  "bool": {
						"must": [
						  { "match": { "address.city": "Pune" } }
						]
					  }
					},
					"score_mode": "avg"
				  }
				}
			  }`},
			want: strings.NewReader(`{
				"query": {
				  "nested": {
					"path": "address",
					"query": {
					  "bool": {
						"must": [
						  { "match": { "address.city": "Pune" } }
						]
					  }
					},
					"score_mode": "avg"
				  }
				}
			  }`),
		},

		{name: "test103",
			args: args{query: `{
				"min_score": 1,
				"query" : {
					"match" : { "dept" : "Computer Science" }
				}
			}`},
			want: strings.NewReader(`{
				"min_score": 1,
				"query" : {
					"match" : { "dept" : "Computer Science" }
				}
			}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConstructQuery(tt.args.query)
			if !reflect.DeepEqual(got, tt.want) {
				log.Printf("%T", got)
				log.Printf("%T", tt.want)
				t.Errorf("ConstructQuery() got = %v, want %v", got, tt.want)
			}
			log.Println("Test Case : ", tt.name, " successfully passed.")
			println()
			println()

		})
	}
}
