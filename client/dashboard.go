package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Dashboard object definition
type Dashboard struct {
	ID          uint64   `json:"id"`
	UID         string   `json:"uid"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	IsStarred   bool     `json:"isStarred"`
	URI         string   `json:"uri"`
	Slug        string   `json:"slug"`
	FolderID    uint64   `json:"folderId"`
	FolderUID   string   `json:"folderUid"`
	FolderTitle string   `json:"folderTitle"`
	FolderURL   string   `json:"folderUrl"`
}

// DashboardList list dashboards
func (c Client) DashboardList() ([]Dashboard, error) {
	resp, err := c.Get("/search", map[string]string{"type": "dash-db"})
	if err != nil {
		return nil, err
	}

	dashboards := []Dashboard{}
	err = json.Unmarshal(resp.Body(), &dashboards)
	if err != nil {
		return nil, err
	}
	return dashboards, nil
}

// DashboardImport import a dashboard
func (c Client) DashboardImport(fpath string, folderid uint64) error {
	contents, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}

	dashboardm := map[string]interface{}{}
	err = json.Unmarshal(contents, &dashboardm)
	if err != nil {
		return fmt.Errorf("Fail to extract dashboard: %s", err.Error())
	}

	dashboardm["id"] = nil
	dashboardm["uid"] = nil

	payload := map[string]interface{}{}
	payload["dashboard"] = dashboardm
	payload["folderId"] = folderid
	payload["overwrite"] = false

	_, err = c.Post("/dashboards/db", payload)
	if err != nil {
		return err
	}
	return nil
}

// DashboardExport export a dashboard
func (c Client) DashboardExport(fpath, uid string) error {
	resp, err := c.Get("/dashboards/uid/"+uid, nil)
	if err != nil {
		return err
	}

	// Format output and ignore the "meta" field
	var body []byte
	raw := resp.Body()
	dashboardm := map[string]interface{}{}
	err = json.Unmarshal(raw, &dashboardm)
	if err != nil {
		return fmt.Errorf("Fail to decode dashboard due to: %s", err.Error())
	}
	body, _ = json.MarshalIndent(dashboardm["dashboard"], "", "  ")

	// export the dashboard to a file
	err = ioutil.WriteFile(fpath, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

// DashboardDelete delete a dashboard
func (c Client) DashboardDelete(uid string) error {
	_, err := c.Delete("/dashboards/uid/" + uid)
	return err
}
