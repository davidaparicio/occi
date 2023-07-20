package types

type Service struct {
	Route struct {
		URL  string `json:"url,omitempty"`
		Vars []struct {
			Value string `json:"value,omitempty"`
			Key   string `json:"key,omitempty"`
		} `json:"vars,omitempty"`
		Path string `json:"path,omitempty"`
	} `json:"route,omitempty"`
	CreationDate    string `json:"creationDate,omitempty"`
	Details         []any  `json:"details,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
	NextBillingDate string `json:"nextBillingDate,omitempty"`
	Plan            struct {
		Product struct {
			Name any `json:"name,omitempty"`
		} `json:"product,omitempty"`
		Code any `json:"code,omitempty"`
	} `json:"plan,omitempty"`
	EngagementDate any    `json:"engagementDate,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	Resource       struct {
		State       string `json:"state,omitempty"`
		DisplayName string `json:"displayName,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"resource,omitempty"`
	State string `json:"state,omitempty"`
	Renew struct {
		Mode              string   `json:"mode,omitempty"`
		DayOfMonth        int      `json:"dayOfMonth,omitempty"`
		PossibleIntervals []string `json:"possibleIntervals,omitempty"`
		PossibleModes     []string `json:"possibleModes,omitempty"`
		Interval          string   `json:"interval,omitempty"`
	} `json:"renew,omitempty"`
}
