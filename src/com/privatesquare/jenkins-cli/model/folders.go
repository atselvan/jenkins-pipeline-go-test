package model

type FolderInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Folders struct {
	Jobs []FolderInfo `json:"jobs"`
}
