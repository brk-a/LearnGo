package main

import (
	"net/http"

	"github.com/brk-a/0x05-rss-feeds/internal/database"
)

type authenticatedHandler  func(http.ResponseWriter, *http.Request, database.User)

func (cgf *apiConfig) middlewareAuth(handler authenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}