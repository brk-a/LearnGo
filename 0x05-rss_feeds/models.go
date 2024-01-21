package main

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	APIKey string `json:"api_key"`
}
func databaseUserToUser(dbUser database.User) User {
	return User {
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		APIKey: dbUser.ApiKey,
	}
}

type Feed struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Url string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}
func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed {
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}
func databaseFeedsToFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbf:=range dbFeed{
		feeds = append(feeds, databaseFeedToFeed(dbf))
	}

	return feeds
}

type FeedFollows struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}
func databaseFeedFollowToFeedFollow(dbFeedFollows database.FeedFollows) FeedFollows {
	return FeedFollows {
		ID: dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdatedAt: dbFeedFollows.UpdatedAt,
		UserID: dbFeedFollows.UserID,
		FeedID: dbFeedFollows.FeedID,
	}
}
unc databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollows) []FeedFollows {
	feedFollowss := []FeedFollows{}

	for _, dbff:=range dbFeedFollows{
		feedFollows = append(feeds, databaseFeedFollowToFeedFollow(dbff))
	}

	return feedFollows
}