package bitflyer
//
//import (
//	"context"
//	"net/http"
//	"net/http/httptest"
//	"net/url"
//	"reflect"
//	"testing"
//	"time"
//)
//
//func TestClient_GetCollateral(t *testing.T) {
//	srv := serverCollateralMock()
//	defer srv.Close()
//	u, _ := url.Parse(srv.URL)
//
//	tests := []struct {
//		name    string
//		want    *GetCollateralOutput
//		wantErr bool
//	}{
//		{
//			name: "Normal",
//			want: &GetCollateralOutput{
//				Collateral:        100000,
//				OpenPositionPnl:   -715,
//				RequireCollateral: 19857,
//				KeepRate:          5.000,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := Client{
//				BaseURL: u,
//				HTTPClient: &http.Client{
//					Timeout: time.Minute,
//				},
//			}
//			got, err := c.GetCollateral(context.Background())
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetCollateral() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetCollateral() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func serverCollateralMock() *httptest.Server {
//	handler := http.NewServeMux()
//	handler.HandleFunc("/v1/getcollateral", func(w http.ResponseWriter, r *http.Request) {
//		_, _ = w.Write([]byte(`{
//  "collateral": 100000,
//  "open_position_pnl": -715,
//  "require_collateral": 19857,
//  "keep_rate": 5.000
//}`))
//	})
//
//	srv := httptest.NewServer(handler)
//
//	return srv
//}
