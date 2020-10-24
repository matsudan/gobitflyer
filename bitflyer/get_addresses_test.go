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
//func TestClient_GetAddressList(t *testing.T) {
//	srv := serverAddressesMock()
//	defer srv.Close()
//	u, _ := url.Parse(srv.URL)
//
//	tests := []struct {
//		name    string
//		want    *GetAddressListOutput
//		wantErr bool
//	}{
//		{
//			name: "Normal",
//			want: &GetAddressListOutput{
//				Addresses: []*Address{
//					{
//						Type: "Normal",
//						CurrencyCode: "BTC",
//						Address: "3AYrDq8zhF82NJ2ZaLwBMPmaNziaKPaxa7",
//					},
//					{
//						Type: "NORMAL",
//						CurrencyCode: "ETH",
//						Address: "0x7fbB2CC24a3C0cd3789a44e9073381Ca6470853f",
//					},
//				},
//			},
//		},
//	}
//		for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := Client{
//				BaseURL: u,
//				HTTPClient: &http.Client{
//					Timeout: time.Minute,
//				},
//			}
//			got, err := c.GetAddressList(context.Background())
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetAddressList() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAddressList() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func serverAddressesMock() *httptest.Server {
//	handler := http.NewServeMux()
//	handler.HandleFunc("/v1/getaddresses", func(w http.ResponseWriter, r *http.Request) {
//		_, _ = w.Write([]byte(`[
//  {
//    "type": "NORMAL",
//    "currency_code": "BTC",
//    "address": "3AYrDq8zhF82NJ2ZaLwBMPmaNziaKPaxa7"
//  },
//  {
//    "type": "NORMAL",
//    "currency_code": "ETH",
//    "address": "0x7fbB2CC24a3C0cd3789a44e9073381Ca6470853f"
//  }
//]`))
//	})
//
//	srv := httptest.NewServer(handler)
//
//	return srv
//}
