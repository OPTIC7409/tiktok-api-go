package tiktok

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

// User represents a TikTok user profile
type User struct {
	ID             string `json:"id"`
	UniqueID       string `json:"uniqueId"`
	Nickname       string `json:"nickname"`
	AvatarURL      string `json:"avatarUrl"`
	FollowerCount  int64  `json:"followerCount"`
	FollowingCount int64  `json:"followingCount"`
	HeartCount     int64  `json:"heartCount"`
	VideoCount     int64  `json:"videoCount"`
	Signature      string `json:"signature"`
	Verified       bool   `json:"verified"`
}

// GetUserByUsername fetches a user's profile by their username
func (c *Client) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("User-Agent", c.userAgent).
		SetQueryParam("uniqueId", username).
		SetResult(&User{}).
		Get("https://api2.musical.ly/aweme/v1/user/profile/other/")

	if err != nil {
		return nil, errors.Wrap(err, "failed to get user profile")
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get user profile: status code %d", resp.StatusCode())
	}

	user := resp.Result().(*User)
	return user, nil
}

// GetUserVideos fetches a list of videos from a user
func (c *Client) GetUserVideos(ctx context.Context, userID string, cursor int64) ([]*Video, int64, error) {
	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("User-Agent", c.userAgent).
		SetQueryParams(map[string]string{
			"user_id": userID,
			"count":   "30",
			"cursor":  fmt.Sprintf("%d", cursor),
		}).
		SetResult(&struct {
			Videos []*Video `json:"aweme_list"`
			Cursor int64    `json:"cursor"`
		}{}).
		Get("https://api2.musical.ly/aweme/v1/aweme/post/")

	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to get user videos")
	}

	if resp.IsError() {
		return nil, 0, fmt.Errorf("failed to get user videos: status code %d", resp.StatusCode())
	}

	result := resp.Result().(*struct {
		Videos []*Video `json:"aweme_list"`
		Cursor int64    `json:"cursor"`
	})

	return result.Videos, result.Cursor, nil
}
