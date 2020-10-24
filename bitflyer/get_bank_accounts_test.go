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

func TestClient_GetBankAccountList(t *testing.T) {
	type fields struct {
		num int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetBankAccountListOutput
		wantErr bool
	}{
		{
			name: "1 account",
			fields: fields{
				num: 1,
			},
			want: &GetBankAccountListOutput{
				BankAccounts: []*BankAccount{
					{
						ID:            3402,
						IsVerified:    true,
						BankName:      "二菱東京UFJ",
						BranchName:    "人形町支店",
						AccountType:   "普通",
						AccountNumber: "1111111",
						AccountName:   "タナカ",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2 account",
			fields: fields{
				num: 2,
			},
			want: &GetBankAccountListOutput{
				BankAccounts: []*BankAccount{
					{
						ID:            3402,
						IsVerified:    true,
						BankName:      "二菱東京UFJ",
						BranchName:    "人形町支店",
						AccountType:   "普通",
						AccountNumber: "1111111",
						AccountName:   "タナカ",
					},
					{
						ID:            3403,
						IsVerified:    true,
						BankName:      "二井住友銀行",
						BranchName:    "人形町支店",
						AccountType:   "普通",
						AccountNumber: "1111115",
						AccountName:   "タナカ",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := serverBankAccountsMock(tt.fields.num)
			defer srv.Close()
			u, _ := url.Parse(srv.URL)

			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.GetBankAccountList(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBankAccountList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBankAccountList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverBankAccountsMock(num int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/me/getbankaccounts", func(w http.ResponseWriter, r *http.Request) {
		if num == 1 {
			_, _ = w.Write([]byte(`[{
				"id": 3402,
				"is_verified": true,
				"bank_name": "二菱東京UFJ",
				"branch_name": "人形町支店",
				"account_type": "普通",
				"account_number": "1111111",
				"account_name": "タナカ"}]`,
			))
		}
		if num == 2 {
			_, _ = w.Write([]byte(`[
				{
					"id": 3402,
					"is_verified": true,
					"bank_name": "二菱東京UFJ",
					"branch_name": "人形町支店",
					"account_type": "普通",
					"account_number": "1111111",
					"account_name": "タナカ"
				}, 
				{
					"id": 3403,
					"is_verified": true,
					"bank_name": "二井住友銀行",
					"branch_name": "人形町支店",
					"account_type": "普通",
					"account_number": "1111115",
					"account_name": "タナカ"
				}]`,
			))
		}
	})

	srv := httptest.NewServer(handler)

	return srv
}
