package norm

import (
	"exercises/7-phnorm/norm/db"
	"reflect"
	"testing"
)

func Test_NormalizeAll(t *testing.T) {
	type args struct {
		nums []db.PhoneNumber
	}
	tests := []struct {
		name string
		args args
		want []db.PhoneNumber
	}{
		{
			name: "one correct",
			args: args{nums: []db.PhoneNumber{{0, "1234567890", false, false}}},
			want: []db.PhoneNumber{{0, "1234567890", false, false}},
		},
		{
			name: "one with extra chars",
			args: args{nums: []db.PhoneNumber{{0, "1 2 3 45 6 7 8 9 0 #@&*(!)^", false, false}}},
			want: []db.PhoneNumber{{0, "1234567890", true, false}},
		},
		{
			name: "2, no duplicates",
			args: args{nums: []db.PhoneNumber{{0, "1234567890", false, false}, {0, "1234567899", false, false}}},
			want: []db.PhoneNumber{{0, "1234567890", false, false}, {0, "1234567899", false, false}},
		},
		{
			name: "2, one duplicate",
			args: args{nums: []db.PhoneNumber{{0, "1234567890", false, false}, {0, "1234567890", false, false}}},
			want: []db.PhoneNumber{{0, "1234567890", false, false}, {0, "1234567890", false, true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeAll(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
