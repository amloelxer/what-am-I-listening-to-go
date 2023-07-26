package models

type Album struct {
	Name        string `json:"name,omitempty"`
	Artist      string `json:"artist,omitempty"`
	ImageArtURL string `json:"imageArtURL,omitempty"`
}
