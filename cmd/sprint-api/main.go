package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"

	"github.com/Nithiszz/sprint-api/pkg/app"
	"github.com/Nithiszz/sprint-api/pkg/service/book"
	"github.com/acoshift/ds"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var jsonConfig = []byte(`{
  "type": "service_account",
  "project_id": "sprintapi-159415",
  "private_key_id": "ad42e1dd826bca29c4c8cc9c105aa439159105c3",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCufYphpM9RbxmC\nmHtUi++nRl1lQ3xaLaNse6gUY6NrAv44WiF5NBWQl0ifnCb3WU2g7CfmFoKitFKw\nHcbzLyQk5xHJ2qxAhMYNL6OUc14Q0FxbxLoCd3nP5SqTzjPGtqqe1E1O/+3B6iBf\nG0Wv/b22WmLpuJQZJaLLP105JMmKVyRaW6ZGalgaASvaIEsXPrNNWE4m7ToRXI4d\nlCq8hA1ZCJOuOUBl+WgaIHTzvn8u0eIlRGXY8Roqvy7sVjd18KFabPP42ltUlf6W\nPRMlGQ8io00DOe5TQtjQ5UugkBaOH/ggbM/UuElBc/45Mq1EsGm6E2IXtSZygGvn\nMsZ9FHZDAgMBAAECggEAULcqCDkg3d6ywkKe6mFBHvPhyDILl/t8mXYqLiRZN+tO\nherLiTGauCQDKDInpEvfKQ2U9056Z6FajrV3jo7D/X4WMHXDMKU6qGbdGJK1dLmt\nv5WlJfb2lkDADVdZhBaDnq0+hcjFxunyx4vqFJsf4va0wsDrYSaTw8kv6nsl76PT\n6xhQ6VLJ7inV2NlUDDKagHhvm14Rrf8FrE20LgtFLYMEFF3wO3x0OXu0hzyBfp5q\nI/8ChPmOU/+8lRbj9wMq/H2Zr6+32dAG/B26gApQKXzx1/csnYOVk9pjsevdLuUL\nvu/nrxD132K2lsn/VHcxuwXo+UPrVwDCaj+B5MZl4QKBgQDhv7uzPqRsnlZX4kkr\ncBqusymxpujAvYWwDKt/sGtcUrBgeXXFFIxk/xchuncCbns029kMpDJqQZrvEEh6\nCMIHP41WZEqDGa7rZ3MM29aJS0TcjMeTNmpi9yPZRjoEIQ3igOpGgBhtlUmD32IH\nsl/D1DGnBrfGabGWDYlPWaGOLwKBgQDF32TWrVppiGdF1+TLxTbPLIyUbcwXIriJ\nuq4YBAc8LoEf/ez0megZlcjcTUzz/lI2O//7cJBzmstGRTUmQYPW8PNbZjUN9Q08\noGkpiPUaXOHW76HGURkEVDu5p5bUwj0g56lka58baGc6IIX7Agei34A3CXMm/kS9\no980SlMILQKBgQCzkE/b3oro/vUNCKhMzenbZhVXAr2Gefm5tApCEfEDyZ4Tzuj3\nb0XRG/qpUGlTXM/RlsHJxV14mWCEob4Um5zCKTHiMvn119sD0hB4fPDj2iQXDj+8\n//6VY6F3NN325Nfnf8VZeJB1sdZ895VBLAfye3lXMwfA4ddo1LGQlWnk7wKBgQC+\nJXh/m8KgW0Xypg7lijSrTcIh+IkBSopPQCeASI6zVHUdSyRjwWp6+6cznMzwQ1q5\nZ1sMQxVtIjlo46S5iIerC4ywLj2Zlf0MX5HvKf9vQAZ+R3UMYG5L7K4gMF4PQkD/\n196983XIHsHj59EYbtDrwR8yxE/2Dq38FvetBidYWQKBgQDJBb6x0P1DicWj3nUs\nOcdsqtB6ug09dYmuz3fGWkr/dd3mc2pVcOH//iiLHepJV3WBY3dRWcEn6zb8fSNQ\nVj2RkFv9wfVOawyXo+opkAcaoaqdNNJS9fqh56nixp1y73Ov01DNlE9CXRln3Pst\nu1zJppKttnxF6GUfGctNLZcH0A==\n-----END PRIVATE KEY-----\n",
  "client_email": "backend@sprintapi-159415.iam.gserviceaccount.com",
  "client_id": "114654682130742014455",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://accounts.google.com/o/oauth2/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/backend%40sprintapi-159415.iam.gserviceaccount.com"
}
`)

func main() {
	googleConfig, err := google.JWTConfigFromJSON(jsonConfig, datastore.ScopeDatastore)
	if err != nil {
		log.Fatal(err)
	}

	tokenSource := googleConfig.TokenSource(context.Background())

	client, err := ds.NewClient(context.Background(), "sprintapi-159415", option.WithTokenSource(tokenSource))

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	app.RegisterBookService(mux, book.New(book.Config{
		Datastore: client,
	}))

	handler := LogMiddleware(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

//LogResponseWriter result
type LogResponseWriter struct {
	http.ResponseWriter
	header int
}

func (w *LogResponseWriter) WriteHeader(v int) {
	w.header = v
	w.ResponseWriter.WriteHeader(v)
}

// LogMiddleware Log request server
func LogMiddleware(h http.Handler) http.Handler {
	// init var for LogMiddleware
	log.Println("log LogMiddleware inited!")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path
		tw := &LogResponseWriter{ResponseWriter: w}
		ip := r.RemoteAddr
		h.ServeHTTP(tw, r)
		end := time.Now()
		fmt.Printf("%s | %3d | %13v | %s | %s | %s\n", end.Format(time.RFC3339), tw.header, end.Sub(start), ip, r.Method, path)

	})
}
