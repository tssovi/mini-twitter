package tweets

import (
	"encoding/json"
	"net/http"

	"../db/models"
)

func getFollowingUsers(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	data := TweetService.getFollowingUsers(user)
	json.NewEncoder(w).Encode(data)
}

func getFollowers(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	list := TweetService.getFollowers(user)
	json.NewEncoder(w).Encode(list)
}

func followUser(w http.ResponseWriter, r *http.Request) {
	followSerializer := FollowSerializer{}
	err := json.NewDecoder(r.Body).Decode(&followSerializer)
	followSerializer.FollowerID = r.Context().Value("user").(*models.User).ID
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = TweetService.followUser(&followSerializer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("Follower added.")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	serializer := CreateTweetSerializer{}
	err := json.NewDecoder(r.Body).Decode(&serializer)
	serializer.UserID = r.Context().Value("user").(*models.User).ID
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = TweetService.createPost(&serializer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("Tweet successfully created.")
}

func getFeed(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)
	feed := TweetService.getFeed(user)
	json.NewEncoder(w).Encode(feed)
}
