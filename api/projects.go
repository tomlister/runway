package api

import (
	"encoding/json"
	"net/http"
	"time"
)

// TODO: Clarify struct types
// ProjectInfo represents the information regarding a project
type ProjectInfo struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Namespace struct {
		Owner string `json:"owner"`
		Slug  string `json:"slug"`
	} `json:"namespace"`
	Stats struct {
		Views           int `json:"views"`
		Downloads       int `json:"downloads"`
		RecentViews     int `json:"recentViews"`
		RecentDownloads int `json:"recentDownloads"`
		Stars           int `json:"stars"`
		Watchers        int `json:"watchers"`
	} `json:"stats"`
	Category    string    `json:"category"`
	Visibility  string    `json:"visibility"`
	Description string    `json:"description"`
	LastUpdated time.Time `json:"lastUpdated"`
	UserActions struct {
		Starred  bool `json:"starred"`
		Watching bool `json:"watching"`
		Flagged  bool `json:"flagged"`
	} `json:"userActions"`
	Settings struct {
		Homepage string `json:"homepage"`
		Issues   string `json:"issues"`
		Source   string `json:"source"`
		Support  string `json:"support"`
		License  struct {
			Name interface{} `json:"name"`
			URL  string      `json:"url"`
			Type interface{} `json:"type"`
		} `json:"license"`
		Donation struct {
			Enable         bool        `json:"enable"`
			Email          interface{} `json:"email"`
			DefaultAmount  int         `json:"defaultAmount"`
			OneTimeAmounts interface{} `json:"oneTimeAmounts"`
			MonthlyAmounts interface{} `json:"monthlyAmounts"`
		} `json:"donation"`
		Keywords  []string `json:"keywords"`
		ForumSync bool     `json:"forumSync"`
	} `json:"settings"`
	TopicID          interface{}   `json:"topicId"`
	PostID           interface{}   `json:"postId"`
	PromotedVersions []interface{} `json:"promotedVersions"`
}

func GetProject(author string, project string) (ProjectInfo, bool, error) {
	var info ProjectInfo
	projectsBaseURL := BASE_URL + "v1/projects/"
	resp, err := http.Get(projectsBaseURL + author + "/" + project)
	if err != nil {
		return info, false, err
	}
	if resp.StatusCode == 404 {
		return info, false, err
	}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return info, false, err
	}
	return info, true, nil
}
