package bitflyer

import (
	"context"
	"github.com/matsudan/gobitflyer/bitflyer/types"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetBoardState(t *testing.T) {
	srv := serverBoardStatehMock()
	defer srv.Close()
	u, _ := url.Parse(srv.URL)

	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetBoardStateOutput
		wantErr bool
	}{
		{
			name: "Normal",
			args: args{
				productCode: "BTC_JPY",
			},
			want: &GetBoardStateOutput{
				Health: types.HealthNormal,
				State: types.StateMatured,
				Data: BoardStateData{
					SpecialQuotation: 410897,
				},
			},
			wantErr: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.GetBoardState(context.Background(), tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoardState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBoardState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverBoardStatehMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/getboardstate", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"health": "NORMAL", "state": "MATURED", "data": {"special_quotation": 410897}}`))
	})

	srv := httptest.NewServer(handler)

	return srv
}