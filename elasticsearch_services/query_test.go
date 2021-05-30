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

		{name: "test123",
			args: args{query: `{"query":{"match_all":{}}}`},
			want: strings.NewReader(`{"query":{"match_all":{}}}`),
		},
		{name: "test102",
			args: args{query: `{"quech_all":"hi"}`},
			want: strings.NewReader(`{"quech_all":"hi"}`),
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

		})
	}
}
