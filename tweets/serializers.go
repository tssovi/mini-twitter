package tweets

// FollowSerializer Serializer for follow API
type FollowSerializer struct {
	UserID     uint `json:"user_id"`
	FollowerID uint `json:"follower_id"`
}

// CreateTweetSerializer Serializer for creating tweet
type CreateTweetSerializer struct {
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}
