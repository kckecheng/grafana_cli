package client

import (
	"encoding/json"
	"fmt"
	"time"
)

// Annotation object
type Annotation struct {
	ID          int64                  `json:"id"`
	AlertID     int64                  `json:"alertId"`
	AlertName   string                 `json:"alertName"`
	DashboardID int64                  `json:"dashboardId"`
	PanelID     int64                  `json:"panelId"`
	UserID      int64                  `json:"userId"`
	NewState    string                 `json:"newState"`
	PrevState   string                 `json:"prevState"`
	Created     int64                  `json:"created"`
	Updated     int64                  `json:"updated"`
	Time        int64                  `json:"time"`
	TimeEnd     int64                  `json:"timeEnd"`
	Text        string                 `json:"text"`
	Tags        []string               `json:"tags"`
	Login       string                 `json:"login"`
	Email       string                 `json:"email"`
	AvatarURL   string                 `json:"avatarUrl"`
	Data        map[string]interface{} `json:"data"`
}

// AnnotationList list annotations for a dashboard
// Limitation - only 100 (default) annotations can be gotten for specified time period
func (c Client) AnnotationList(dashboardID int64, from, to time.Time) ([]Annotation, error) {
	params := map[string]string{}
	if dashboardID > 0 {
		params["dashboardId"] = fmt.Sprintf("%d", dashboardID)
	}

	tzeroMS := timeToEpochMS(time.Unix(0, 0))
	fromMS := timeToEpochMS(from)
	toMS := timeToEpochMS(to)
	if fromMS > tzeroMS && toMS > tzeroMS && fromMS <= toMS {
		params["from"] = fmt.Sprintf("%d", fromMS)
		params["to"] = fmt.Sprintf("%d", toMS)
	}

	resp, err := c.Get("/annotations", params)
	if err != nil {
		return nil, err
	}

	var annotations []Annotation
	err = json.Unmarshal(resp.Body(), &annotations)
	if err != nil {
		return nil, err
	}
	return annotations, nil
}

// AnnotationCreate create an annotation
func (c Client) AnnotationCreate(dashboardID, panelID int64, from, to time.Time, text string, tag ...string) error {
	var tags []string
	for _, t := range tag {
		tags = append(tags, t)
	}
	payload := map[string]interface{}{
		"dashboardId": dashboardID,
		"panelId":     panelID,
		"time":        timeToEpochMS(from),
		"timeEnd":     timeToEpochMS(to),
		"text":        text,
		"tags":        tags,
	}

	_, err := c.Post("/annotations", payload)
	return err
}

// AnnotationDelete delete an annotation
func (c Client) AnnotationDelete(id int64) error {
	_, err := c.Delete(fmt.Sprintf("/annotations/%d", id))
	return err
}
