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

func TestClient_GetCollateral(t *testing.T) {
	type fields struct {
		num int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetCollateralOutput
		wantErr bool
	}{
		{
			name: "1 collateral",
			fields: fields{
				num: 1,
			},
			want: &GetCollateralOutput{
				Collateral:        100000,
				OpenPositionPnl:   -715,
				RequireCollateral: 19857,
				KeepRate:          5.000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := serverCollateralMock(tt.fields.num)
			defer srv.Close()
			u, _ := url.Parse(srv.URL)

			c := Client{
				BaseURL: u,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
			}
			got, err := c.GetCollateral(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCollateral() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCollateral() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func serverCollateralMock(num int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/me/getcollateral", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`
			{
				"collateral": 100000,
				"open_position_pnl": -715,
				"require_collateral": 19857,
				"keep_rate": 5.000
			}`,
		))
	})

	srv := httptest.NewServer(handler)

	return srv
}
