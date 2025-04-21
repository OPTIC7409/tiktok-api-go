package tiktok

// Comment represents a TikTok comment
type Comment struct {
	ID         string `json:"cid"`
	Text       string `json:"text"`
	CreateTime int64  `json:"create_time"`
	User       struct {
		ID        string `json:"uid"`
		UniqueID  string `json:"unique_id"`
		Nickname  string `json:"nickname"`
		AvatarURL string `json:"avatar_thumb"`
	} `json:"user"`
	ReplyCount int64 `json:"reply_count"`
	DiggCount  int64 `json:"digg_count"`
	IsAuthor   bool  `json:"is_author"`
	IsPinned   bool  `json:"is_pinned"`
}
