package main

import (
	"fmt"
	"net/http"

	"github.com/brk-a/0x05-rss-feeds/internal/database"
)

type authenticatedHandler  func(http.ResponseWriter, *http.Request, database.User)

func (cgf *apiConfig) middlewareAuth(handler authenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err!=nil {
			respondWithError(w, 403, fmt.Sprintf("cannot fetch API key: %v", err))
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err!=nil {
			respondWithError(w, 400, fmt.Sprintf("cannot get user: %v", err))
		}
		
		handler(w, r, user)
	}
}