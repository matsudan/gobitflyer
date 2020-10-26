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

func TestClient_GetWithdrawalList(t *testing.T) {
	type fields struct {
		num int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetWithdrawalListOutput
		wantErr bool
	}{
		{
			name: "Empty",
			fields: fields{
				0,
			},
			want: &GetWithdrawalListOutput{
				Withdrawals: []*Withdrawal{},
			},
			wantErr: false,
		},
		{
			name: "1 withdrawal",
			fields: fields{
				1,
			},
			want: &GetWithdrawalListOutput{
				Withdrawals: []*Withdrawal{
					{
						ID:           700,
						OrderID:      "MWD20151020-090909-011111",
						CurrencyCode: "JPY",
						Amount:       12000,
						Status:       types.WithdrawalStatusCompleted,
						EventDate:    "2015-10-20T09:09:09.416",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2 withdrawals",
			fields: fields{
				2,
			},
			want: &GetWithdrawalListOutput{
				Withdrawals: []*Withdrawal{
					{
						ID:           700,
						OrderID:      "MWD20151020-090909-011111",
						CurrencyCode: "JPY",
						Amount:       12000,
						Status:       types.WithdrawalStatusCompleted,
						EventDate:    "2015-10-20T09:09:09.416",
					},
					{
						ID:           701,
						OrderID:      "MWD20151020-090909-022222",
						CurrencyCode: "JPY",
						Amount:       52000,
						Status:       types.WithdrawalStatusPending,
						EventDate:    "2015-10-20T09:09:09.416",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := serverWithdrawalsMock(tt.fields.num)
			defer srv.Close()
			u, _ := url.Parse(srv.URL)

			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.GetWithdrawalList(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWithdrawalList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWithdrawalList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverWithdrawalsMock(num int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/me/getwithdrawals", func(w http.ResponseWriter, r *http.Request) {
		switch num {
		case 1:
			_, _ = w.Write([]byte(`[
				{
    				"id": 700,
    				"order_id": "MWD20151020-090909-011111",
    				"currency_code": "JPY",
					"amount": 12000,
    				"status": "COMPLETED",
    				"event_date": "2015-10-20T09:09:09.416"
				}]`,
			))
		case 2:
			_, _ = w.Write([]byte(`[
				{
    				"id": 700,
    				"order_id": "MWD20151020-090909-011111",
    				"currency_code": "JPY",
					"amount": 12000,
    				"status": "COMPLETED",
    				"event_date": "2015-10-20T09:09:09.416"
				},
				{
    				"id": 701,
    				"order_id": "MWD20151020-090909-022222",
    				"currency_code": "JPY",
					"amount": 52000,
    				"status": "PENDING",
    				"event_date": "2015-10-20T09:09:09.416"
				}]`,
			))
		default:
			_, _ = w.Write([]byte(`[]`))
		}
	})

	srv := httptest.NewServer(handler)

	return srv
}
