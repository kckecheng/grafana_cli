package client

import "encoding/json"

// Folder object
type Folder struct {
	ID        int64    `json:"id"`
	UID       string   `json:"uid"`
	Title     string   `json:"title"`
	URI       string   `json:"uri"`
	URL       string   `json:"url"`
	Slug      string   `json:"slug"`
	Type      string   `json:"type"`
	Tags      []string `json:"tags"`
	IsStarred bool     `json:"isStarred"`
}

// FolderList list folders
func (c Client) FolderList() ([]Folder, error) {
	resp, err := c.Get("/search", map[string]string{"type": "dash-folder"})
	if err != nil {
		return nil, err
	}

	var folders []Folder
	err = json.Unmarshal(resp.Body(), &folders)
	if err != nil {
		return nil, err
	}
	return folders, nil
}

// FolderCreate create a folder
func (c Client) FolderCreate(name string) error {
	_, err := c.Post("/folders", map[string]interface{}{"title": name})
	return err
}

// FolderDelete delete a folder
func (c Client) FolderDelete(uid string) error {
	_, err := c.Delete("/folders/" + uid)
	return err
}
