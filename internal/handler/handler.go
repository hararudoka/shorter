package handler

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"math/rand"
	"shorter/internal/storage"
	"time"
)

type handler struct {
	db *storage.DB
}

type Template struct {
	templates *template.Template
}

func New(db *storage.DB) *handler {
	return &handler{
		db: db,
	}
}

type Service interface {
	Register(h handler, g echo.Group)
}

func (h handler) REGISTER(group echo.Group, service Service) {
	service.Register(h, group)
}

func makeShort() string {
	rand.Seed(time.Now().Unix())
	var s string
	ra := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 6; i++ {
		s += string(ra[rand.Intn(len(ra))])
	}
	return s
}
