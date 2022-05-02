package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hararudoka/shrt/internal/handler/nethttp"
	"github.com/hararudoka/shrt/internal/service"
	"github.com/hararudoka/shrt/internal/storage"
)

func main() {
	db, err := storage.Open()
	if err != nil {
		panic(err)
	}

	s := service.New(db)

	handler := nethttp.New(*s)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), handler); err != nil {
		log.Fatal(err)
	}
}
