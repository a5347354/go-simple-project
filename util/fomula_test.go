package util

import "testing"

func TestRoundValue(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Rounding Test 1",
			args: args{
				value: 32.44444,
			},
			want: 32,
		},
		{
			name: "Rounding Test 2",
			args: args{
				value: 32.71111,
			},
			want: 33,
		},
		{
			name: "Rounding Test 3",
			args: args{
				value: 32.5222,
			},
			want: 33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundValue(tt.args.value); got != tt.want {
				t.Errorf("RoundValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
