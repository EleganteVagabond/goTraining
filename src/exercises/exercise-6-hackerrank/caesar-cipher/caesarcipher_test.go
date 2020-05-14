package cipher

import "testing"

func TestCaesarCipher(t *testing.T) {
	type args struct {
		raw   string
		shift int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{"", 5},
			want: "",
		},
		{
			name: "no alpha",
			args: args{"-@[]/^@@$%^^&*(!)", 3},
			want: "-@[]/^@@$%^^&*(!)",
		},
		{
			name: "edge cases",
			args: args{"AZaz", 3},
			want: "DCdc",
		},
		{
			name: "edge cases, long shift",
			args: args{"AZaz", 29},
			want: "DCdc",
		},
		{
			name: "edge cases, shift 26",
			args: args{"AZaz", 26},
			want: "AZaz",
		},
		{
			name: "edge cases, shift 25",
			args: args{"AZaz", 25},
			want: "ZYzy",
		},
		{
			name: "alpha with punc",
			args: args{"There's-a-starman-waiting-in-the-sky", 3},
			want: "Wkhuh'v-d-vwdupdq-zdlwlqj-lq-wkh-vnb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaesarCipher(tt.args.raw, tt.args.shift); got != tt.want {
				t.Errorf("CaesarCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
