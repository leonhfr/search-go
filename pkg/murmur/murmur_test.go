package murmur

import (
	"testing"
)

var (
	text1 = "The quick brown fox jumps over the lazy dog"
	text2 = "The fox jumps cog"
)

func Benchmark_Sum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum(text1, 1)
	}
}

func Test_Sum(t *testing.T) {
	type args struct {
		text string
		seed uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{text1, args{text1, 1}, uint32(0x78e69e27)},
		{text2, args{text2, 1}, uint32(3709117227)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.text, tt.args.seed); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
