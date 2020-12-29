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

// Panel object definition - only the id and title
type Panel struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

// DashboardList list dashboards
func (c Client) DashboardList() ([]Dashboard, error) {
	resp, err := c.Get("/search", map[string]string{"type": "dash-db"})
	if err != nil {
		return nil, err
	}
	code := resp.StatusCode()
	if code != 200 {
		return nil, fmt.Errorf("Unexpected return code: %v", code)
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

	resp, err := c.Post("/dashboards/db", payload)
	if err != nil {
		return err
	}
	code := resp.StatusCode()
	if code != 200 {
		return fmt.Errorf("Unexpected return code: %v", code)
	}
	return nil
}

// DashboardExport export a dashboard
func (c Client) DashboardExport(fpath, uid string) error {
	resp, err := c.Get("/dashboards/uid/"+uid, nil)
	if err != nil {
		return err
	}
	code := resp.StatusCode()
	if code != 200 {
		return fmt.Errorf("Unexpected return code: %v", code)
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
	resp, err := c.Delete("/dashboards/uid/" + uid)
	if err != nil {
		return err
	}
	code := resp.StatusCode()
	if code != 200 {
		return fmt.Errorf("Unexpected return code: %v", code)
	}
	return nil
}

// DashboardPanelList list panels of a dashboard
func (c Client) DashboardPanelList(uid string) ([]Panel, error) {
	resp, err := c.Get("/dashboards/uid/"+uid, nil)
	if err != nil {
		return nil, err
	}
	code := resp.StatusCode()
	if code != 200 {
		return nil, fmt.Errorf("Unexpected return code: %v", code)
	}

	dashboardDetail := struct {
		Meta      map[string]interface{} `json:"meta"`
		Dashboard struct {
			Panels []Panel `json:"panels"`
		} `json:"dashboard"`
	}{}
	err = json.Unmarshal(resp.Body(), &dashboardDetail)
	if err != nil {
		return nil, err
	}
	return dashboardDetail.Dashboard.Panels, nil
}
