package main

import (
	"log"
	"net/http"

	"github.com/Nithiszz/sprint-api/pkg/app"
	"github.com/Nithiszz/sprint-api/pkg/service/book"
)

func main() {
	mux := http.NewServeMux()
	app.RegisterBookService(mux, book.New())
	log.Fatal(http.ListenAndServe(":8080", mux))
}
