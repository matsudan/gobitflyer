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

func TestClient_SendChildOrder(t *testing.T) {
	tests := []struct {
		name    string
		args    *SendChildOrderInput
		want    *SendChildOrderOutput
		wantErr bool
	}{
		{
			name: "",
			args: &SendChildOrderInput{
				ProductCode:    "BTC_JPY",
				ChildOrderType: types.ChildOrderTypeLimit,
				Side:           types.OrderSideBuy,
				Price:          30000,
				Size:           0.1,
				MinuteToExpire: 10000,
				TimeInForce:    types.TimeInForceGTC,
			},
			want: &SendChildOrderOutput{
				ChildOrderAcceptanceID: "JRF20150707-050237-639234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := serverChildOrderMock()
			defer srv.Close()
			u, _ := url.Parse(srv.URL)
			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.SendChildOrder(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendChildOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendChildOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverChildOrderMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/me/sendchildorder", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"child_order_acceptance_id": "JRF20150707-050237-639234"}`))
	})

	srv := httptest.NewServer(handler)

	return srv
}
