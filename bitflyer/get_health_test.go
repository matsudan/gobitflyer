package bitflyer

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/matsudan/gobitflyer/bitflyer/types"
)

func TestClient_GetHealth(t *testing.T) {
	srv := serverHealthMock()
	defer srv.Close()
	u, _ := url.Parse(srv.URL)

	tests := []struct {
		name    string
		want    *GetHealthOutput
		wantErr bool
	}{
		{
			name: "Normal",
			want: &GetHealthOutput{
				Status: types.ExchangeStatus("NORMAL"),
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
			got, err := c.GetHealth(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHealth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHealth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverHealthMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/gethealth", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"Status": "NORMAL"}`))
	})

	srv := httptest.NewServer(handler)

	return srv
}
