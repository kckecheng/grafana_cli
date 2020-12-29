package client

import (
	"encoding/json"
	"fmt"
)

// Folder object
type Folder struct {
	ID        uint64   `json:"id"`
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
	code := resp.StatusCode()
	if code != 200 {
		return nil, fmt.Errorf("Unexpected return code: %v", code)
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
	resp, err := c.Post("/folders", map[string]interface{}{"title": name})
	if err != nil {
		return err
	}
	code := resp.StatusCode()
	if code != 200 {
		return fmt.Errorf("Unexpected return code: %v", code)
	}
	return nil
}

// FolderDelete delete a folder
func (c Client) FolderDelete(uid string) error {
	resp, err := c.Delete("/folders/" + uid)
	if err != nil {
		return err
	}
	code := resp.StatusCode()
	if code != 200 {
		return fmt.Errorf("Unexpected return code: %v", code)
	}
	return nil
}
