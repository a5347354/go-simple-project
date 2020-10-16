package regex

import "testing"

func Test_emailValidation(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Success normal 1",
			args: args{
				email: "luke@linecorp.com",
			},
			wantErr: false,
		},
		{
			name: "Test Success normal 2",
			args: args{
				email: "luke.wang@linecorp.com",
			},
			wantErr: false,
		},
		{
			name: "Test Success normal 3",
			args: args{
				email: "abcd@gmail-yahoo.com",
			},
			wantErr: false,
		},
		{
			name: "Test Success normal 4",
			args: args{
				email: "abcd@gmail.yahoo",
			},
			wantErr: false,
		},
		{
			name: "Test Fail informal 1",
			args: args{
				email: "ç$€§/az@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "Test Fail informal 2",
			args: args{
				email: "abcd@gmailyahoo",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := emailValidation(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("emailValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
