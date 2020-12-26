package client

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// Snapshot object
type Snapshot struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	OrgID       int64     `json:"orgId"`
	UserID      int64     `json:"userId"`
	External    bool      `json:"external"`
	ExternalURL string    `json:"externalUrl"`
	Expires     time.Time `json:"expires"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

// SnapshotList list snapshots
func (c Client) SnapshotList() ([]Snapshot, error) {
	resp, err := c.Get("/dashboard/snapshots", nil)
	if err != nil {
		return nil, err
	}

	var snapshots []Snapshot
	err = json.Unmarshal(resp.Body(), &snapshots)
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}

// SnapshotExport export an existing snapshot as a file which
// can be imported as a dashboard by calling DashboardImport
func (c Client) SnapshotExport(key, fpath string) error {
	resp, err := c.Get("/snapshots/"+key, nil)
	if err != nil {
		return err
	}

	var snapshot map[string]interface{}
	err = json.Unmarshal(resp.Body(), &snapshot)
	if err != nil {
		return nil
	}

	body, err := json.MarshalIndent(snapshot["dashboard"], "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fpath, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

// SnaposhotCreate - there is no way to create a snapshot with data directly with current API,
// please refer to https://community.grafana.com/t/snapshot-using-http-api-does-nothing/8834/18
// func (c Client) SnapshotCreate(dashboardUID string, from time.Time, to time.Time) error {
// 	return nil
// }

// SnapshotDelete delete a snapshot
func (c Client) SnapshotDelete(key string) error {
	_, err := c.Delete("/snapshots/" + key)
	return err
}
