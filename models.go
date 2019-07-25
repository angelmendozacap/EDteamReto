package main

// Photo defines the photo struct
type Photo struct {
	AlbumID      uint   `json:"albumId"`
	ID           uint   `json:"id,omitempty"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
