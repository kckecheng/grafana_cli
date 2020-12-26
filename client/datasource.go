package client

import (
	"encoding/json"
	"io/ioutil"
)

// DataSource object
type DataSource struct {
	ID          int64                  `json:"id"`
	OrgID       int64                  `json:"orgId"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	TypeLogoURL string                 `json:"typeLogoUrl"`
	Access      string                 `json:"access"`
	URL         string                 `json:"url"`
	Password    string                 `json:"password"`
	User        string                 `json:"User"`
	Database    string                 `json:"database"`
	BasicAuth   bool                   `json:"basicAuth"`
	IsDefault   bool                   `json:"isDefault"`
	JSONData    map[string]interface{} `json:"jsonData"`
	ReadOnly    bool                   `json:"readOnly"`
}

// DataSourceList list data sources
func (c Client) DataSourceList() ([]DataSource, error) {
	resp, err := c.Get("/datasources", nil)
	if err != nil {
		return nil, err
	}

	var dses []DataSource
	err = json.Unmarshal(resp.Body(), &dses)
	if err != nil {
		return nil, err
	}
	return dses, nil
}

// DataSourceExport export a data source
func (c Client) DataSourceExport(name, fpath string) error {
	resp, err := c.Get("/datasources/name/"+name, nil)
	if err != nil {
		return err
	}

	// Decode data source
	var ds map[string]interface{}
	err = json.Unmarshal(resp.Body(), &ds)
	if err != nil {
		return err
	}

	// Format for output
	body, err := json.MarshalIndent(ds, "", "  ")
	if err != nil {
		return err
	}

	// Write to an external file
	err = ioutil.WriteFile(fpath, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

// DataSourceImport import a data source
func (c Client) DataSourceImport(name, fpath string) error {
	var ds map[string]interface{}
	raw, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &ds)
	if err != nil {
		return err
	}

	if name != "" {
		ds["name"] = name
	}
	_, err = c.Post("/datasources", ds)
	if err != nil {
		return err
	}
	return nil
}

// DataSourceDelete delete a data source
func (c Client) DataSourceDelete(name string) error {
	_, err := c.Delete("/datasources/name/" + name)
	return err
}
