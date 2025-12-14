package ignite

import (
	"bytes"
	"bufio"
	"testing"
)

func Test_response_ReadFrom(t *testing.T) {
	rr := bytes.NewBuffer([]byte{1, 0, 0, 0, 1})

	type args struct {
		rr *bufio.Reader
	}
	tests := []struct {
		name    string
		r       *response
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "1",
			r:    &response{},
			args: args{
				rr: bufio.NewReader(rr),
			},
			want: 4 + 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadFrom(tt.args.rr)
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
