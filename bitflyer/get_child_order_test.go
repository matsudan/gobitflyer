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

func TestClient_GetChildOrderList(t *testing.T) {
	type fields struct {
		num int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetChildOrderListOutput
		wantErr bool
	}{
		{
			name: "Empty",
			fields: fields{
				num: 0,
			},
			want: &GetChildOrderListOutput{
				ChildOrders: []*ChildOrder{},
			},
		},
		{
			name: "1 child order",
			fields: fields{
				num: 1,
			},
			want: &GetChildOrderListOutput{
				ChildOrders: []*ChildOrder{
					{
						ID:                     138398,
						ChildOrderID:           "JOR20150707-084555-022523",
						ProductCode:            "BTC_JPY",
						Side:                   types.OrderSideBuy,
						ChildOrderType:         types.ChildOrderTypeLimit,
						Price:                  30000,
						AveragePrice:           30000,
						Size:                   0.1,
						ChildOrderState:        types.ChildOrderStateCompleted,
						ExpireDate:             "2015-07-14T07:25:52",
						ChildOrderDate:         "2015-07-07T08:45:53",
						ChildOrderAcceptanceID: "JRF20150707-084552-031927",
						OutstandingSize:        0,
						CancelSize:             0,
						ExecutedSize:           0.1,
						TotalCommission:        0,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2 child order2",
			fields: fields{
				num: 2,
			},
			want: &GetChildOrderListOutput{
				ChildOrders: []*ChildOrder{
					{
						ID:                     138398,
						ChildOrderID:           "JOR20150707-084555-022523",
						ProductCode:            "BTC_JPY",
						Side:                   types.OrderSideBuy,
						ChildOrderType:         types.ChildOrderTypeLimit,
						Price:                  30000,
						AveragePrice:           30000,
						Size:                   0.1,
						ChildOrderState:        types.ChildOrderStateCompleted,
						ExpireDate:             "2015-07-14T07:25:52",
						ChildOrderDate:         "2015-07-07T08:45:53",
						ChildOrderAcceptanceID: "JRF20150707-084552-031927",
						OutstandingSize:        0,
						CancelSize:             0,
						ExecutedSize:           0.1,
						TotalCommission:        0,
					},
					{
						ID:                     138397,
						ChildOrderID:           "JOR20150707-084549-022519",
						ProductCode:            "BTC_JPY",
						Side:                   types.OrderSideSell,
						ChildOrderType:         types.ChildOrderTypeLimit,
						Price:                  30000,
						AveragePrice:           0,
						Size:                   0.1,
						ChildOrderState:        types.ChildOrderStateCancelled,
						ExpireDate:             "2015-07-14T07:25:47",
						ChildOrderDate:         "2015-07-07T08:45:47",
						ChildOrderAcceptanceID: "JRF20150707-084547-396699",
						OutstandingSize:        0,
						CancelSize:             0.1,
						ExecutedSize:           0,
						TotalCommission:        0,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := serverChildOrdersMock(tt.fields.num)
			defer srv.Close()
			u, _ := url.Parse(srv.URL)

			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.GetChildOrderList(context.Background(), nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChildOrderList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChildOrderList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverChildOrdersMock(num int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/me/getchildorders", func(w http.ResponseWriter, r *http.Request) {
		switch num {
		case 1:
			_, _ = w.Write([]byte(`[
				{
					"id": 138398,
					"child_order_id": "JOR20150707-084555-022523",
					"product_code": "BTC_JPY",
					"side": "BUY",
					"child_order_type": "LIMIT",
					"price": 30000,
					"average_price": 30000,
					"size": 0.1,
					"child_order_state": "COMPLETED",
					"expire_date": "2015-07-14T07:25:52",
					"child_order_date": "2015-07-07T08:45:53",
					"child_order_acceptance_id": "JRF20150707-084552-031927",
					"outstanding_size": 0,
					"cancel_size": 0,
					"executed_size": 0.1,
					"total_commission": 0
				}]`,
			))
		case 2:
			_, _ = w.Write([]byte(`[
				{
					"id": 138398,
					"child_order_id": "JOR20150707-084555-022523",
					"product_code": "BTC_JPY",
					"side": "BUY",
					"child_order_type": "LIMIT",
					"price": 30000,
					"average_price": 30000,
					"size": 0.1,
					"child_order_state": "COMPLETED",
					"expire_date": "2015-07-14T07:25:52",
					"child_order_date": "2015-07-07T08:45:53",
					"child_order_acceptance_id": "JRF20150707-084552-031927",
					"outstanding_size": 0,
					"cancel_size": 0,
					"executed_size": 0.1,
					"total_commission": 0
				},
				{
					"id": 138397,
					"child_order_id": "JOR20150707-084549-022519",
					"product_code": "BTC_JPY",
					"side": "SELL",
					"child_order_type": "LIMIT",
					"price": 30000,
					"average_price": 0,
					"size": 0.1,
					"child_order_state": "CANCELED",
					"expire_date": "2015-07-14T07:25:47",
					"child_order_date": "2015-07-07T08:45:47",
					"child_order_acceptance_id": "JRF20150707-084547-396699",
					"outstanding_size": 0,
					"cancel_size": 0.1,
					"executed_size": 0,
					"total_commission": 0
				}]`,
			))
		default:
			_, _ = w.Write([]byte(`[]`))
		}
	})

	srv := httptest.NewServer(handler)

	return srv
}
