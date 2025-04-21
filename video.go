package tiktok

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

// Video represents a TikTok video
type Video struct {
	ID          string `json:"aweme_id"`
	Description string `json:"desc"`
	CreateTime  int64  `json:"create_time"`
	Author      struct {
		ID       string `json:"uid"`
		UniqueID string `json:"unique_id"`
		Nickname string `json:"nickname"`
	} `json:"author"`
	Statistics struct {
		PlayCount     int64 `json:"play_count"`
		DiggCount     int64 `json:"digg_count"`
		ShareCount    int64 `json:"share_count"`
		CommentCount  int64 `json:"comment_count"`
		DownloadCount int64 `json:"download_count"`
	} `json:"statistics"`
	Music struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Author   string `json:"author"`
		Duration int    `json:"duration"`
		PlayURL  string `json:"play_url"`
	} `json:"music"`
	Video struct {
		PlayAddr struct {
			URLList []string `json:"url_list"`
		} `json:"play_addr"`
		Duration int `json:"duration"`
		Width    int `json:"width"`
		Height   int `json:"height"`
	} `json:"video"`
}

// GetVideoByID fetches a video by its ID
func (c *Client) GetVideoByID(ctx context.Context, videoID string) (*Video, error) {
	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("User-Agent", c.userAgent).
		SetQueryParam("aweme_id", videoID).
		SetResult(&struct {
			Video *Video `json:"aweme_detail"`
		}{}).
		Get("https://api2.musical.ly/aweme/v1/aweme/detail/")

	if err != nil {
		return nil, errors.Wrap(err, "failed to get video")
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get video: status code %d", resp.StatusCode())
	}

	result := resp.Result().(*struct {
		Video *Video `json:"aweme_detail"`
	})

	return result.Video, nil
}

// GetTrendingVideos fetches trending videos
func (c *Client) GetTrendingVideos(ctx context.Context, count int) ([]*Video, error) {
	if count <= 0 {
		count = 20
	}

	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("User-Agent", c.userAgent).
		SetQueryParams(map[string]string{
			"count":    fmt.Sprintf("%d", count),
			"language": c.language,
			"region":   c.region,
		}).
		SetResult(&struct {
			Videos []*Video `json:"aweme_list"`
		}{}).
		Get("https://api2.musical.ly/aweme/v1/feed/")

	if err != nil {
		return nil, errors.Wrap(err, "failed to get trending videos")
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get trending videos: status code %d", resp.StatusCode())
	}

	result := resp.Result().(*struct {
		Videos []*Video `json:"aweme_list"`
	})

	return result.Videos, nil
}

// GetVideoComments fetches comments for a video
func (c *Client) GetVideoComments(ctx context.Context, videoID string, cursor int64) ([]*Comment, int64, error) {
	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetHeader("User-Agent", c.userAgent).
		SetQueryParams(map[string]string{
			"aweme_id": videoID,
			"count":    "20",
			"cursor":   fmt.Sprintf("%d", cursor),
		}).
		SetResult(&struct {
			Comments []*Comment `json:"comments"`
			Cursor   int64      `json:"cursor"`
		}{}).
		Get("https://api2.musical.ly/aweme/v1/comment/list/")

	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to get video comments")
	}

	if resp.IsError() {
		return nil, 0, fmt.Errorf("failed to get video comments: status code %d", resp.StatusCode())
	}

	result := resp.Result().(*struct {
		Comments []*Comment `json:"comments"`
		Cursor   int64      `json:"cursor"`
	})

	return result.Comments, result.Cursor, nil
}
