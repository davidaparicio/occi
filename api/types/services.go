package types

import "time"

type Services struct {
	Route   any `json:"route,omitempty"`
	Billing struct {
		NextBillingDate time.Time `json:"nextBillingDate,omitempty"`
		ExpirationDate  time.Time `json:"expirationDate,omitempty"`
		Plan            struct {
			Code        string `json:"code,omitempty"`
			InvoiceName string `json:"invoiceName,omitempty"`
		} `json:"plan,omitempty"`
		Pricing struct {
			Capacities      []string `json:"capacities,omitempty"`
			Description     string   `json:"description,omitempty"`
			Interval        int      `json:"interval,omitempty"`
			Duration        string   `json:"duration,omitempty"`
			MinimumQuantity int      `json:"minimumQuantity,omitempty"`
			MaximumQuantity int      `json:"maximumQuantity,omitempty"`
			MinimumRepeat   int      `json:"minimumRepeat,omitempty"`
			MaximumRepeat   any      `json:"maximumRepeat,omitempty"`
			Price           struct {
				CurrencyCode string  `json:"currencyCode,omitempty"`
				Text         string  `json:"text,omitempty"`
				Value        float64 `json:"value,omitempty"`
			} `json:"price,omitempty"`
			PriceInUcents           int    `json:"priceInUcents,omitempty"`
			PricingMode             string `json:"pricingMode,omitempty"`
			PricingType             string `json:"pricingType,omitempty"`
			EngagementConfiguration any    `json:"engagementConfiguration,omitempty"`
		} `json:"pricing,omitempty"`
		Group     any `json:"group,omitempty"`
		Lifecycle struct {
			Current struct {
				PendingActions  []any     `json:"pendingActions,omitempty"`
				TerminationDate any       `json:"terminationDate,omitempty"`
				CreationDate    time.Time `json:"creationDate,omitempty"`
				State           string    `json:"state,omitempty"`
			} `json:"current,omitempty"`
			Capacities struct {
				Actions []string `json:"actions,omitempty"`
			} `json:"capacities,omitempty"`
		} `json:"lifecycle,omitempty"`
		Renew struct {
			Current struct {
				Mode     string    `json:"mode,omitempty"`
				NextDate time.Time `json:"nextDate,omitempty"`
				Period   string    `json:"period,omitempty"`
			} `json:"current,omitempty"`
			Capacities struct {
				Mode []string `json:"mode,omitempty"`
			} `json:"capacities,omitempty"`
		} `json:"renew,omitempty"`
		Engagement        any `json:"engagement,omitempty"`
		EngagementRequest any `json:"engagementRequest,omitempty"`
	} `json:"billing,omitempty"`
	Resource struct {
		DisplayName string `json:"displayName,omitempty"`
		Name        string `json:"name,omitempty"`
		State       string `json:"state,omitempty"`
		Product     struct {
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"product,omitempty"`
		ResellingProvider any `json:"resellingProvider,omitempty"`
	} `json:"resource,omitempty"`
	ServiceID       int `json:"serviceId,omitempty"`
	ParentServiceID int `json:"parentServiceId,omitempty"`
	Customer        struct {
		Contacts []struct {
			CustomerCode string `json:"customerCode,omitempty"`
			Type         string `json:"type,omitempty"`
		} `json:"contacts,omitempty"`
	} `json:"customer,omitempty"`
	Tags []any `json:"tags,omitempty"`
}
