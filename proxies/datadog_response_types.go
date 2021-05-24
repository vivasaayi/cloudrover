package proxies

type DDMonitorSearchResponse struct {
	Counts   DDMonitorSearchSummary   `json:"counts"`
	Metadata DDMonitorsSearchMetadata `json:"metadata"`
	Monitors []DDMonitors             `json:"monitors"`
}

type DDMonitorSearchSummary struct {
	Status []DDSearchResponseSummary     `json:"status"`
	Muted  []DDSearchResponseBoolSummary `json:"muted"`
	Type   []DDSearchResponseSummary     `json:"type"`
}

type DDMonitors struct {
	Id              int      `json:"id"`
	Status          string   `json:"status"`
	Classification  string   `json:"classification"`
	LastTriggeredTs int      `json:"last_triggered_ts"`
	Name            string   `json:"name"`
	Query           string   `json:"query"`
	OrgId           int      `json:"org_id"`
	Type            string   `json:"type"`
	Tags            []string `json:"tags"`
}

type DDMonitorsSearchMetadata struct {
	TotalCont int `json:"total_count"`
	PageCount int `json:"page_count"`
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
}

type DDSearchResponseSummary struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

type DDSearchResponseBoolSummary struct {
	Count int  `json:"count"`
	Name  bool `json:"name"`
}
