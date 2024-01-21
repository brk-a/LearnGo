package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/brk-a/0x05-rss-feeds/internal/database"
)

func (apiCfg apiConfig)handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err!=nil {
		respondWithError(w, 400, fmt.Sprintf("cannot decode request body %v", err))
		return
	}

	feedFollows, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.ID,
	})
	if err!=nil {
		respondWithError(w, 400, fmt.Sprintf("cannot create feed follow %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollows))
}

func (apiCfg apiConfig)handlerGetAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetAllFeedFollows(r.Context())
	if err!=nil {
		respondWithError(w, 400, fmt.Sprintf("cannot get feed follows %v", err))
		return
	}
	
	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}