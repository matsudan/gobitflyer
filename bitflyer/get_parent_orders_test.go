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

func TestClient_GetParentOrderList(t *testing.T) {
	type fields struct {
		num int
	}
	type args struct {
		parentOrderState *types.ParentOrderState
		paginationQuery  *PaginationQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetParentOrderListOutput
		wantErr bool
	}{
		{
			name: "Empty",
			fields: fields{
				num: 0,
			},
			args: args{
				parentOrderState: nil,
				paginationQuery:  nil,
			},
			want: &GetParentOrderListOutput{
				ParentOrders: []*ParentOrder{},
			},
			wantErr: false,
		},
		{
			name: "1 parent order",
			fields: fields{
				num: 1,
			},
			args: args{
				parentOrderState: nil,
				paginationQuery:  nil,
			},
			want: &GetParentOrderListOutput{
				ParentOrders: []*ParentOrder{
					{
						ID:                      138398,
						ParentOrderID:           "JCO20150707-084555-022523",
						ProductCode:             "BTC_JPY",
						Side:                    types.OrderSideBuy,
						ParentOrderType:         "STOP",
						Price:                   30000,
						AveragePrice:            30000,
						Size:                    0.1,
						ParentOrderState:        types.ParentOrderStateCompleted,
						ExpireDate:              "2015-07-14T07:25:52",
						ParentOrderDate:         "2015-07-07T08:45:53",
						ParentOrderAcceptanceID: "JRF20150707-084552-031927",
						OutstandingSize:         0,
						CancelSize:              0,
						ExecutedSize:            0.1,
						TotalCommission:         0,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2 parent orders",
			fields: fields{
				num: 2,
			},
			args: args{
				parentOrderState: nil,
				paginationQuery:  nil,
			},
			want: &GetParentOrderListOutput{
				ParentOrders: []*ParentOrder{
					{
						ID:                      138398,
						ParentOrderID:           "JCO20150707-084555-022523",
						ProductCode:             "BTC_JPY",
						Side:                    types.OrderSideBuy,
						ParentOrderType:         "STOP",
						Price:                   30000,
						AveragePrice:            30000,
						Size:                    0.1,
						ParentOrderState:        types.ParentOrderStateCompleted,
						ExpireDate:              "2015-07-14T07:25:52",
						ParentOrderDate:         "2015-07-07T08:45:53",
						ParentOrderAcceptanceID: "JRF20150707-084552-031927",
						OutstandingSize:         0,
						CancelSize:              0,
						ExecutedSize:            0.1,
						TotalCommission:         0,
					},
					{
						ID:                      138397,
						ParentOrderID:           "JCO20150707-084549-022519",
						ProductCode:             "BTC_JPY",
						Side:                    types.OrderSideSell,
						ParentOrderType:         "IFD",
						Price:                   30000,
						AveragePrice:            0,
						Size:                    0.1,
						ParentOrderState:        types.ParentOrderStateCancelled,
						ExpireDate:              "2015-07-14T07:25:47",
						ParentOrderDate:         "2015-07-07T08:45:47",
						ParentOrderAcceptanceID: "JRF20150707-084547-396699",
						OutstandingSize:         0,
						CancelSize:              0.1,
						ExecutedSize:            0,
						TotalCommission:         0,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := serverParentOrdersMock(tt.fields.num)
			defer srv.Close()
			u, _ := url.Parse(srv.URL)

			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.GetParentOrderList(context.Background(), tt.args.parentOrderState, tt.args.paginationQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetParentOrderList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetParentOrderList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverParentOrdersMock(num int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/me/getparentorders", func(w http.ResponseWriter, r *http.Request) {
		switch num {
		case 1:
			_, _ = w.Write([]byte(`[
				{
					"id": 138398,
					"parent_order_id": "JCO20150707-084555-022523",
					"product_code": "BTC_JPY",
					"side": "BUY",
					"parent_order_type": "STOP",
					"price": 30000,
					"average_price": 30000,
					"size": 0.1,
					"parent_order_state": "COMPLETED",
					"expire_date": "2015-07-14T07:25:52",
					"parent_order_date": "2015-07-07T08:45:53",
					"parent_order_acceptance_id": "JRF20150707-084552-031927",
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
					"parent_order_id": "JCO20150707-084555-022523",
					"product_code": "BTC_JPY",
					"side": "BUY",
					"parent_order_type": "STOP",
					"price": 30000,
					"average_price": 30000,
					"size": 0.1,
					"parent_order_state": "COMPLETED",
					"expire_date": "2015-07-14T07:25:52",
					"parent_order_date": "2015-07-07T08:45:53",
					"parent_order_acceptance_id": "JRF20150707-084552-031927",
					"outstanding_size": 0,
					"cancel_size": 0,
					"executed_size": 0.1,
					"total_commission": 0
				},
				{
					"id": 138397,
					"parent_order_id": "JCO20150707-084549-022519",
					"product_code": "BTC_JPY",
					"side": "SELL",
					"parent_order_type": "IFD",
					"price": 30000,
					"average_price": 0,
					"size": 0.1,
					"parent_order_state": "CANCELED",
					"expire_date": "2015-07-14T07:25:47",
					"parent_order_date": "2015-07-07T08:45:47",
					"parent_order_acceptance_id": "JRF20150707-084547-396699",
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
