package camelcase

import "testing"

func TestCamelWordCount(t *testing.T) {
	type args struct {
		camelstr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{""},
			want: 0,
		},
		{
			name: "all lower",
			args: args{"everythingislowercase"},
			want: 1,
		},
		{
			name: "1 upper",
			args: args{"upperA"},
			want: 2,
		},
		{
			name: "1 upper full word",
			args: args{"upperArse"},
			want: 2,
		},
		{
			name: "5 uppers",
			args: args{"upperABCDE"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelWordCount(tt.args.camelstr); got != tt.want {
				t.Errorf("CamelWordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
