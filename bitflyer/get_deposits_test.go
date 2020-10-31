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

func TestClient_GetDepositList(t *testing.T) {
	type fields struct {
		num int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetDepositListOutput
		wantErr bool
	}{
		{
			name: "Empty deposit",
			fields: fields{
				num: 0,
			},
			want: &GetDepositListOutput{
				Deposits: []*Deposit{},
			},
			wantErr: false,
		},
		{
			name: "1 deposit",
			fields: fields{
				num: 1,
			},
			want: &GetDepositListOutput{
				Deposits: []*Deposit{
					{
						ID:           300,
						OrderID:      "MDP20151014-101010-033333",
						CurrencyCode: "JPY",
						Amount:       10000,
						Status:       types.DepositStatusPending,
						EventDate:    "2015-10-14T10:10:10.001",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2 deposit",
			fields: fields{
				num: 2,
			},
			want: &GetDepositListOutput{
				Deposits: []*Deposit{
					{
						ID:           300,
						OrderID:      "MDP20151014-101010-033333",
						CurrencyCode: "JPY",
						Amount:       10000,
						Status:       types.DepositStatusPending,
						EventDate:    "2015-10-14T10:10:10.001",
					},
					{
						ID:           301,
						OrderID:      "MDP20151014-101010-099999",
						CurrencyCode: "JPY",
						Amount:       30000,
						Status:       types.DepositStatusCompleted,
						EventDate:    "2015-10-14T10:10:10.001",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := serverDepositsMock(tt.fields.num)
			defer srv.Close()
			u, _ := url.Parse(srv.URL)

			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.GetDepositList(context.Background(), nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDepositList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDepositList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverDepositsMock(num int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/me/getdeposits", func(w http.ResponseWriter, r *http.Request) {
		switch num {
		case 1:
			_, _ = w.Write([]byte(`[
				{
					"id": 300,
					"order_id": "MDP20151014-101010-033333",
					"currency_code": "JPY",
					"amount": 10000,
					"status": "PENDING",
					"event_date": "2015-10-14T10:10:10.001"
				}]`,
			))
		case 2:
			_, _ = w.Write([]byte(`[
				{
					"id": 300,
					"order_id": "MDP20151014-101010-033333",
					"currency_code": "JPY",
					"amount": 10000,
					"status": "PENDING",
					"event_date": "2015-10-14T10:10:10.001"
				},
				{
					"id": 301,
					"order_id": "MDP20151014-101010-099999",
					"currency_code": "JPY",
					"amount": 30000,
					"status": "COMPLETED",
					"event_date": "2015-10-14T10:10:10.001"
				}]`,
			))
		default:
			_, _ = w.Write([]byte(`[]`))
		}
	})

	srv := httptest.NewServer(handler)

	return srv
}
