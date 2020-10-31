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

func TestClient_GetParentOrder(t *testing.T) {
	type args struct {
		parentOrderID           string
		parentOrderAcceptanceID string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetParentOrderOutput
		wantErr bool
	}{
		{
			name: "Normal",
			args: args{
				parentOrderID:           "JCP20150825-046876-036161",
				parentOrderAcceptanceID: "",
			},
			want: &GetParentOrderOutput{
				ParentOrderDetails{
					ID:            4242,
					ParentOrderID: "JCP20150825-046876-036161",
					OrderMethod:   "IFDOCO",
					ExpireDate:    "2015-09-24T04:35:59.277",
					TimeInForce:   "GTC",
					Parameters: []*Parameter{
						{
							ProductCode:   "BTC_JPY",
							ConditionType: "LIMIT",
							Side:          types.OrderSideBuy,
							Price:         30000,
							Size:          0.1,
							TriggerPrice:  0,
							Offset:        0,
						},
						{
							ProductCode:   "BTC_JPY",
							ConditionType: "LIMIT",
							Side:          types.OrderSideSell,
							Price:         32000,
							Size:          0.1,
							TriggerPrice:  0,
							Offset:        0,
						},
						{
							ProductCode:   "BTC_JPY",
							ConditionType: "STOP_LIMIT",
							Side:          types.OrderSideSell,
							Price:         28800,
							Size:          0.1,
							TriggerPrice:  29000,
							Offset:        0,
						},
					},
					ParentOrderAcceptanceID: "JRF20150925-060559-396699",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid argument 1",
			args: args{
				parentOrderID:           "JRF20150925-060559-396699",
				parentOrderAcceptanceID: "JRF20150925-060559-396699",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid argument 2",
			args: args{
				parentOrderID:           "",
				parentOrderAcceptanceID: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := serverParentOrderMock()
			defer srv.Close()
			u, _ := url.Parse(srv.URL)

			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.GetParentOrder(context.Background(), tt.args.parentOrderID, tt.args.parentOrderAcceptanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetParentOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetParentOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverParentOrderMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/me/getparentorder", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`
			{
				"id": 4242,
				"parent_order_id": "JCP20150825-046876-036161",
				"order_method": "IFDOCO",
				"expire_date": "2015-09-24T04:35:59.277",
				"time_in_force": "GTC",
				"parameters": [
					{
						"product_code": "BTC_JPY",
						"condition_type": "LIMIT",
						"side": "BUY",
						"price": 30000,
						"size": 0.1,
						"trigger_price": 0,
						"offset": 0
					},
					{
						"product_code": "BTC_JPY",
						"condition_type": "LIMIT",
						"side": "SELL",
						"price": 32000,
						"size": 0.1,
						"trigger_price": 0,
						"offset": 0
					},
					{
						"product_code": "BTC_JPY",
						"condition_type": "STOP_LIMIT",
						"side": "SELL",
						"price": 28800,
						"size": 0.1,
						"trigger_price": 29000,
						"offset": 0
					}
				],
				"parent_order_acceptance_id": "JRF20150925-060559-396699"
			}`,
		))
	})

	srv := httptest.NewServer(handler)

	return srv
}
