package bitflyer

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCheckResponse(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "200",
			args: args{
				r: &http.Response{
					StatusCode: 200,
				},
			},
			wantErr: false,
		},
		{
			name: "401",
			args: args{
				r: &http.Response{
					StatusCode: 401,
					Body: ioutil.NopCloser(bytes.NewBufferString(
						`{"status":-500,"error_message":"ACCESS-KEY header is required","data":null}`,
					)),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckResponse(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CheckResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
