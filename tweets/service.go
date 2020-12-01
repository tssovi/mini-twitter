package tweets

import (
	"errors"
	"log"

	"../common"
	"../db/models"
)

var db = common.Dependencies.DB

// ITweetService Service exposing tweet related operations
type ITweetService struct {
}

func (ts *ITweetService) getFollowingUsers(user *models.User) []map[string]interface{} {
	db.Preload("FollowingUsers.User").Find(&user)
	data := []map[string]interface{}{}
	for i := 0; i < len(user.FollowingUsers); i++ {
		data = append(data, map[string]interface{}{
			"username": user.FollowingUsers[i].User.Username,
			"name":     user.FollowingUsers[i].User.Name,
		})
	}
	return data
}

func (ts *ITweetService) getFollowers(user *models.User) []map[string]interface{} {
	db.Preload("Followers.FollowerUser").Find(&user)
	data := []map[string]interface{}{}
	for i := 0; i < len(user.Followers); i++ {
		data = append(data, map[string]interface{}{
			"username": user.Followers[i].FollowerUser.Username,
			"name":     user.Followers[i].FollowerUser.Name,
		})
	}
	return data
}

func (ts *ITweetService) getFeed(user *models.User) []map[string]interface{} {
	db.Preload("FollowingUsers").Find(&user)
	followingUsers := []uint{}
	for i := 0; i < len(user.FollowingUsers); i++ {
		followingUsers = append(followingUsers, user.FollowingUsers[i].UserID)
	}
	tweets := []models.Tweet{}
	db.Limit(100).Where("user_id IN (?)", followingUsers).Preload("User").Find(&tweets)
	log.Println(len(tweets))

	data := []map[string]interface{}{}
	for i := 0; i < len(tweets); i++ {
		data = append(data, map[string]interface{}{
			"username":   tweets[i].User.Username,
			"name":       tweets[i].User.Name,
			"content":    tweets[i].Content,
			"created_at": tweets[i].CreatedAt,
		})
	}
	return data
}

func (ts *ITweetService) createPost(serializer *CreateTweetSerializer) error {
	tweet := models.Tweet{
		UserID:  serializer.UserID,
		Content: serializer.Content,
	}
	errs := db.Create(&tweet).GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

func (ts *ITweetService) followUser(serializer *FollowSerializer) error {
	if serializer.UserID == serializer.FollowerID {
		return errors.New("cannot follow self")
	}
	follow := models.Follower{
		UserID:     serializer.UserID,
		FollowerID: serializer.FollowerID,
	}
	errs := db.Create(&follow).GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

// TweetService Singleton instance of TweetService
var TweetService = ITweetService{}
