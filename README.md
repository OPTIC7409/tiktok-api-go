# TikTok API Go Client

A Go client for the TikTok API that allows you to interact with TikTok's data programmatically.

## Features

- Get user profiles and videos
- Fetch video details and comments
- Get trending videos
- Configurable client options (language, region, user agent)

## Installation

```bash
go get github.com/OPTIC7409/tiktok-api-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    tiktok "github.com/OPTIC7409/tiktok-api-go"
)

func main() {
    // Create a new client
    client := tiktok.NewClient(
        tiktok.WithLanguage("en"),
        tiktok.WithRegion("US"),
    )

    // Get user profile
    ctx := context.Background()
    user, err := client.GetUserByUsername(ctx, "username")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User: %+v\n", user)

    // Get user's videos
    videos, cursor, err := client.GetUserVideos(ctx, user.ID, 0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Videos: %+v\n", videos)

    // Get trending videos
    trending, err := client.GetTrendingVideos(ctx, 20)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Trending: %+v\n", trending)

    // Get video comments
    comments, nextCursor, err := client.GetVideoComments(ctx, videos[0].ID, 0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Comments: %+v\n", comments)
}
```

## Configuration Options

The client can be configured with several options:

- `WithUserAgent(userAgent string)`: Set a custom User-Agent header
- `WithLanguage(language string)`: Set the preferred language (e.g., "en", "es")
- `WithRegion(region string)`: Set the preferred region (e.g., "US", "UK")

## Error Handling

All API methods return errors that can be handled using the standard Go error handling patterns. The errors include both network-related issues and API-specific error responses.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details 