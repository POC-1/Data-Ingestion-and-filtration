package setup

import (
	"log"
	"reflect"
	"testing"
)

func TestGetJsonByteVal(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "Test101",
			args: args{
				filePath: `/Users/himalisaini/Desktop/internship/Go/src/github.com/himalisaini/poc1/json files/students.json`,
			},
			want: []byte{32, 91, 10, 32, 32, 32, 32, 32, 123, 10, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 82, 97, 104, 117, 108, 34, 44, 10, 32, 34, 105, 100, 34, 58, 32, 49, 50, 51, 44, 10, 32, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 32, 123, 10, 32, 34, 115, 116, 114, 101, 101, 116, 34, 58, 32, 34, 83, 101, 110, 97, 112, 97, 116, 105, 32, 66, 97, 112, 97, 116, 32, 82, 111, 97, 100, 34, 44, 10, 32, 34, 104, 111, 117, 115, 101, 110, 111, 34, 58, 32, 49, 44, 10, 32, 34, 99, 105, 116, 121, 34, 58, 32, 34, 80, 117, 110, 101, 34, 10, 32, 125, 44, 10, 32, 34, 100, 101, 112, 116, 34, 58, 32, 34, 67, 111, 109, 112, 117, 116, 101, 114, 32, 83, 99, 105, 101, 110, 99, 101, 34, 44, 10, 32, 34, 99, 111, 110, 116, 97, 99, 116, 34, 58, 32, 123, 10, 32, 34, 112, 114, 105, 109, 97, 114, 121, 34, 58, 32, 57, 49, 48, 48, 48, 48, 48, 48, 48, 48, 44, 10, 32, 34, 115, 101, 99, 111, 110, 100, 97, 114, 121, 34, 58, 32, 57, 49, 48, 48, 48, 48, 48, 49, 48, 10, 32, 125, 10, 32, 125, 44, 10, 123, 10, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 83, 97, 109, 105, 114, 34, 44, 10, 32, 34, 105, 100, 34, 58, 32, 49, 50, 52, 44, 10, 32, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 32, 123, 10, 32, 34, 115, 116, 114, 101, 101, 116, 34, 58, 32, 34, 71, 97, 110, 100, 104, 105, 32, 82, 111, 97, 100, 34, 44, 10, 32, 34, 104, 111, 117, 115, 101, 110, 111, 34, 58, 32, 49, 44, 10, 32, 34, 99, 105, 116, 121, 34, 58, 32, 34, 65, 104, 109, 101, 100, 97, 98, 97, 100, 34, 10, 32, 125, 44, 10, 32, 34, 100, 101, 112, 116, 34, 58, 32, 34, 67, 111, 109, 112, 117, 116, 101, 114, 32, 65, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 34, 44, 10, 32, 34, 99, 111, 110, 116, 97, 99, 116, 34, 58, 32, 123, 10, 32, 34, 112, 114, 105, 109, 97, 114, 121, 34, 58, 32, 57, 49, 48, 48, 48, 48, 48, 48, 48, 49, 44, 10, 32, 34, 115, 101, 99, 111, 110, 100, 97, 114, 121, 34, 58, 32, 57, 49, 48, 48, 48, 48, 48, 49, 49, 10, 32, 125, 10, 125, 44, 10, 123, 10, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 74, 117, 108, 105, 101, 34, 44, 10, 32, 34, 105, 100, 34, 58, 32, 49, 50, 53, 44, 10, 32, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 32, 123, 10, 32, 34, 115, 116, 114, 101, 101, 116, 34, 58, 32, 34, 120, 121, 122, 32, 115, 116, 114, 101, 101, 34, 44, 10, 32, 34, 104, 111, 117, 115, 101, 110, 111, 34, 58, 32, 53, 44, 10, 32, 34, 99, 105, 116, 121, 34, 58, 32, 34, 80, 117, 110, 101, 100, 34, 10, 32, 125, 44, 10, 32, 34, 100, 101, 112, 116, 34, 58, 32, 34, 67, 111, 109, 112, 117, 116, 101, 114, 32, 84, 101, 99, 104, 110, 111, 108, 111, 103, 121, 34, 44, 10, 32, 34, 99, 111, 110, 116, 97, 99, 116, 34, 58, 32, 123, 10, 32, 34, 112, 114, 105, 109, 97, 114, 121, 34, 58, 32, 57, 49, 48, 48, 48, 48, 48, 48, 48, 51, 44, 10, 32, 34, 115, 101, 99, 111, 110, 100, 97, 114, 121, 34, 58, 32, 57, 49, 48, 48, 48, 48, 48, 52, 49, 10, 32, 125, 10, 32, 125, 10, 93},
		}, //Valid Json File
		{name: "Test102",
			args: args{
				filePath: `/Users/himalisaini/Desktop/internship/Go/src`,
			},
			want: []byte(""),
		}, //Invalid Path to Json File
		{name: "Test103",
			args: args{
				filePath: `/Users/`,
			},
			want: []byte(""),
		}, //Invalid Path to Json File
		{name: "Test104",
			args: args{
				filePath: `/Users/himalisaini/Desktop/internship/Go/src/github.com/himalisaini/poc1/util/config.go`,
			},
			want: []byte(""),
		}, //Not Json File
		{name: "Test105",
			args: args{
				filePath: `/Users/himalisaini/Desktop/internship/Go/src/github.com/himalisaini/poc1/json files/test.json`,
			},
			want: []byte{91, 10, 32, 32, 32, 32, 123, 10, 32, 32, 32, 32, 34, 102, 114, 117, 105, 116, 34, 58, 32, 34, 65, 112, 112, 108, 101, 34, 44, 10, 32, 32, 32, 32, 34, 115, 105, 122, 101, 34, 58, 32, 34, 76, 97, 114, 103, 101, 34, 44, 10, 32, 32, 32, 32, 34, 99, 111, 108, 111, 114, 34, 58, 32, 34, 82, 101, 100, 34, 10, 32, 32, 32, 32, 125, 10, 93},
		}, //Valid Path Json File
		{name: "Test106",
			args: args{
				filePath: `/Users/himalisaini/Desktop/internship/Go/src/github.com/himalisaini/poc1/json files/incorrect.json`,
			},
			want: []byte(""),
		}, //Invalid Json File
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetJsonByteVal(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJsonByteVal() = %v, want %v", got, tt.want)
			}
			log.Println("Test Case : ", tt.name, " successfully passed.")
			println()
			println()
		})
	}
}
