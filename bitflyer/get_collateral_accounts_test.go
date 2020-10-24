package bitflyer

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetCollateralAccountList(t *testing.T) {
	type fields struct {
		num int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetCollateralAccountsOutput
		wantErr bool
	}{
		{
			name: "Empty collateral account",
			fields: fields{
				num: 0,
			},
			want: &GetCollateralAccountsOutput{
				CollateralAccounts: []*CollateralAccount{},
			},
			wantErr: false,
		},
		{
			name: "1 collateral account",
			fields: fields{
				num: 1,
			},
			want: &GetCollateralAccountsOutput{
				CollateralAccounts: []*CollateralAccount{
					{
						CurrencyCode: "JPY",
						Amount: 10000,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2 collateral account",
			fields: fields{
				num: 2,
			},
			want: &GetCollateralAccountsOutput{
				CollateralAccounts: []*CollateralAccount{
					{
						CurrencyCode: "JPY",
						Amount: 10000,
					},
					{
						CurrencyCode: "BTC",
						Amount: 1.23,
					},
				},
			},
			wantErr: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := serverCollateralAccountsMock(tt.fields.num)
			defer srv.Close()
			u, _ := url.Parse(srv.URL)

			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.GetCollateralAccountList(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCollateralAccountList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCollateralAccountList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverCollateralAccountsMock(num int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/me/getcollateralaccounts", func(w http.ResponseWriter, r *http.Request) {
		switch num {
		case 1:
			_, _ = w.Write([]byte(`[
				{
					"currency_code": "JPY",
					"amount": 10000
				}]`,
			))
		case 2:
			_, _ = w.Write([]byte(`[
				{
					"currency_code": "JPY",
					"amount": 10000
				},
				{
					"currency_code": "BTC",
					"amount": 1.23
				}]`,
			))
		default:
			_, _ = w.Write([]byte(`[]`))
		}
	})

	srv := httptest.NewServer(handler)

	return srv
}