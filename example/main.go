package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	tiktok "github.com/OPTIC7409/tiktok-api-go"
)

func main() {
	// Create a new client
	client := tiktok.NewClient(
		tiktok.WithLanguage("en"),
		tiktok.WithRegion("US"),
	)

	// Create context
	ctx := context.Background()

	// Get trending videos
	fmt.Println("Fetching trending videos...")
	videos, err := client.GetTrendingVideos(ctx, 5)
	if err != nil {
		log.Fatal(err)
	}

	// Pretty print the results
	for i, video := range videos {
		fmt.Printf("\nVideo %d:\n", i+1)
		fmt.Printf("ID: %s\n", video.ID)
		fmt.Printf("Author: %s (@%s)\n", video.Author.Nickname, video.Author.UniqueID)
		fmt.Printf("Description: %s\n", video.Description)
		fmt.Printf("Stats: %d plays, %d likes, %d comments\n",
			video.Statistics.PlayCount,
			video.Statistics.DiggCount,
			video.Statistics.CommentCount,
		)

		// Get video comments
		comments, _, err := client.GetVideoComments(ctx, video.ID, 0)
		if err != nil {
			fmt.Printf("Failed to get comments: %v\n", err)
			continue
		}

		fmt.Printf("\nTop comments:\n")
		for j, comment := range comments {
			if j >= 3 {
				break
			}
			fmt.Printf("- %s: %s\n", comment.User.Nickname, comment.Text)
		}
		fmt.Println(strings.Repeat("-", 50))
	}

	// Save the last result as JSON for inspection
	if len(videos) > 0 {
		data, _ := json.MarshalIndent(videos[0], "", "  ")
		os.WriteFile("example_video.json", data, 0644)
		fmt.Println("\nSaved first video details to example_video.json")
	}
}
