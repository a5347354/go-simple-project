package util

import (
	"crypto/aes"
	"testing"
)

func TestAES256Decrypter(t *testing.T) {
	type args struct {
		plaintext string
		key       string
		iv        string
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "AES256Decrypter",
			args: args{
				plaintext: "c353fd22a82acf1a04b2724789e3de0a",
				key:       "szKET12uzhWYzdtBaTbkpS3WlyYLxvHk",
				iv:        "BJvsSswSfhKamPHP",
				blockSize: aes.BlockSize,
			},
			want: "Luke Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AES256Decrypter(tt.args.plaintext, tt.args.key, tt.args.iv, tt.args.blockSize); got != tt.want {
				t.Errorf("AES256Decrypter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAES256(t *testing.T) {
	type args struct {
		plaintext string
		key       string
		iv        string
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "AES256",
			args: args{
				plaintext: "Luke Test",
				key:       "szKET12uzhWYzdtBaTbkpS3WlyYLxvHk",
				iv:        "BJvsSswSfhKamPHP",
				blockSize: aes.BlockSize,
			},
			want: "c353fd22a82acf1a04b2724789e3de0a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AES256(tt.args.plaintext, tt.args.key, tt.args.iv, tt.args.blockSize); got != tt.want {
				t.Errorf("AES256() = %v, want %v", got, tt.want)
			}
		})
	}
}
