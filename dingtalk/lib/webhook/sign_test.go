package webhook

import (
	"testing"
)

func Test_getSign(t *testing.T) {
	type args struct {
		timestamp int64
		secret    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{ "demo1", args{1658047001719, "abcdef"}, "LwEbrSBm2YdHvsPI/KpKQQRLqYaAtwhjdORXPYjCxQY=" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSign(tt.args.timestamp, tt.args.secret); got != tt.want {
				t.Errorf("getSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
