package bloom

import (
	"testing"
)

func Test_optimalBits(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{"400", 400, 5752},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalBits(tt.args); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
