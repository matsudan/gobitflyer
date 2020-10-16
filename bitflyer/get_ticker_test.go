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

func TestClient_GetTicker(t *testing.T) {
	srv := serverTickerMock()
	defer srv.Close()
	u, _ := url.Parse(srv.URL)

	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetTickerOutput
		wantErr bool
	}{
		{
			name: "Normal",
			args: args{
				"BTC_JPY",
			},
			want: &GetTickerOutput{
				ProductCode:     "BTC_JPY",
				State:           types.StateRunning,
				Timestamp:       "2015-07-08T02:50:59.97",
				TickID:          3579,
				BestBid:         30000,
				BestAsk:         36640,
				BestBidSize:     0.1,
				BestAskSize:     5,
				TotalBidDepth:   15.13,
				TotalAskDepth:   20,
				MarketBidSize:   0,
				MarketAskSize:   0,
				LTP:             31690,
				Volume:          16819.26,
				VolumeByProduct: 6819.26,
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
			got, err := c.GetTicker(context.Background(), tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverTickerMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/ticker", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`
{"product_code": "BTC_JPY",
  "state": "RUNNING",
  "timestamp": "2015-07-08T02:50:59.97",
  "tick_id": 3579,
  "best_bid": 30000,
  "best_ask": 36640,
  "best_bid_size": 0.1,
  "best_ask_size": 5,
  "total_bid_depth": 15.13,
  "total_ask_depth": 20,
  "market_bid_size": 0,
  "market_ask_size": 0,
  "ltp": 31690,
  "volume": 16819.26,
  "volume_by_product": 6819.26
}`))
	})

	srv := httptest.NewServer(handler)

	return srv
}
